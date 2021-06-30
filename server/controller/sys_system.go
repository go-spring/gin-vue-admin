package controller

import (
	"fmt"

	"gin-vue-admin/filter"
	"gin-vue-admin/global/response"
	"gin-vue-admin/model"
	resp "gin-vue-admin/model/response"
	"gin-vue-admin/service"

	"github.com/go-spring/spring-boot"
	"github.com/go-spring/spring-web"
)

func init() {
	SpringBoot.RegisterBean(new(SystemController)).Init(func(c *SystemController) {

		r := SpringBoot.Route("/system",
			SpringBoot.FilterBean((*filter.TraceFilter)(nil)),
			SpringBoot.FilterBean((*filter.JwtFilter)(nil)),
			SpringBoot.FilterBean((*filter.CasbinRcbaFilter)(nil)))

		r.PostMapping("/getSystemConfig", c.GetSystemConfig)
		r.PostMapping("/setSystemConfig", c.SetSystemConfig)
	})
}

type SystemController struct {
	SysSystemService *service.SysSystemService `autowire:""`
}

// @Tags system
// @Summary 获取配置文件内容
// @Security ApiKeyAuth
// @Produce  application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"返回成功"}"
// @Router /system/getSystemConfig [post]
func (controller *SystemController) GetSystemConfig(webCtx SpringWeb.WebContext) {
	err, config := controller.SysSystemService.GetSystemConfig()
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("获取失败，%v", err), webCtx)
	} else {
		response.OkWithData(resp.SysConfigResponse{Config: config}, webCtx)
	}
}

// @Tags system
// @Summary 设置配置文件内容
// @Security ApiKeyAuth
// @Produce  application/json
// @Param data body model.System true "设置配置文件内容"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"返回成功"}"
// @Router /system/setSystemConfig [post]
func (controller *SystemController) SetSystemConfig(webCtx SpringWeb.WebContext) {
	var sys model.System
	_ = webCtx.Bind(&sys)
	err := controller.SysSystemService.SetSystemConfig(sys)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("设置失败，%v", err), webCtx)
	} else {
		response.OkWithData("设置成功", webCtx)
	}
}

// 本方法开发中 开发者windows系统 缺少linux系统所需的包 因此搁置
// @Tags system
// @Summary 设置配置文件内容
// @Security ApiKeyAuth
// @Produce  application/json
// @Param data body model.System true "设置配置文件内容"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"返回成功"}"
// @Router /system/ReloadSystem [post]
func (controller *SystemController) ReloadSystem(webCtx SpringWeb.WebContext) {
	var sys model.System
	_ = webCtx.Bind(&sys)
	err := controller.SysSystemService.SetSystemConfig(sys)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("设置失败，%v", err), webCtx)
	} else {
		response.OkWithMessage("设置成功", webCtx)
	}
}
