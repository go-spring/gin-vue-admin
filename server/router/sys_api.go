package router

import (
	"gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"

	"github.com/go-spring/go-spring-web/spring-gin"
	"github.com/go-spring/go-spring/spring-boot"
)

func init() {

	a := SpringBoot.Route("/api",
		SpringGin.Filter(middleware.JWTAuth()),
		SpringGin.Filter(middleware.CasbinHandler()))

	a.POST("/createApi", SpringGin.Gin(v1.CreateApi))
	a.POST("/deleteApi", SpringGin.Gin(v1.DeleteApi))
	a.POST("/getApiList", SpringGin.Gin(v1.GetApiList))
	a.POST("/getApiById", SpringGin.Gin(v1.GetApiById))
	a.POST("/updateApi", SpringGin.Gin(v1.UpdateApi))
	a.POST("/getAllApis", SpringGin.Gin(v1.GetAllApis))
}
