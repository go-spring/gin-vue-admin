package router

import (
	"gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"

	"github.com/go-spring/go-spring-web/spring-gin"
	"github.com/go-spring/go-spring/spring-boot"
)

func init() {
	SpringBoot.RegisterBean(new(v1.ApiController)).Init(func(c *v1.ApiController) {

		r := SpringBoot.Route("/api",
			SpringGin.Filter(middleware.JWTAuth()),
			SpringGin.Filter(middleware.CasbinHandler()))

		r.PostMapping("/createApi", c.CreateApi)
		r.PostMapping("/deleteApi", c.DeleteApi)
		r.PostMapping("/getApiList", c.GetApiList)
		r.PostMapping("/getApiById", c.GetApiById)
		r.PostMapping("/updateApi", c.UpdateApi)
		r.PostMapping("/getAllApis", c.GetAllApis)
	})
}
