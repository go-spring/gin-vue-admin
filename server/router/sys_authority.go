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

		r.POST("/createAuthority", SpringGin.Gin(c.CreateAuthority))
		r.POST("/deleteAuthority", SpringGin.Gin(c.DeleteAuthority))
		r.PUT("/updateAuthority", SpringGin.Gin(c.UpdateAuthority))
		r.POST("/copyAuthority", SpringGin.Gin(c.CopyAuthority))
		r.POST("/getAuthorityList", SpringGin.Gin(c.GetAuthorityList))
		r.POST("/setDataAuthority", SpringGin.Gin(c.SetDataAuthority))
	})
}
