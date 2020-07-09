package router

import (
	"gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"

	"github.com/go-spring/go-spring-web/spring-gin"
	"github.com/go-spring/go-spring/spring-boot"
)

func init() {

	SpringBoot.RegisterBean(new(v1.CustomerController)).Init(func(controller *v1.CustomerController) {
		c := SpringBoot.Route("/customer",
			SpringGin.Filter(middleware.JWTAuth()),
			SpringGin.Filter(middleware.CasbinHandler()))

		c.POST("/customer", SpringGin.Gin(controller.CreateExaCustomer))
		c.PUT("/customer", SpringGin.Gin(controller.UpdateExaCustomer))
		c.DELETE("/customer", SpringGin.Gin(controller.DeleteExaCustomer))
		c.GET("/customer", SpringGin.Gin(controller.GetExaCustomer))
		c.GET("/customerList", SpringGin.Gin(controller.GetExaCustomerList))
	})

}
