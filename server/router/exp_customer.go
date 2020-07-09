package router

import (
	"gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"

	"github.com/go-spring/go-spring-web/spring-gin"
	"github.com/go-spring/go-spring/spring-boot"
)

func init() {
	SpringBoot.RegisterBean(new(v1.CustomerController)).Init(func(c *v1.CustomerController) {

		r := SpringBoot.Route("/customer",
			SpringGin.Filter(middleware.JWTAuth()),
			SpringGin.Filter(middleware.CasbinHandler()))

		r.POST("/customer", SpringGin.Gin(c.CreateExaCustomer))
		r.PUT("/customer", SpringGin.Gin(c.UpdateExaCustomer))
		r.DELETE("/customer", SpringGin.Gin(c.DeleteExaCustomer))
		r.GET("/customer", SpringGin.Gin(c.GetExaCustomer))
		r.GET("/customerList", SpringGin.Gin(c.GetExaCustomerList))
	})
}
