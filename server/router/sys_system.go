package router

import (
	"gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"

	"github.com/go-spring/go-spring-web/spring-gin"
	"github.com/go-spring/go-spring/spring-boot"
)

func init() {

	SpringBoot.RegisterBean(new(v1.SystemController)).Init(func(controller *v1.SystemController) {
		u := SpringBoot.Route("/system",
			SpringGin.Filter(middleware.JWTAuth()),
			SpringGin.Filter(middleware.CasbinHandler()))

		u.POST("/getSystemConfig", SpringGin.Gin(controller.GetSystemConfig))
		u.POST("/setSystemConfig", SpringGin.Gin(controller.SetSystemConfig))
	})

}
