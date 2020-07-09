package router

import (
	"gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"

	"github.com/go-spring/go-spring-web/spring-gin"
	"github.com/go-spring/go-spring/spring-boot"
)

func init() {
	SpringBoot.RegisterBean(new(v1.UserController)).Init(func(c *v1.UserController) {

		r := SpringBoot.Route("/user",
			SpringGin.Filter(middleware.JWTAuth()),
			SpringGin.Filter(middleware.CasbinHandler()))

		r.POST("/changePassword", SpringGin.Gin(c.ChangePassword))
		r.POST("/uploadHeaderImg", SpringGin.Gin(c.UploadHeaderImg))
		r.POST("/getUserList", SpringGin.Gin(c.GetUserList))
		r.POST("/setUserAuthority", SpringGin.Gin(c.SetUserAuthority))
		r.DELETE("/deleteUser", SpringGin.Gin(c.DeleteUser))
	})
}
