package controller

import (
	"gin-vue-admin/global/response"
	resp "gin-vue-admin/model/response"
	"github.com/dchest/captcha"
	"github.com/go-spring/spring-boot"

	"github.com/go-spring/spring-web"
)

// @Tags base
// @Summary 生成验证码
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /base/captcha [post]
func (controller *BaseController) Captcha(webCtx SpringWeb.WebContext) {
	captchaId := captcha.NewLen(int(SpringBoot.GetIntProperty("captcha.key-long")))
	response.OkDetailed(resp.SysCaptchaResponse{
		CaptchaId: captchaId,
		PicPath:   "/base/captcha/" + captchaId + ".png",
	}, "验证码获取成功", webCtx)
}

// @Tags base
// @Summary 生成验证码图片路径
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /base/captcha/:captchaId [get]
func (controller *BaseController) CaptchaImg(webCtx SpringWeb.WebContext) {
	controller.GinCaptchaService.GinCaptchaServeHTTP(webCtx.ResponseWriter(), webCtx.Request())
}
