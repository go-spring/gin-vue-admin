package router

import (
	"gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"

	"github.com/go-spring/go-spring-web/spring-gin"
	"github.com/go-spring/go-spring/spring-boot"
)

func init() {

	u := SpringBoot.Route("/system",
		SpringGin.Filter(middleware.JWTAuth()),
		SpringGin.Filter(middleware.CasbinHandler()))

	u.POST("/getSystemConfig", SpringGin.Gin(v1.GetSystemConfig))
	u.POST("/setSystemConfig", SpringGin.Gin(v1.SetSystemConfig))
}
