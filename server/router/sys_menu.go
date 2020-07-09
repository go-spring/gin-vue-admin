package router

import (
	"gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"

	"github.com/go-spring/go-spring-web/spring-gin"
	"github.com/go-spring/go-spring/spring-boot"
)

func init() {
	SpringBoot.RegisterBean(new(v1.MenuController)).Init(func(c *v1.MenuController) {

		r := SpringBoot.Route("/menu",
			SpringGin.Filter(middleware.JWTAuth()),
			SpringGin.Filter(middleware.CasbinHandler()))

		r.POST("/getMenu", SpringGin.Gin(c.GetMenu))
		r.POST("/getMenuList", SpringGin.Gin(c.GetMenuList))
		r.POST("/addBaseMenu", SpringGin.Gin(c.AddBaseMenu))
		r.POST("/getBaseMenuTree", SpringGin.Gin(c.GetBaseMenuTree))
		r.POST("/addMenuAuthority", SpringGin.Gin(c.AddMenuAuthority))
		r.POST("/getMenuAuthority", SpringGin.Gin(c.GetMenuAuthority))
		r.POST("/deleteBaseMenu", SpringGin.Gin(c.DeleteBaseMenu))
		r.POST("/updateBaseMenu", SpringGin.Gin(c.UpdateBaseMenu))
		r.POST("/getBaseMenuById", SpringGin.Gin(c.GetBaseMenuById))
	})
}
