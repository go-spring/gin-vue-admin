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

		r.PostMapping("/customer", c.CreateExaCustomer)
		r.PUT("/customer", c.UpdateExaCustomer)
		r.DELETE("/customer", c.DeleteExaCustomer)
		r.GetMapping("/customer", c.GetExaCustomer)
		r.GetMapping("/customerList", c.GetExaCustomerList)
	})
}
