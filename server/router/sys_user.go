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

		r.PostMapping("/changePassword", c.ChangePassword)
		r.PostMapping("/uploadHeaderImg", c.UploadHeaderImg)
		r.PostMapping("/getUserList", c.GetUserList)
		r.PostMapping("/setUserAuthority", c.SetUserAuthority)
		r.DELETE("/deleteUser", c.DeleteUser)
	})
}
