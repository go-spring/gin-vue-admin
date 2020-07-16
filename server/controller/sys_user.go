package controller

import (
	"fmt"
	"mime/multipart"
	"time"

	"gin-vue-admin/filter"
	"gin-vue-admin/global"
	"gin-vue-admin/global/response"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
	resp "gin-vue-admin/model/response"
	"gin-vue-admin/service"
	"gin-vue-admin/utils"

	"github.com/dchest/captcha"
	"github.com/dgrijalva/jwt-go"
	"github.com/go-redis/redis"
	"github.com/go-spring/go-spring-web/spring-web"
	"github.com/go-spring/go-spring/spring-boot"
)

func init() {
	SpringBoot.RegisterBean(new(BaseController)).Init(func(c *BaseController) {

		r := SpringBoot.Route("/base")

		r.PostMapping("/login", c.Login)
		r.PostMapping("/captcha", c.Captcha)
		r.PostMapping("/register", c.Register)
		r.GetMapping("/captcha/:captchaId", c.CaptchaImg)
	})

	SpringBoot.RegisterBean(new(UserController)).Init(func(c *UserController) {

		r := SpringBoot.Route("/user",
			SpringBoot.FilterBean((*filter.JwtFilter)(nil)),
			SpringBoot.FilterBean((*filter.CasbinRcbaFilter)(nil)))

		r.PostMapping("/changePassword", c.ChangePassword)
		r.PostMapping("/uploadHeaderImg", c.UploadHeaderImg)
		r.PostMapping("/getUserList", c.GetUserList)
		r.PostMapping("/setUserAuthority", c.SetUserAuthority)
		r.DELETE("/deleteUser", c.DeleteUser)
	})
}

type BaseController struct {
	SysUserService      *service.SysUserService      `autowire:""`
	JwtBlackListService *service.JwtBlackListService `autowire:""`
}

// @Tags Base
// @Summary 用户注册账号
// @Produce  application/json
// @Param data body model.SysUser true "用户注册接口"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"注册成功"}"
// @Router /base/register [post]
func (controller *BaseController) Register(webCtx SpringWeb.WebContext) {
	var R request.RegisterStruct
	_ = webCtx.Bind(&R)
	UserVerify := utils.Rules{
		"Username":    {utils.NotEmpty()},
		"NickName":    {utils.NotEmpty()},
		"Password":    {utils.NotEmpty()},
		"AuthorityId": {utils.NotEmpty()},
	}
	UserVerifyErr := utils.Verify(R, UserVerify)
	if UserVerifyErr != nil {
		response.FailWithMessage(UserVerifyErr.Error(), webCtx)
		return
	}
	user := &model.SysUser{Username: R.Username, NickName: R.NickName, Password: R.Password, HeaderImg: R.HeaderImg, AuthorityId: R.AuthorityId}
	err, userReturn := controller.SysUserService.Register(*user)
	if err != nil {
		response.FailWithDetailed(response.ERROR, resp.SysUserResponse{User: userReturn}, fmt.Sprintf("%v", err), webCtx)
	} else {
		response.OkDetailed(resp.SysUserResponse{User: userReturn}, "注册成功", webCtx)
	}
}

// @Tags Base
// @Summary 用户登录
// @Produce  application/json
// @Param data body request.RegisterAndLoginStruct true "用户登录接口"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"登陆成功"}"
// @Router /base/login [post]
func (controller *BaseController) Login(webCtx SpringWeb.WebContext) {
	var L request.RegisterAndLoginStruct
	_ = webCtx.Bind(&L)

	UserVerify := utils.Rules{
		"CaptchaId": {utils.NotEmpty()},
		"Captcha":   {utils.NotEmpty()},
		"Username":  {utils.NotEmpty()},
		"Password":  {utils.NotEmpty()},
	}

	UserVerifyErr := utils.Verify(L, UserVerify)
	if UserVerifyErr != nil {
		response.FailWithMessage(UserVerifyErr.Error(), webCtx)
		return
	}

	if captcha.VerifyString(L.CaptchaId, L.Captcha) {
		U := &model.SysUser{Username: L.Username, Password: L.Password}
		if err, user := controller.SysUserService.Login(U); err != nil {
			response.FailWithMessage(fmt.Sprintf("用户名密码错误或%v", err), webCtx)
		} else {
			controller.tokenNext(webCtx, *user)
		}
	} else {
		response.FailWithMessage("验证码错误", webCtx)
	}
}

// 登录以后签发jwt
func (controller *BaseController) tokenNext(webCtx SpringWeb.WebContext, user model.SysUser) {

	j := &filter.JWT{
		SigningKey: []byte(global.GVA_CONFIG.JWT.SigningKey), // 唯一签名
	}
	clams := request.CustomClaims{
		UUID:        user.UUID,
		ID:          user.ID,
		NickName:    user.NickName,
		AuthorityId: user.AuthorityId,
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix() - 1000,       // 签名生效时间
			ExpiresAt: time.Now().Unix() + 60*60*24*7, // 过期时间 一周
			Issuer:    "qmPlus",                       // 签名的发行者
		},
	}
	token, err := j.CreateToken(clams)
	if err != nil {
		response.FailWithMessage("获取token失败", webCtx)
		return
	}
	if !global.GVA_CONFIG.System.UseMultipoint {
		response.OkWithData(resp.LoginResponse{
			User:      user,
			Token:     token,
			ExpiresAt: clams.StandardClaims.ExpiresAt * 1000,
		}, webCtx)
		return
	}
	var loginJwt model.JwtBlacklist
	loginJwt.Jwt = token
	err, jwtStr := controller.JwtBlackListService.GetRedisJWT(user.Username)
	if err == redis.Nil {
		if err := controller.JwtBlackListService.SetRedisJWT(loginJwt, user.Username); err != nil {
			response.FailWithMessage("设置登录状态失败", webCtx)
			return
		}
		response.OkWithData(resp.LoginResponse{
			User:      user,
			Token:     token,
			ExpiresAt: clams.StandardClaims.ExpiresAt * 1000,
		}, webCtx)
	} else if err != nil {
		response.FailWithMessage(fmt.Sprintf("%v", err), webCtx)
	} else {
		var blackJWT model.JwtBlacklist
		blackJWT.Jwt = jwtStr
		if err := controller.JwtBlackListService.JsonInBlacklist(blackJWT); err != nil {
			response.FailWithMessage("jwt作废失败", webCtx)
			return
		}
		if err := controller.JwtBlackListService.SetRedisJWT(loginJwt, user.Username); err != nil {
			response.FailWithMessage("设置登录状态失败", webCtx)
			return
		}
		response.OkWithData(resp.LoginResponse{
			User:      user,
			Token:     token,
			ExpiresAt: clams.StandardClaims.ExpiresAt * 1000,
		}, webCtx)
	}
}

type UserController struct {
	SysUserService *service.SysUserService `autowire:""`
}

// @Tags SysUser
// @Summary 用户修改密码
// @Security ApiKeyAuth
// @Produce  application/json
// @Param data body request.ChangePasswordStruct true "用户修改密码"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"修改成功"}"
// @Router /user/changePassword [put]
func (controller *UserController) ChangePassword(webCtx SpringWeb.WebContext) {
	var params request.ChangePasswordStruct
	_ = webCtx.Bind(&params)
	UserVerify := utils.Rules{
		"Username":    {utils.NotEmpty()},
		"Password":    {utils.NotEmpty()},
		"NewPassword": {utils.NotEmpty()},
	}
	UserVerifyErr := utils.Verify(params, UserVerify)
	if UserVerifyErr != nil {
		response.FailWithMessage(UserVerifyErr.Error(), webCtx)
		return
	}
	U := &model.SysUser{Username: params.Username, Password: params.Password}
	if err, _ := controller.SysUserService.ChangePassword(U, params.NewPassword); err != nil {
		response.FailWithMessage("修改失败，请检查用户名密码", webCtx)
	} else {
		response.OkWithMessage("修改成功", webCtx)
	}
}

type UserHeaderImg struct {
	HeaderImg multipart.File `json:"headerImg"`
}

// @Tags SysUser
// @Summary 用户上传头像
// @Security ApiKeyAuth
// @accept multipart/form-data
// @Produce  application/json
// @Param headerImg formData file true "用户上传头像"
// @Param username formData string true "用户上传头像"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"上传成功"}"
// @Router /user/uploadHeaderImg [post]
func (controller *UserController) UploadHeaderImg(webCtx SpringWeb.WebContext) {
	claims := webCtx.Get("claims")
	// 获取头像文件
	// 这里我们通过断言获取 claims内的所有内容
	waitUse := claims.(*request.CustomClaims)
	uuid := waitUse.UUID
	header, err := webCtx.FormFile("headerImg")
	// 便于找到用户 以后从jwt中取
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("上传文件失败，%v", err), webCtx)
	} else {
		// 文件上传后拿到文件路径
		err, filePath, _ := utils.Upload(header)
		if err != nil {
			response.FailWithMessage(fmt.Sprintf("接收返回值失败，%v", err), webCtx)
		} else {
			// 修改数据库后得到修改后的user并且返回供前端使用
			err, user := controller.SysUserService.UploadHeaderImg(uuid, filePath)
			if err != nil {
				response.FailWithMessage(fmt.Sprintf("修改数据库链接失败，%v", err), webCtx)
			} else {
				response.OkWithData(resp.SysUserResponse{User: *user}, webCtx)
			}
		}
	}
}

// @Tags SysUser
// @Summary 分页获取用户列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取用户列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /user/getUserList [post]
func (controller *UserController) GetUserList(webCtx SpringWeb.WebContext) {
	var pageInfo request.PageInfo
	_ = webCtx.Bind(&pageInfo)
	PageVerifyErr := utils.Verify(pageInfo, utils.CustomizeMap["PageVerify"])
	if PageVerifyErr != nil {
		response.FailWithMessage(PageVerifyErr.Error(), webCtx)
		return
	}
	err, list, total := controller.SysUserService.GetUserInfoList(pageInfo)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("获取数据失败，%v", err), webCtx)
	} else {
		response.OkWithData(resp.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, webCtx)
	}
}

// @Tags SysUser
// @Summary 设置用户权限
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.SetUserAuth true "设置用户权限"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"修改成功"}"
// @Router /user/setUserAuthority [post]
func (controller *UserController) SetUserAuthority(webCtx SpringWeb.WebContext) {
	var sua request.SetUserAuth
	_ = webCtx.Bind(&sua)
	UserVerify := utils.Rules{
		"UUID":        {utils.NotEmpty()},
		"AuthorityId": {utils.NotEmpty()},
	}
	UserVerifyErr := utils.Verify(sua, UserVerify)
	if UserVerifyErr != nil {
		response.FailWithMessage(UserVerifyErr.Error(), webCtx)
		return
	}
	err := controller.SysUserService.SetUserAuthority(sua.UUID, sua.AuthorityId)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("修改失败，%v", err), webCtx)
	} else {
		response.OkWithMessage("修改成功", webCtx)
	}
}

// @Tags SysUser
// @Summary 删除用户
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.GetById true "删除用户"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"修改成功"}"
// @Router /user/deleteUser [delete]
func (controller *UserController) DeleteUser(webCtx SpringWeb.WebContext) {
	var reqId request.GetById
	_ = webCtx.Bind(&reqId)
	IdVerifyErr := utils.Verify(reqId, utils.CustomizeMap["IdVerify"])
	if IdVerifyErr != nil {
		response.FailWithMessage(IdVerifyErr.Error(), webCtx)
		return
	}
	err := controller.SysUserService.DeleteUser(reqId.Id)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("删除失败，%v", err), webCtx)
	} else {
		response.OkWithMessage("删除成功", webCtx)
	}
}
