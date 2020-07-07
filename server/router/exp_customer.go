package router

import (
	"gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"

	"github.com/go-spring/go-spring-web/spring-gin"
	"github.com/go-spring/go-spring/spring-boot"
)

func init() {

	c := SpringBoot.Route("/customer",
		SpringGin.Filter(middleware.JWTAuth()),
		SpringGin.Filter(middleware.CasbinHandler()))

	c.POST("/customer", SpringGin.Gin(v1.CreateExaCustomer))
	c.PUT("/customer", SpringGin.Gin(v1.UpdateExaCustomer))
	c.DELETE("/customer", SpringGin.Gin(v1.DeleteExaCustomer))
	c.GET("/customer", SpringGin.Gin(v1.GetExaCustomer))
	c.GET("/customerList", SpringGin.Gin(v1.GetExaCustomerList))
}
