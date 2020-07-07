package router

import (
	"gin-vue-admin/api/v1"

	"github.com/go-spring/go-spring-web/spring-gin"
	"github.com/go-spring/go-spring/spring-boot"
)

func init() {

	r := SpringBoot.Route("/base")

	r.HandlePost("/login", SpringGin.Gin(v1.Login))
	r.HandlePost("/captcha", SpringGin.Gin(v1.Captcha))
	r.HandlePost("/register", SpringGin.Gin(v1.Register))
	r.HandleGet("/captcha/:captchaId", SpringGin.Gin(v1.CaptchaImg))
}
