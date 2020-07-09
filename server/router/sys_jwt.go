package router

import (
	"gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"
	"github.com/go-spring/go-spring-web/spring-gin"
	"github.com/go-spring/go-spring/spring-boot"
)

func init() {

	SpringBoot.RegisterBean(new(v1.JwTController)).Init(func(controller *v1.JwTController) {
		a := SpringBoot.Route("/jwt",
			SpringGin.Filter(middleware.JWTAuth()),
			SpringGin.Filter(middleware.CasbinHandler()))

		a.POST("/jsonInBlacklist", SpringGin.Gin(controller.JsonInBlacklist))
	})

}
