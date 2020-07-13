package v1

import (
	"fmt"

	"gin-vue-admin/global/response"
	"gin-vue-admin/model"
	"gin-vue-admin/service"

	"github.com/gin-gonic/gin"
	"github.com/go-spring/go-spring-web/spring-web"
)

type JwTController struct {
}

// @Tags jwt
// @Summary jwt加入黑名单
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"拉黑成功"}"
// @Router /jwt/jsonInBlacklist [post]
func (controller *JwTController) JsonInBlacklist(webCtx SpringWeb.WebContext) {
	c := webCtx.NativeContext().(*gin.Context)

	token := c.Request.Header.Get("x-token")
	modelJwt := model.JwtBlacklist{
		Jwt: token,
	}
	err := service.JsonInBlacklist(modelJwt)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("jwt作废失败，%v", err), webCtx)
	} else {
		response.OkWithMessage("jwt作废成功", webCtx)
	}
}
