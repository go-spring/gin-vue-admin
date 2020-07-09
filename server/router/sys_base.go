package router

import (
	"gin-vue-admin/api/v1"

	"github.com/go-spring/go-spring-web/spring-gin"
	"github.com/go-spring/go-spring/spring-boot"
)

func init() {
	SpringBoot.RegisterBean(new(v1.BaseController)).Init(func(c *v1.BaseController) {

		r := SpringBoot.Route("/base")

		r.PostMapping("/login", c.Login)
		r.HandlePost("/captcha", SpringGin.Gin(c.Captcha))
		r.HandlePost("/register", SpringGin.Gin(c.Register))
		r.HandleGet("/captcha/:captchaId", SpringGin.Gin(c.CaptchaImg))
	})
}
