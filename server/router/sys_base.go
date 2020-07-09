package router

import (
	"gin-vue-admin/api/v1"

	"github.com/go-spring/go-spring-web/spring-gin"
	"github.com/go-spring/go-spring/spring-boot"
)

func init() {

	SpringBoot.RegisterBean(new(v1.BaseController)).Init(func(controller *v1.BaseController) {
		r := SpringBoot.Route("/base")

		r.HandlePost("/login", SpringGin.Gin(controller.Login))
		r.HandlePost("/captcha", SpringGin.Gin(controller.Captcha))
		r.HandlePost("/register", SpringGin.Gin(controller.Register))
		r.HandleGet("/captcha/:captchaId", SpringGin.Gin(controller.CaptchaImg))
	})

}
