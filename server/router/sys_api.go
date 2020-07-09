package router

import (
	"gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"
	"github.com/go-spring/go-spring-web/spring-gin"
	"github.com/go-spring/go-spring/spring-boot"
)

func init() {

	SpringBoot.RegisterBean(new(v1.ApiController)).Init(func(controller *v1.ApiController) {
		a := SpringBoot.Route("/api",
			SpringGin.Filter(middleware.JWTAuth()),
			SpringGin.Filter(middleware.CasbinHandler()))

		a.POST("/createApi", SpringGin.Gin(controller.CreateApi))
		a.POST("/deleteApi", SpringGin.Gin(controller.DeleteApi))
		a.POST("/getApiList", SpringGin.Gin(controller.GetApiList))
		a.POST("/getApiById", SpringGin.Gin(controller.GetApiById))
		a.POST("/updateApi", SpringGin.Gin(controller.UpdateApi))
		a.POST("/getAllApis", SpringGin.Gin(controller.GetAllApis))
	})

}
