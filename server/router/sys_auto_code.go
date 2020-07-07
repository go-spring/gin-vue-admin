package router

import (
	"gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
	"github.com/go-spring/go-spring-web/spring-gin"
	"github.com/go-spring/go-spring/spring-boot"
)

func InitAutoCodeRouter(Router *gin.RouterGroup) {

	ac := SpringBoot.Route("/autoCode",
		SpringGin.Filter(middleware.JWTAuth()),
		SpringGin.Filter(middleware.CasbinHandler()))

	ac.POST("/createTemp", SpringGin.Gin(v1.CreateTemp))

	AutoCodeRouter := Router.Group("autoCode").Use(middleware.JWTAuth()).Use(middleware.CasbinHandler())
	{
		AutoCodeRouter.POST("createTemp", v1.CreateTemp) // 创建自动化代码
	}
}
