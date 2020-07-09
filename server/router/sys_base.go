package router

import (
	"gin-vue-admin/api/v1"

	"github.com/go-spring/go-spring-web/spring-gin"
	"github.com/go-spring/go-spring/spring-boot"
)

func init() {
	SpringBoot.RegisterBean(new(v1.Base)).Init(func(base *v1.Base) {
		r := SpringBoot.Route("/base")
		r.HandlePost("/login", SpringGin.Gin(v1.Login))
		r.HandlePost("/captcha", SpringGin.Gin(v1.Captcha))
		r.HandlePost("/register", SpringGin.Gin(base.Register))
		r.HandleGet("/captcha/:captchaId", SpringGin.Gin(v1.CaptchaImg))
	})
}
