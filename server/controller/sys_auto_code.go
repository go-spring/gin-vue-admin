package controller

import (
	"fmt"
	"net/url"
	"os"

	"gin-vue-admin/global/response"
	"gin-vue-admin/middleware"
	"gin-vue-admin/model"
	"gin-vue-admin/service"
	"gin-vue-admin/utils"

	"github.com/go-spring/go-spring-web/spring-gin"
	"github.com/go-spring/go-spring-web/spring-web"
	"github.com/go-spring/go-spring/spring-boot"
)

func init() {
	SpringBoot.RegisterBean(new(AutoCodeController)).Init(func(c *AutoCodeController) {

		r := SpringBoot.Route("/autoCode",
			SpringGin.Filter(middleware.JWTAuth()),
			SpringGin.Filter(middleware.CasbinHandler()))

		r.PostMapping("/createTemp", c.CreateTemp)
	})
}

type AutoCodeController struct {
	SysApiService      *service.SysApiService      `autowire:""`
	SysAutoCodeService *service.SysAutoCodeService `autowire:""`
}

// @Tags SysApi
// @Summary 自动代码模板
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AutoCodeStruct true "创建自动代码"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /autoCode/createTemp [post]
func (controller *AutoCodeController) CreateTemp(webCtx SpringWeb.WebContext) {
	var a model.AutoCodeStruct
	_ = webCtx.Bind(&a)
	AutoCodeVerify := utils.Rules{
		"Abbreviation": {utils.NotEmpty()},
		"StructName":   {utils.NotEmpty()},
		"PackageName":  {utils.NotEmpty()},
		"Fields":       {utils.NotEmpty()},
	}
	WKVerifyErr := utils.Verify(a, AutoCodeVerify)
	if WKVerifyErr != nil {
		response.FailWithMessage(WKVerifyErr.Error(), webCtx)
		return
	}
	if a.AutoCreateApiToSql {
		apiList := [5]model.SysApi{
			{
				Path:        "/" + a.Abbreviation + "/" + "create" + a.StructName,
				Description: "新增" + a.Description,
				ApiGroup:    a.Abbreviation,
				Method:      "POST",
			},
			{
				Path:        "/" + a.Abbreviation + "/" + "delete" + a.StructName,
				Description: "删除" + a.Description,
				ApiGroup:    a.Abbreviation,
				Method:      "DELETE",
			},
			{
				Path:        "/" + a.Abbreviation + "/" + "update" + a.StructName,
				Description: "更新" + a.Description,
				ApiGroup:    a.Abbreviation,
				Method:      "PUT",
			},
			{
				Path:        "/" + a.Abbreviation + "/" + "find" + a.StructName,
				Description: "根据ID获取" + a.Description,
				ApiGroup:    a.Abbreviation,
				Method:      "GET",
			},
			{
				Path:        "/" + a.Abbreviation + "/" + "get" + a.StructName + "List",
				Description: "获取" + a.Description + "列表",
				ApiGroup:    a.Abbreviation,
				Method:      "GET",
			},
		}
		for _, v := range apiList {
			errC := controller.SysApiService.CreateApi(v)
			if errC != nil {
				webCtx.Header("success", "false")
				webCtx.Header("msg", url.QueryEscape(fmt.Sprintf("自动化创建失败，%v，请自行清空垃圾数据", errC)))
				return
			}
		}
	}
	err := controller.SysAutoCodeService.CreateTemp(a)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("创建失败，%v", err), webCtx)
		os.Remove("./ginvueadmin.zip")
	} else {
		webCtx.Header("Content-Disposition", fmt.Sprintf("attachment; filename=%s", "ginvueadmin.zip")) // fmt.Sprintf("attachment; filename=%s", filename)对下载的文件重命名
		webCtx.Header("Content-Type", "application/json")
		webCtx.Header("success", "true")
		webCtx.File("./ginvueadmin.zip")
		os.Remove("./ginvueadmin.zip")
	}
}
