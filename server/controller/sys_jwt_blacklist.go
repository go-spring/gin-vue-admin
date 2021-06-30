package controller

import (
	"fmt"

	"gin-vue-admin/filter"
	"gin-vue-admin/global/response"
	"gin-vue-admin/model"
	"gin-vue-admin/service"

	"github.com/go-spring/spring-boot"
	"github.com/go-spring/spring-web"
)

func init() {
	SpringBoot.RegisterBean(new(JwTController)).Init(func(c *JwTController) {

		r := SpringBoot.Route("/jwt",
			SpringBoot.FilterBean((*filter.TraceFilter)(nil)),
			SpringBoot.FilterBean((*filter.JwtFilter)(nil)),
			SpringBoot.FilterBean((*filter.CasbinRcbaFilter)(nil)))

		r.PostMapping("/jsonInBlacklist", c.JsonInBlacklist)
	})
}

type JwTController struct {
	JwtBlackListService *service.JwtBlackListService `autowire:""`
}

// @Tags jwt
// @Summary jwt加入黑名单
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"拉黑成功"}"
// @Router /jwt/jsonInBlacklist [post]
func (controller *JwTController) JsonInBlacklist(webCtx SpringWeb.WebContext) {
	token := webCtx.GetHeader("x-token")
	modelJwt := model.JwtBlacklist{
		Jwt: token,
	}
	err := controller.JwtBlackListService.JsonInBlacklist(modelJwt)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("jwt作废失败，%v", err), webCtx)
	} else {
		response.OkWithMessage("jwt作废成功", webCtx)
	}
}
