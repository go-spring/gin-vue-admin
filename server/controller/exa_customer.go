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
	SpringBoot.RegisterBean(new(CustomerController)).Init(func(c *CustomerController) {

		r := SpringBoot.Route("/customer",
			SpringGin.Filter(middleware.JWTAuth()),
			SpringGin.Filter(middleware.CasbinHandler()))

		r.PostMapping("/customer", c.CreateExaCustomer)
		r.PUT("/customer", c.UpdateExaCustomer)
		r.DELETE("/customer", c.DeleteExaCustomer)
		r.GetMapping("/customer", c.GetExaCustomer)
		r.GetMapping("/customerList", c.GetExaCustomerList)
	})
}

type CustomerController struct {
	ExaCustomerService *service.ExaCustomerService `autowire:""`
}

// @Tags SysApi
// @Summary 创建客户
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ExaCustomer true "创建客户"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /customer/customer [post]
func (controller *CustomerController) CreateExaCustomer(webCtx SpringWeb.WebContext) {
	var cu model.ExaCustomer
	_ = webCtx.Bind(&cu)
	CustomerVerify := utils.Rules{
		"CustomerName":      {utils.NotEmpty()},
		"CustomerPhoneData": {utils.NotEmpty()},
	}
	CustomerVerifyErr := utils.Verify(cu, CustomerVerify)
	if CustomerVerifyErr != nil {
		response.FailWithMessage(CustomerVerifyErr.Error(), webCtx)
		return
	}
	claims := webCtx.Get("claims")
	waitUse := claims.(*request.CustomClaims)
	cu.SysUserID = waitUse.ID
	cu.SysUserAuthorityID = waitUse.AuthorityId
	err := controller.ExaCustomerService.CreateExaCustomer(cu)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("删除失败：%v", err), webCtx)
	} else {
		response.OkWithMessage("创建成功", webCtx)
	}
}

// @Tags SysApi
// @Summary 删除客户
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ExaCustomer true "删除客户"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /customer/customer [delete]
func (controller *CustomerController) DeleteExaCustomer(webCtx SpringWeb.WebContext) {
	var cu model.ExaCustomer
	_ = webCtx.Bind(&cu)
	CustomerVerify := utils.Rules{
		"ID": {utils.NotEmpty()},
	}
	CustomerVerifyErr := utils.Verify(cu.Model, CustomerVerify)
	if CustomerVerifyErr != nil {
		response.FailWithMessage(CustomerVerifyErr.Error(), webCtx)
		return
	}
	err := controller.ExaCustomerService.DeleteExaCustomer(cu)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("删除失败：%v", err), webCtx)
	} else {
		response.OkWithMessage("删除成功", webCtx)
	}
}

// @Tags SysApi
// @Summary 更新客户信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ExaCustomer true "创建客户"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /customer/customer [put]
func (controller *CustomerController) UpdateExaCustomer(webCtx SpringWeb.WebContext) {
	var cu model.ExaCustomer
	_ = webCtx.Bind(&cu)
	IdCustomerVerify := utils.Rules{
		"ID": {utils.NotEmpty()},
	}
	IdCustomerVerifyErr := utils.Verify(cu.Model, IdCustomerVerify)
	if IdCustomerVerifyErr != nil {
		response.FailWithMessage(IdCustomerVerifyErr.Error(), webCtx)
		return
	}
	CustomerVerify := utils.Rules{
		"CustomerName":      {utils.NotEmpty()},
		"CustomerPhoneData": {utils.NotEmpty()},
	}
	CustomerVerifyErr := utils.Verify(cu, CustomerVerify)
	if CustomerVerifyErr != nil {
		response.FailWithMessage(CustomerVerifyErr.Error(), webCtx)
		return
	}
	err := controller.ExaCustomerService.UpdateExaCustomer(&cu)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("更新失败：%v", err), webCtx)
	} else {
		response.OkWithMessage("更新成功", webCtx)
	}
}

// @Tags SysApi
// @Summary 获取单一客户信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ExaCustomer true "获取单一客户信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /customer/customer [get]
func (controller *CustomerController) GetExaCustomer(webCtx SpringWeb.WebContext) {
	var cu model.ExaCustomer
	_ = webCtx.Bind(&cu)
	IdCustomerVerify := utils.Rules{
		"ID": {utils.NotEmpty()},
	}
	IdCustomerVerifyErr := utils.Verify(cu.Model, IdCustomerVerify)
	if IdCustomerVerifyErr != nil {
		response.FailWithMessage(IdCustomerVerifyErr.Error(), webCtx)
		return
	}
	err, customer := controller.ExaCustomerService.GetExaCustomer(cu.ID)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("获取失败：%v", err), webCtx)
	} else {
		response.OkWithData(resp.ExaCustomerResponse{Customer: customer}, webCtx)
	}
}

// @Tags SysApi
// @Summary 获取权限客户列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "获取权限客户列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /customer/customerList [get]
func (controller *CustomerController) GetExaCustomerList(webCtx SpringWeb.WebContext) {
	claims := webCtx.Get("claims")
	waitUse := claims.(*request.CustomClaims)
	var pageInfo request.PageInfo
	_ = webCtx.Bind(&pageInfo)
	PageVerifyErr := utils.Verify(pageInfo, utils.CustomizeMap["PageVerify"])
	if PageVerifyErr != nil {
		response.FailWithMessage(PageVerifyErr.Error(), webCtx)
		return
	}
	err, customerList, total := controller.ExaCustomerService.GetCustomerInfoList(waitUse.AuthorityId, pageInfo)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("获取失败：%v", err), webCtx)
	} else {
		response.OkWithData(resp.PageResult{
			List:     customerList,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, webCtx)
	}
}
