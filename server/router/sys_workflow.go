package router

import (
	"gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"

	"github.com/go-spring/go-spring-web/spring-gin"
	"github.com/go-spring/go-spring/spring-boot"
)

func init() {

	SpringBoot.RegisterBean(new(v1.WorkFlowController)).Init(func(controller *v1.WorkFlowController) {
		w := SpringBoot.Route("/workflow",
			SpringGin.Filter(middleware.JWTAuth()),
			SpringGin.Filter(middleware.CasbinHandler()))

		w.POST("/createWorkFlow", SpringGin.Gin(controller.CreateWorkFlow))
	})

}
