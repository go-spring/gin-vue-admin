package core

import (
	_ "gin-vue-admin/controller"
	_ "gin-vue-admin/docs"
	"gin-vue-admin/filter"
	"gin-vue-admin/global"
	"gin-vue-admin/initialize"
	"github.com/gin-gonic/gin"
	"github.com/go-spring/go-spring-web/spring-gin"
	"github.com/go-spring/go-spring-web/spring-web"
	"github.com/go-spring/go-spring/spring-boot"
	_ "github.com/go-spring/go-spring/starter-gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func RunWindowsServer() {
	if global.GVA_CONFIG.System.UseMultipoint {
		// 初始化redis服务
		initialize.Redis()
	}

	// 添加 swagger 接口
	SpringBoot.GET("/swagger/*any", SpringGin.Gin(ginSwagger.WrapHandler(swaggerFiles.Handler)))

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
