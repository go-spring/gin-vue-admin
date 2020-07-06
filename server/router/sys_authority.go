package router

import (
	"gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
	"github.com/go-spring/go-spring-web/spring-gin"
	"github.com/go-spring/go-spring/spring-boot"
)

func InitAuthorityRouter(Router *gin.RouterGroup) {

	a := SpringBoot.Route("/authority",
		SpringGin.Filter(middleware.JWTAuth()),
		SpringGin.Filter(middleware.CasbinHandler()))

	a.POST("/createAuthority", SpringGin.Gin(v1.CreateAuthority))
	a.POST("/deleteAuthority", SpringGin.Gin(v1.DeleteAuthority))
	a.PUT("/updateAuthority", SpringGin.Gin(v1.UpdateAuthority))
	a.POST("/copyAuthority", SpringGin.Gin(v1.CopyAuthority))
	a.POST("/getAuthorityList", SpringGin.Gin(v1.GetAuthorityList))
	a.POST("/setDataAuthority", SpringGin.Gin(v1.SetDataAuthority))

	AuthorityRouter := Router.Group("authority").Use(middleware.JWTAuth()).Use(middleware.CasbinHandler())
	{
		AuthorityRouter.POST("createAuthority", v1.CreateAuthority)   // 创建角色
		AuthorityRouter.POST("deleteAuthority", v1.DeleteAuthority)   // 删除角色
		AuthorityRouter.PUT("updateAuthority", v1.UpdateAuthority)    // 更新角色
		AuthorityRouter.POST("copyAuthority", v1.CopyAuthority)       // 更新角色
		AuthorityRouter.POST("getAuthorityList", v1.GetAuthorityList) // 获取角色列表
		AuthorityRouter.POST("setDataAuthority", v1.SetDataAuthority) // 设置角色资源权限
	}
}
