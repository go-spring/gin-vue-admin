package router

import (
	"gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"

	"github.com/go-spring/go-spring-web/spring-gin"
	"github.com/go-spring/go-spring/spring-boot"
)

func init() {
	SpringBoot.RegisterBean(new(v1.ApiController)).Init(func(c *v1.ApiController) {

		r := SpringBoot.Route("/api",
			SpringGin.Filter(middleware.JWTAuth()),
			SpringGin.Filter(middleware.CasbinHandler()))

		r.POST("/createApi", SpringGin.Gin(c.CreateApi))
		r.POST("/deleteApi", SpringGin.Gin(c.DeleteApi))
		r.POST("/getApiList", SpringGin.Gin(c.GetApiList))
		r.POST("/getApiById", SpringGin.Gin(c.GetApiById))
		r.POST("/updateApi", SpringGin.Gin(c.UpdateApi))
		r.POST("/getAllApis", SpringGin.Gin(c.GetAllApis))
	})
}
