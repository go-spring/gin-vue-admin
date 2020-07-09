package router

import (
	"gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"

	"github.com/go-spring/go-spring-web/spring-gin"
	"github.com/go-spring/go-spring/spring-boot"
)

func init() {

	SpringBoot.RegisterBean(new(v1.MenuController)).Init(func(controller *v1.MenuController) {
		m := SpringBoot.Route("/menu",
			SpringGin.Filter(middleware.JWTAuth()),
			SpringGin.Filter(middleware.CasbinHandler()))

		m.POST("/getMenu", SpringGin.Gin(controller.GetMenu))
		m.POST("/getMenuList", SpringGin.Gin(controller.GetMenuList))
		m.POST("/addBaseMenu", SpringGin.Gin(controller.AddBaseMenu))
		m.POST("/getBaseMenuTree", SpringGin.Gin(controller.GetBaseMenuTree))
		m.POST("/addMenuAuthority", SpringGin.Gin(controller.AddMenuAuthority))
		m.POST("/getMenuAuthority", SpringGin.Gin(controller.GetMenuAuthority))
		m.POST("/deleteBaseMenu", SpringGin.Gin(controller.DeleteBaseMenu))
		m.POST("/updateBaseMenu", SpringGin.Gin(controller.UpdateBaseMenu))
		m.POST("/getBaseMenuById", SpringGin.Gin(controller.GetBaseMenuById))
	})

}
