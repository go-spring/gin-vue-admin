package controller

import (
	"fmt"

	"gin-vue-admin/filter"
	"gin-vue-admin/global/response"
	"gin-vue-admin/model/request"
	resp "gin-vue-admin/model/response"
	"gin-vue-admin/service"
	"gin-vue-admin/utils"

	"github.com/gin-gonic/gin"
	"github.com/go-spring/spring-boot"
	"github.com/go-spring/spring-web"
)

func init() {
	SpringBoot.RegisterBean(new(CasbinController)).Init(func(c *CasbinController) {

		r := SpringBoot.Route("/casbin",
			SpringBoot.FilterBean((*filter.TraceFilter)(nil)),
			SpringBoot.FilterBean((*filter.JwtFilter)(nil)),
			SpringBoot.FilterBean((*filter.CasbinRcbaFilter)(nil)))

		r.PostMapping("/updateCasbin", c.UpdateCasbin)
		r.PostMapping("/getPolicyPathByAuthorityId", c.GetPolicyPathByAuthorityId)
		r.GetMapping("/casbinTest/:pathParam", c.CasbinTest)
	})
}

type CasbinController struct {
	SysCasbinService *service.SysCasbinService `autowire:""`
}

// @Tags casbin
// @Summary 更改角色api权限
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.CasbinInReceive true "更改角色api权限"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /casbin/UpdateCasbin [post]
func (controller *CasbinController) UpdateCasbin(webCtx SpringWeb.WebContext) {
	var cmr request.CasbinInReceive
	_ = webCtx.Bind(&cmr)
	AuthorityIdVerifyErr := utils.Verify(cmr, utils.CustomizeMap["AuthorityIdVerify"])
	if AuthorityIdVerifyErr != nil {
		response.FailWithMessage(AuthorityIdVerifyErr.Error(), webCtx)
		return
	}
	err := controller.SysCasbinService.UpdateCasbin(cmr.AuthorityId, cmr.CasbinInfos)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("添加规则失败，%v", err), webCtx)
	} else {
		response.OkWithMessage("添加规则成功", webCtx)
	}
}

// @Tags casbin
// @Summary 获取权限列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.CasbinInReceive true "获取权限列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /casbin/getPolicyPathByAuthorityId [post]
func (controller *CasbinController) GetPolicyPathByAuthorityId(webCtx SpringWeb.WebContext) {
	var cmr request.CasbinInReceive
	_ = webCtx.Bind(&cmr)
	AuthorityIdVerifyErr := utils.Verify(cmr, utils.CustomizeMap["AuthorityIdVerify"])
	if AuthorityIdVerifyErr != nil {
		response.FailWithMessage(AuthorityIdVerifyErr.Error(), webCtx)
		return
	}
	paths := controller.SysCasbinService.GetPolicyPathByAuthorityId(cmr.AuthorityId)
	response.OkWithData(resp.PolicyPathResponse{Paths: paths}, webCtx)
}

// @Tags casbin
// @Summary casb RBAC RESTFUL测试路由
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.CasbinInReceive true "获取权限列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /casbin/CasbinTest [get]
func (controller *CasbinController) CasbinTest(webCtx SpringWeb.WebContext) {
	// 测试restful以及占位符代码  随意书写
	pathParam := webCtx.PathParam("pathParam")
	query := webCtx.QueryParam("query")
	response.OkDetailed(gin.H{"pathParam": pathParam, "query": query}, "获取规则成功", webCtx)
}
