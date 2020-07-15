package controller

import (
	"fmt"

	"gin-vue-admin/global/response"
	"gin-vue-admin/middleware"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
	resp "gin-vue-admin/model/response"
	"gin-vue-admin/service"
	"gin-vue-admin/utils"

	"github.com/go-spring/go-spring-web/spring-gin"
	"github.com/go-spring/go-spring-web/spring-web"
	"github.com/go-spring/go-spring/spring-boot"
)

func init() {
	SpringBoot.RegisterBean(new(ApiController)).Init(func(c *ApiController) {

		r := SpringBoot.Route("/api",
			SpringGin.Filter(middleware.JWTAuth()),
			SpringGin.Filter(middleware.CasbinHandler()))

		r.PostMapping("/createApi", c.CreateApi)
		r.PostMapping("/deleteApi", c.DeleteApi)
		r.PostMapping("/getApiList", c.GetApiList)
		r.PostMapping("/getApiById", c.GetApiById)
		r.PostMapping("/updateApi", c.UpdateApi)
		r.PostMapping("/getAllApis", c.GetAllApis)
	})
}

type ApiController struct {
	SysApiService *service.SysApiService `autowire:""`
}

// @Tags SysApi
// @Summary 创建基础api
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SysApi true "创建api"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /api/createApi [post]
func (controller *ApiController) CreateApi(webCtx SpringWeb.WebContext) {
	var api model.SysApi
	_ = webCtx.Bind(&api)
	ApiVerify := utils.Rules{
		"Path":        {utils.NotEmpty()},
		"Description": {utils.NotEmpty()},
		"ApiGroup":    {utils.NotEmpty()},
		"Method":      {utils.NotEmpty()},
	}
	ApiVerifyErr := utils.Verify(api, ApiVerify)
	if ApiVerifyErr != nil {
		response.FailWithMessage(ApiVerifyErr.Error(), webCtx)
		return
	}
	err := controller.SysApiService.CreateApi(api)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("创建失败，%v", err), webCtx)
	} else {
		response.OkWithMessage("创建成功", webCtx)
	}
}

// @Tags SysApi
// @Summary 删除指定api
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SysApi true "删除api"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /api/deleteApi [post]
func (controller *ApiController) DeleteApi(webCtx SpringWeb.WebContext) {
	var a model.SysApi
	_ = webCtx.Bind(&a)
	ApiVerify := utils.Rules{
		"ID": {utils.NotEmpty()},
	}
	ApiVerifyErr := utils.Verify(a.Model, ApiVerify)
	if ApiVerifyErr != nil {
		response.FailWithMessage(ApiVerifyErr.Error(), webCtx)
		return
	}
	err := service.DeleteApi(a)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("删除失败，%v", err), webCtx)
	} else {
		response.OkWithMessage("删除成功", webCtx)
	}
}

// 条件搜索后端看此api

// @Tags SysApi
// @Summary 分页获取API列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.SearchApiParams true "分页获取API列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /api/getApiList [post]
func (controller *ApiController) GetApiList(webCtx SpringWeb.WebContext) {
	// 此结构体仅本方法使用
	var sp request.SearchApiParams
	_ = webCtx.Bind(&sp)
	PageVerifyErr := utils.Verify(sp.PageInfo, utils.CustomizeMap["PageVerify"])
	if PageVerifyErr != nil {
		response.FailWithMessage(PageVerifyErr.Error(), webCtx)
		return
	}
	err, list, total := service.GetAPIInfoList(sp.SysApi, sp.PageInfo, sp.OrderKey, sp.Desc)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("获取数据失败，%v", err), webCtx)
	} else {
		response.OkWithData(resp.PageResult{
			List:     list,
			Total:    total,
			Page:     sp.PageInfo.Page,
			PageSize: sp.PageInfo.PageSize,
		}, webCtx)
	}
}

// @Tags SysApi
// @Summary 根据id获取api
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.GetById true "根据id获取api"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /api/getApiById [post]
func (controller *ApiController) GetApiById(webCtx SpringWeb.WebContext) {
	var idInfo request.GetById
	_ = webCtx.Bind(&idInfo)
	IdVerifyErr := utils.Verify(idInfo, utils.CustomizeMap["IdVerify"])
	if IdVerifyErr != nil {
		response.FailWithMessage(IdVerifyErr.Error(), webCtx)
		return
	}
	err, api := service.GetApiById(idInfo.Id)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("获取数据失败，%v", err), webCtx)
	} else {
		response.OkWithData(resp.SysAPIResponse{Api: api}, webCtx)
	}
}

// @Tags SysApi
// @Summary 创建基础api
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SysApi true "创建api"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /api/updateApi [post]
func (controller *ApiController) UpdateApi(webCtx SpringWeb.WebContext) {
	var api model.SysApi
	_ = webCtx.Bind(&api)
	ApiVerify := utils.Rules{
		"Path":        {utils.NotEmpty()},
		"Description": {utils.NotEmpty()},
		"ApiGroup":    {utils.NotEmpty()},
		"Method":      {utils.NotEmpty()},
	}
	ApiVerifyErr := utils.Verify(api, ApiVerify)
	if ApiVerifyErr != nil {
		response.FailWithMessage(ApiVerifyErr.Error(), webCtx)
		return
	}
	err := service.UpdateApi(api)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("修改数据失败，%v", err), webCtx)
	} else {
		response.OkWithMessage("修改数据成功", webCtx)
	}
}

// @Tags SysApi
// @Summary 获取所有的Api 不分页
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /api/getAllApis [post]
func (controller *ApiController) GetAllApis(webCtx SpringWeb.WebContext) {
	err, apis := service.GetAllApis()
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("获取数据失败，%v", err), webCtx)
	} else {
		response.OkWithData(resp.SysAPIListResponse{Apis: apis}, webCtx)
	}
}
