package router

import (
	"gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"

	"github.com/go-spring/go-spring-web/spring-gin"
	"github.com/go-spring/go-spring/spring-boot"
)

func init() {
	SpringBoot.RegisterBean(new(v1.AuthorityController)).Init(func(c *v1.AuthorityController) {

		r := SpringBoot.Route("/authority",
			SpringGin.Filter(middleware.JWTAuth()),
			SpringGin.Filter(middleware.CasbinHandler()))

		r.PostMapping("/createAuthority", c.CreateAuthority)
		r.PostMapping("/deleteAuthority", c.DeleteAuthority)
		r.PUT("/updateAuthority", c.UpdateAuthority)
		r.PostMapping("/copyAuthority", c.CopyAuthority)
		r.PostMapping("/getAuthorityList", c.GetAuthorityList)
		r.PostMapping("/setDataAuthority", c.SetDataAuthority)
	})
}
