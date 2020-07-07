package router

import (
	"gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
	"github.com/go-spring/go-spring-web/spring-gin"
	"github.com/go-spring/go-spring/spring-boot"
)

func InitApiRouter(Router *gin.RouterGroup) {

	a := SpringBoot.Route("/api",
		SpringGin.Filter(middleware.JWTAuth()),
		SpringGin.Filter(middleware.CasbinHandler()))

	a.POST("/createApi", SpringGin.Gin(v1.CreateApi))
	a.POST("/deleteApi", SpringGin.Gin(v1.DeleteApi))
	a.POST("/getApiList", SpringGin.Gin(v1.GetApiList))
	a.POST("/getApiById", SpringGin.Gin(v1.GetApiById))
	a.POST("/updateApi", SpringGin.Gin(v1.UpdateApi))
	a.POST("/getAllApis", SpringGin.Gin(v1.GetAllApis))

	ApiRouter := Router.Group("api").Use(middleware.JWTAuth()).Use(middleware.CasbinHandler())
	{
		ApiRouter.POST("createApi", v1.CreateApi)   // 创建Api
		ApiRouter.POST("deleteApi", v1.DeleteApi)   // 删除Api
		ApiRouter.POST("getApiList", v1.GetApiList) // 获取Api列表
		ApiRouter.POST("getApiById", v1.GetApiById) // 获取单条Api消息
		ApiRouter.POST("updateApi", v1.UpdateApi)   // 更新api
		ApiRouter.POST("getAllApis", v1.GetAllApis) // 获取所有api
	}
}
