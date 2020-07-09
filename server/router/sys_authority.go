package router

import (
	"gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"
	"github.com/go-spring/go-spring-web/spring-gin"
	"github.com/go-spring/go-spring/spring-boot"
)

func init() {

	SpringBoot.RegisterBean(new(v1.AuthorityController)).Init(func(controller *v1.AuthorityController) {
		a := SpringBoot.Route("/authority",
			SpringGin.Filter(middleware.JWTAuth()),
			SpringGin.Filter(middleware.CasbinHandler()))

		a.POST("/createAuthority", SpringGin.Gin(controller.CreateAuthority))
		a.POST("/deleteAuthority", SpringGin.Gin(controller.DeleteAuthority))
		a.PUT("/updateAuthority", SpringGin.Gin(controller.UpdateAuthority))
		a.POST("/copyAuthority", SpringGin.Gin(controller.CopyAuthority))
		a.POST("/getAuthorityList", SpringGin.Gin(controller.GetAuthorityList))
		a.POST("/setDataAuthority", SpringGin.Gin(controller.SetDataAuthority))
	})

}
