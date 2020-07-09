package router

import (
	"gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"
	"github.com/go-spring/go-spring-web/spring-gin"
	"github.com/go-spring/go-spring/spring-boot"
)

func init() {

	SpringBoot.RegisterBean(new(v1.UserController)).Init(func(controller *v1.UserController) {
		u := SpringBoot.Route("/user",
			SpringGin.Filter(middleware.JWTAuth()),
			SpringGin.Filter(middleware.CasbinHandler()))

		u.POST("/changePassword", SpringGin.Gin(controller.ChangePassword))
		u.POST("/uploadHeaderImg", SpringGin.Gin(controller.UploadHeaderImg))
		u.POST("/getUserList", SpringGin.Gin(controller.GetUserList))
		u.POST("/setUserAuthority", SpringGin.Gin(controller.SetUserAuthority))
		u.DELETE("/deleteUser", SpringGin.Gin(controller.DeleteUser))
	})

}
