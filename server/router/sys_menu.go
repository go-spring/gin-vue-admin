package router

import (
	"gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"

	"github.com/go-spring/go-spring-web/spring-gin"
	"github.com/go-spring/go-spring/spring-boot"
)

func init() {

	m := SpringBoot.Route("/menu",
		SpringGin.Filter(middleware.JWTAuth()),
		SpringGin.Filter(middleware.CasbinHandler()))

	m.POST("/getMenu", SpringGin.Gin(v1.GetMenu))
	m.POST("/getMenuList", SpringGin.Gin(v1.GetMenuList))
	m.POST("/addBaseMenu", SpringGin.Gin(v1.AddBaseMenu))
	m.POST("/getBaseMenuTree", SpringGin.Gin(v1.GetBaseMenuTree))
	m.POST("/addMenuAuthority", SpringGin.Gin(v1.AddMenuAuthority))
	m.POST("/getMenuAuthority", SpringGin.Gin(v1.GetMenuAuthority))
	m.POST("/deleteBaseMenu", SpringGin.Gin(v1.DeleteBaseMenu))
	m.POST("/updateBaseMenu", SpringGin.Gin(v1.UpdateBaseMenu))
	m.POST("/getBaseMenuById", SpringGin.Gin(v1.GetBaseMenuById))
}
