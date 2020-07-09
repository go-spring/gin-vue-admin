package router

import (
	"gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"

	"github.com/go-spring/go-spring-web/spring-gin"
	"github.com/go-spring/go-spring/spring-boot"
)

func init() {

	SpringBoot.RegisterBean(new(v1.CasbinController)).Init(func(controller *v1.CasbinController) {
		c := SpringBoot.Route("/casbin",
			SpringGin.Filter(middleware.JWTAuth()),
			SpringGin.Filter(middleware.CasbinHandler()))

		c.POST("/updateCasbin", SpringGin.Gin(controller.UpdateCasbin))
		c.POST("/getPolicyPathByAuthorityId", SpringGin.Gin(controller.GetPolicyPathByAuthorityId))
		c.GET("/casbinTest/:pathParam", SpringGin.Gin(controller.CasbinTest))
	})

}
