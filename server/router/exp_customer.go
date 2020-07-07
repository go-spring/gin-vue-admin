package router

import (
	"gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
	"github.com/go-spring/go-spring-web/spring-gin"
	"github.com/go-spring/go-spring/spring-boot"
)

func InitCustomerRouter(Router *gin.RouterGroup) {

	c := SpringBoot.Route("/customer",
		SpringGin.Filter(middleware.JWTAuth()),
		SpringGin.Filter(middleware.CasbinHandler()))

	c.POST("/customer", SpringGin.Gin(v1.CreateExaCustomer))
	c.PUT("/customer", SpringGin.Gin(v1.UpdateExaCustomer))
	c.DELETE("/customer", SpringGin.Gin(v1.DeleteExaCustomer))
	c.GET("/customer", SpringGin.Gin(v1.GetExaCustomer))
	c.GET("/customerList", SpringGin.Gin(v1.GetExaCustomerList))

	ApiRouter := Router.Group("customer").Use(middleware.JWTAuth()).Use(middleware.CasbinHandler())
	{
		ApiRouter.POST("customer", v1.CreateExaCustomer)     // 创建客户
		ApiRouter.PUT("customer", v1.UpdateExaCustomer)      // 更新客户
		ApiRouter.DELETE("customer", v1.DeleteExaCustomer)   // 删除客户
		ApiRouter.GET("customer", v1.GetExaCustomer)         // 获取单一客户信息
		ApiRouter.GET("customerList", v1.GetExaCustomerList) // 获取客户列表
	}
}
