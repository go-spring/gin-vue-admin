package controller

import (
	"fmt"

	"gin-vue-admin/filter"
	"gin-vue-admin/global/response"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
	resp "gin-vue-admin/model/response"
	"gin-vue-admin/service"
	"gin-vue-admin/utils"

	"github.com/go-spring/go-spring-web/spring-web"
	"github.com/go-spring/go-spring/spring-boot"
)

func init() {
	SpringBoot.RegisterBean(new(AuthorityController)).Init(func(c *AuthorityController) {

		r := SpringBoot.Route("/authority",
			SpringBoot.FilterBean((*filter.JwtFilter)(nil)),
			SpringBoot.FilterBean((*filter.CasbinRcbaFilter)(nil)))

		r.PostMapping("/createAuthority", c.CreateAuthority)
		r.PostMapping("/deleteAuthority", c.DeleteAuthority)
		r.PUT("/updateAuthority", c.UpdateAuthority)
		r.PostMapping("/copyAuthority", c.CopyAuthority)
		r.PostMapping("/getAuthorityList", c.GetAuthorityList)
		r.PostMapping("/setDataAuthority", c.SetDataAuthority)
	})
}

type AuthorityController struct {
	SysAuthorityService *service.SysAuthorityService `autowire:""`
}

// @Tags authority
// @Summary 创建角色
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SysAuthority true "创建角色"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /authority/createAuthority [post]
func (controller *AuthorityController) CreateAuthority(webCtx SpringWeb.WebContext) {
	var auth model.SysAuthority
	_ = webCtx.Bind(&auth)
	AuthorityVerify := utils.Rules{
		"AuthorityId":   {utils.NotEmpty()},
		"AuthorityName": {utils.NotEmpty()},
		"ParentId":      {utils.NotEmpty()},
	}
	AuthorityVerifyErr := utils.Verify(auth, AuthorityVerify)
	if AuthorityVerifyErr != nil {
		response.FailWithMessage(AuthorityVerifyErr.Error(), webCtx)
		return
	}
	err, authBack := controller.SysAuthorityService.CreateAuthority(auth)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("创建失败，%v", err), webCtx)
	} else {
		response.OkWithData(resp.SysAuthorityResponse{Authority: authBack}, webCtx)
	}
}

// @Tags authority
// @Summary 拷贝角色
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body response.SysAuthorityCopyResponse true "拷贝角色"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"拷贝成功"}"
// @Router /authority/copyAuthority [post]
func (controller *AuthorityController) CopyAuthority(webCtx SpringWeb.WebContext) {
	var copyInfo resp.SysAuthorityCopyResponse
	_ = webCtx.Bind(&copyInfo)
	OldAuthorityVerify := utils.Rules{
		"OldAuthorityId": {utils.NotEmpty()},
	}
	OldAuthorityVerifyErr := utils.Verify(copyInfo, OldAuthorityVerify)
	if OldAuthorityVerifyErr != nil {
		response.FailWithMessage(OldAuthorityVerifyErr.Error(), webCtx)
		return
	}
	AuthorityVerify := utils.Rules{
		"AuthorityId":   {utils.NotEmpty()},
		"AuthorityName": {utils.NotEmpty()},
		"ParentId":      {utils.NotEmpty()},
	}
	AuthorityVerifyErr := utils.Verify(copyInfo.Authority, AuthorityVerify)
	if AuthorityVerifyErr != nil {
		response.FailWithMessage(AuthorityVerifyErr.Error(), webCtx)
		return
	}
	err, authBack := controller.SysAuthorityService.CopyAuthority(copyInfo)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("拷贝失败，%v", err), webCtx)
	} else {
		response.OkWithData(resp.SysAuthorityResponse{Authority: authBack}, webCtx)
	}
}

// @Tags authority
// @Summary 删除角色
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SysAuthority true "删除角色"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /authority/deleteAuthority [post]
func (controller *AuthorityController) DeleteAuthority(webCtx SpringWeb.WebContext) {
	var a model.SysAuthority
	_ = webCtx.Bind(&a)
	AuthorityIdVerifyErr := utils.Verify(a, utils.CustomizeMap["AuthorityIdVerify"])
	if AuthorityIdVerifyErr != nil {
		response.FailWithMessage(AuthorityIdVerifyErr.Error(), webCtx)
		return
	}
	// 删除角色之前需要判断是否有用户正在使用此角色
	err := controller.SysAuthorityService.DeleteAuthority(&a)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("删除失败，%v", err), webCtx)
	} else {
		response.OkWithMessage("删除成功", webCtx)
	}
}

// @Tags authority
// @Summary 设置角色资源权限
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SysAuthority true "设置角色资源权限"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"设置成功"}"
// @Router /authority/updateAuthority [post]
func (controller *AuthorityController) UpdateAuthority(webCtx SpringWeb.WebContext) {
	var auth model.SysAuthority
	_ = webCtx.Bind(&auth)
	AuthorityVerify := utils.Rules{
		"AuthorityId":   {utils.NotEmpty()},
		"AuthorityName": {utils.NotEmpty()},
		"ParentId":      {utils.NotEmpty()},
	}
	AuthorityVerifyErr := utils.Verify(auth, AuthorityVerify)
	if AuthorityVerifyErr != nil {
		response.FailWithMessage(AuthorityVerifyErr.Error(), webCtx)
		return
	}
	err, authority := controller.SysAuthorityService.UpdateAuthority(auth)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("更新失败，%v", err), webCtx)
	} else {
		response.OkWithData(resp.SysAuthorityResponse{Authority: authority}, webCtx)
	}
}

// @Tags authority
// @Summary 分页获取角色列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取用户列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /authority/getAuthorityList [post]
func (controller *AuthorityController) GetAuthorityList(webCtx SpringWeb.WebContext) {
	var pageInfo request.PageInfo
	_ = webCtx.Bind(&pageInfo)
	PageVerifyErr := utils.Verify(pageInfo, utils.CustomizeMap["PageVerify"])
	if PageVerifyErr != nil {
		response.FailWithMessage(PageVerifyErr.Error(), webCtx)
		return
	}
	err, list, total := controller.SysAuthorityService.GetAuthorityInfoList(pageInfo)
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

// @Tags authority
// @Summary 设置角色资源权限
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SysAuthority true "设置角色资源权限"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"设置成功"}"
// @Router /authority/setDataAuthority [post]
func (controller *AuthorityController) SetDataAuthority(webCtx SpringWeb.WebContext) {
	var auth model.SysAuthority
	_ = webCtx.Bind(&auth)
	AuthorityIdVerifyErr := utils.Verify(auth, utils.CustomizeMap["AuthorityIdVerify"])
	if AuthorityIdVerifyErr != nil {
		response.FailWithMessage(AuthorityIdVerifyErr.Error(), webCtx)
		return
	}
	err := controller.SysAuthorityService.SetDataAuthority(auth)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("设置关联失败，%v", err), webCtx)
	} else {
		response.Ok(webCtx)
	}
}
