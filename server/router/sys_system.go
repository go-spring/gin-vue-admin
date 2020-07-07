package router

import (
	"gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
	"github.com/go-spring/go-spring-web/spring-gin"
	"github.com/go-spring/go-spring/spring-boot"
)

func InitSystemRouter(Router *gin.RouterGroup) {

	u := SpringBoot.Route("/system",
		SpringGin.Filter(middleware.JWTAuth()),
		SpringGin.Filter(middleware.CasbinHandler()))

	u.POST("/getSystemConfig", SpringGin.Gin(v1.GetSystemConfig))
	u.POST("/setSystemConfig", SpringGin.Gin(v1.SetSystemConfig))

	UserRouter := Router.Group("system").Use(middleware.JWTAuth()).Use(middleware.CasbinHandler())
	{
		UserRouter.POST("getSystemConfig", v1.GetSystemConfig) // 获取配置文件内容
		UserRouter.POST("setSystemConfig", v1.SetSystemConfig) // 设置配置文件内容
	}
}
