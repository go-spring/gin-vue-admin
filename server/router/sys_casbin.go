package router

import (
	"gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"

	"github.com/go-spring/go-spring-web/spring-gin"
	"github.com/go-spring/go-spring/spring-boot"
)

func init() {

	c := SpringBoot.Route("/casbin",
		SpringGin.Filter(middleware.JWTAuth()),
		SpringGin.Filter(middleware.CasbinHandler()))

	c.POST("/updateCasbin", SpringGin.Gin(v1.UpdateCasbin))
	c.POST("/getPolicyPathByAuthorityId", SpringGin.Gin(v1.GetPolicyPathByAuthorityId))
	c.GET("/casbinTest/:pathParam", SpringGin.Gin(v1.CasbinTest))
}
