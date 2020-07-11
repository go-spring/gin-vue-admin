package router

import (
	"gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"

	"github.com/go-spring/go-spring-web/spring-gin"
	"github.com/go-spring/go-spring/spring-boot"
)

func init() {
	SpringBoot.RegisterBean(new(v1.CasbinController)).Init(func(c *v1.CasbinController) {

		r := SpringBoot.Route("/casbin",
			SpringGin.Filter(middleware.JWTAuth()),
			SpringGin.Filter(middleware.CasbinHandler()))

		r.PostMapping("/updateCasbin", c.UpdateCasbin)
		r.PostMapping("/getPolicyPathByAuthorityId", c.GetPolicyPathByAuthorityId)
		r.GetMapping("/casbinTest/:pathParam", c.CasbinTest)
	})
}
