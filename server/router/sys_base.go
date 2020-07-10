package router

import (
	"gin-vue-admin/api/v1"

	"github.com/go-spring/go-spring/spring-boot"
)

func init() {
	SpringBoot.RegisterBean(new(v1.BaseController)).Init(func(c *v1.BaseController) {

		r := SpringBoot.Route("/base")

		r.PostMapping("/login", c.Login)
		r.PostMapping("/captcha", c.Captcha)
		r.PostMapping("/register", c.Register)
		r.GetMapping("/captcha/:captchaId", c.CaptchaImg)
	})
}
