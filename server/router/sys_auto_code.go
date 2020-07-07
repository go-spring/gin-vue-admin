package router

import (
	"gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"

	"github.com/go-spring/go-spring-web/spring-gin"
	"github.com/go-spring/go-spring/spring-boot"
)

func init() {

	ac := SpringBoot.Route("/autoCode",
		SpringGin.Filter(middleware.JWTAuth()),
		SpringGin.Filter(middleware.CasbinHandler()))

	ac.POST("/createTemp", SpringGin.Gin(v1.CreateTemp))
}
