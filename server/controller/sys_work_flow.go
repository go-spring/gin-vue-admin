package controller

import (
	"fmt"

	"gin-vue-admin/filter"
	"gin-vue-admin/global/response"
	"gin-vue-admin/model"
	"gin-vue-admin/service"
	"gin-vue-admin/utils"

	"github.com/go-spring/spring-boot"
	"github.com/go-spring/spring-web"
)

func init() {
	SpringBoot.RegisterBean(new(WorkFlowController)).Init(func(c *WorkFlowController) {

		r := SpringBoot.Route("/workflow",
			SpringBoot.FilterBean((*filter.TraceFilter)(nil)),
			SpringBoot.FilterBean((*filter.JwtFilter)(nil)),
			SpringBoot.FilterBean((*filter.CasbinRcbaFilter)(nil)))

		r.PostMapping("/createWorkFlow", c.CreateWorkFlow)
	})
}

type WorkFlowController struct {
	SysWorkflowService *service.SysWorkflowService `autowire:""`
}

// @Tags workflow
// @Summary 注册工作流
// @Produce  application/json
// @Param data body model.SysWorkflow true "注册工作流接口"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"注册成功"}"
// @Router /workflow/createWorkFlow [post]
func (controller *WorkFlowController) CreateWorkFlow(webCtx SpringWeb.WebContext) {
	var wk model.SysWorkflow
	_ = webCtx.Bind(&wk)
	WKVerify := utils.Rules{
		"WorkflowNickName":    {utils.NotEmpty()},
		"WorkflowName":        {utils.NotEmpty()},
		"WorkflowDescription": {utils.NotEmpty()},
		"WorkflowStepInfo":    {utils.NotEmpty()},
	}
	WKVerifyErr := utils.Verify(wk, WKVerify)
	if WKVerifyErr != nil {
		response.FailWithMessage(WKVerifyErr.Error(), webCtx)
		return
	}
	err := controller.SysWorkflowService.Create(wk)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("获取失败：%v", err), webCtx)
	} else {
		response.OkWithMessage("获取成功", webCtx)
	}
}
