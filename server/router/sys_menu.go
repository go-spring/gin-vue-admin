package router

import (
	"gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
	SpringGin "github.com/go-spring/go-spring-web/spring-gin"
	SpringBoot "github.com/go-spring/go-spring/spring-boot"
)

func InitMenuRouter(Router *gin.RouterGroup) (R gin.IRoutes) {

	m := SpringBoot.Route("/menu", SpringGin.Filter(middleware.JWTAuth()), SpringGin.Filter(middleware.CasbinHandler()))
	m.POST("/getMenu", SpringGin.Gin(v1.GetMenu))
	m.POST("/getMenuList", SpringGin.Gin(v1.GetMenuList))
	m.POST("/addBaseMenu", SpringGin.Gin(v1.AddBaseMenu))
	m.POST("/getBaseMenuTree", SpringGin.Gin(v1.GetBaseMenuTree))
	m.POST("/addMenuAuthority", SpringGin.Gin(v1.AddMenuAuthority))
	m.POST("/getMenuAuthority", SpringGin.Gin(v1.GetMenuAuthority))
	m.POST("/deleteBaseMenu", SpringGin.Gin(v1.DeleteBaseMenu))
	m.POST("/updateBaseMenu", SpringGin.Gin(v1.UpdateBaseMenu))
	m.POST("/getBaseMenuById", SpringGin.Gin(v1.GetBaseMenuById))

	MenuRouter := Router.Group("menu").Use(middleware.JWTAuth()).Use(middleware.CasbinHandler())
	{
		MenuRouter.POST("getMenu", v1.GetMenu)                   // 获取菜单树
		MenuRouter.POST("getMenuList", v1.GetMenuList)           // 分页获取基础menu列表
		MenuRouter.POST("addBaseMenu", v1.AddBaseMenu)           // 新增菜单
		MenuRouter.POST("getBaseMenuTree", v1.GetBaseMenuTree)   // 获取用户动态路由
		MenuRouter.POST("addMenuAuthority", v1.AddMenuAuthority) //	增加menu和角色关联关系
		MenuRouter.POST("getMenuAuthority", v1.GetMenuAuthority) // 获取指定角色menu
		MenuRouter.POST("deleteBaseMenu", v1.DeleteBaseMenu)     // 删除菜单
		MenuRouter.POST("updateBaseMenu", v1.UpdateBaseMenu)     // 更新菜单
		MenuRouter.POST("getBaseMenuById", v1.GetBaseMenuById)   // 根据id获取菜单
	}
	return MenuRouter
}
