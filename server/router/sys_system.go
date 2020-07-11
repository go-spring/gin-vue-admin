package router

import (
	"gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"

	"github.com/go-spring/go-spring-web/spring-gin"
	"github.com/go-spring/go-spring/spring-boot"
)

func init() {
	SpringBoot.RegisterBean(new(v1.SystemController)).Init(func(c *v1.SystemController) {

		r := SpringBoot.Route("/system",
			SpringGin.Filter(middleware.JWTAuth()),
			SpringGin.Filter(middleware.CasbinHandler()))

		r.PostMapping("/getSystemConfig", c.GetSystemConfig)
		r.PostMapping("/setSystemConfig", c.SetSystemConfig)
	})
}
