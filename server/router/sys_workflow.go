package router

import (
	"gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
	"github.com/go-spring/go-spring-web/spring-gin"
	"github.com/go-spring/go-spring/spring-boot"
)

func InitWorkflowRouter(Router *gin.RouterGroup) {

	w := SpringBoot.Route("/workflow",
		SpringGin.Filter(middleware.JWTAuth()),
		SpringGin.Filter(middleware.CasbinHandler()))

	w.POST("/createWorkFlow", SpringGin.Gin(v1.CreateWorkFlow))

	WorkflowRouter := Router.Group("workflow").Use(middleware.JWTAuth()).Use(middleware.CasbinHandler())
	{
		WorkflowRouter.POST("createWorkFlow", v1.CreateWorkFlow) // 创建工作流
	}
}
