package router

import (
	"gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"

	"github.com/go-spring/go-spring-web/spring-gin"
	"github.com/go-spring/go-spring/spring-boot"
)

func init() {

	u := SpringBoot.Route("/user",
		SpringGin.Filter(middleware.JWTAuth()),
		SpringGin.Filter(middleware.CasbinHandler()))

	u.POST("/changePassword", SpringGin.Gin(v1.ChangePassword))
	u.POST("/uploadHeaderImg", SpringGin.Gin(v1.UploadHeaderImg))
	u.POST("/getUserList", SpringGin.Gin(v1.GetUserList))
	u.POST("/setUserAuthority", SpringGin.Gin(v1.SetUserAuthority))
	u.DELETE("/deleteUser", SpringGin.Gin(v1.DeleteUser))
}
