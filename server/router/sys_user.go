package router

import (
	"gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
	"github.com/go-spring/go-spring-web/spring-gin"
	"github.com/go-spring/go-spring/spring-boot"
)

func InitUserRouter(Router *gin.RouterGroup) {

	u := SpringBoot.Route("/user",
			SpringGin.Filter(middleware.JWTAuth()),
			SpringGin.Filter(middleware.CasbinHandler()))

	u.POST("/changePassword", SpringGin.Gin(v1.ChangePassword))
	u.POST("/uploadHeaderImg", SpringGin.Gin(v1.UploadHeaderImg))
	u.POST("/getUserList", SpringGin.Gin(v1.GetUserList))
	u.POST("/setUserAuthority", SpringGin.Gin(v1.SetUserAuthority))
	u.DELETE("/deleteUser", SpringGin.Gin(v1.DeleteUser))

	UserRouter := Router.Group("user").Use(middleware.JWTAuth()).Use(middleware.CasbinHandler())
	{
		UserRouter.POST("changePassword", v1.ChangePassword)     // 修改密码
		UserRouter.POST("uploadHeaderImg", v1.UploadHeaderImg)   // 上传头像
		UserRouter.POST("getUserList", v1.GetUserList)           // 分页获取用户列表
		UserRouter.POST("setUserAuthority", v1.SetUserAuthority) // 设置用户权限
		UserRouter.DELETE("deleteUser", v1.DeleteUser)           // 删除用户
	}
}
