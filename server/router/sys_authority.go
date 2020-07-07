package router

import (
	"gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"

	"github.com/go-spring/go-spring-web/spring-gin"
	"github.com/go-spring/go-spring/spring-boot"
)

func init() {

	a := SpringBoot.Route("/authority",
		SpringGin.Filter(middleware.JWTAuth()),
		SpringGin.Filter(middleware.CasbinHandler()))

	a.POST("/createAuthority", SpringGin.Gin(v1.CreateAuthority))
	a.POST("/deleteAuthority", SpringGin.Gin(v1.DeleteAuthority))
	a.PUT("/updateAuthority", SpringGin.Gin(v1.UpdateAuthority))
	a.POST("/copyAuthority", SpringGin.Gin(v1.CopyAuthority))
	a.POST("/getAuthorityList", SpringGin.Gin(v1.GetAuthorityList))
	a.POST("/setDataAuthority", SpringGin.Gin(v1.SetDataAuthority))
}
