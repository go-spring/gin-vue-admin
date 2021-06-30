package core

import (
	_ "gin-vue-admin/controller"
	_ "gin-vue-admin/docs"
	"gin-vue-admin/filter"
	_ "gin-vue-admin/utils"

	"github.com/gin-gonic/gin"
	"github.com/go-spring/spring-boot"
	"github.com/go-spring/spring-gin"
	"github.com/go-spring/spring-web"
	_ "github.com/go-spring/starter-gin"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
)

func RunWindowsServer() {

	// TODO 如何对引入的包施加外部限制，考虑开放内部 API
	//if global.GVA_CONFIG.System.UseMultipoint {
	//	// 初始化redis服务
	//	initialize.Redis()
	//}

	// 添加 swagger 接口
	SpringBoot.HandleGet("/swagger/*any", SpringGin.Gin(ginSwagger.WrapHandler(swaggerFiles.Handler)))

	// 1. 引入 go-spring web 组件，关闭 swagger 及默认 filter
	SpringBoot.Config(func(c SpringWeb.WebContainer, port int) {
		c.SetRecoveryFilter(SpringGin.Filter(gin.Recovery()))
		c.SetLoggerFilter(SpringGin.Filter(gin.Logger()))
		c.AddFilter(SpringBoot.FilterBean((*filter.CorsFilter)(nil)))
		c.SetEnableSwagger(false)
	}, "1:${web.server.port}")

	// 2. 启动 SpringWeb 服务器
	SpringBoot.RunApplication("./")
}
