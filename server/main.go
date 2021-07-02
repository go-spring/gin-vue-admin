package main

import (
	_ "gin-vue-admin/controller"
	_ "gin-vue-admin/core"
	_ "gin-vue-admin/docs"
	"gin-vue-admin/filter"
	_ "gin-vue-admin/utils"

	"github.com/gin-gonic/gin"
	"github.com/go-spring/spring-boot"
	"github.com/go-spring/spring-gin"
	"github.com/go-spring/spring-web"
	_ "github.com/go-spring/starter-gin"
	_ "github.com/go-spring/starter-gorm/mysql"
	"github.com/jinzhu/gorm"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
)

// @title Swagger Example API
// @version 0.0.1
// @description This is a sample Server pets
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name x-token
// @BasePath /
func main() {

	// 添加 swagger 接口
	SpringBoot.HandleGet("/swagger/*any", SpringGin.Gin(ginSwagger.WrapHandler(swaggerFiles.Handler)))

	// 1. 引入 go-spring web 组件，关闭 swagger 及默认 filter
	SpringBoot.Config(func(c SpringWeb.WebContainer, port int) {
		c.SetRecoveryFilter(SpringGin.Filter(gin.Recovery()))
		c.SetLoggerFilter(SpringGin.Filter(gin.Logger()))
		c.AddFilter(SpringBoot.FilterBean((*filter.CorsFilter)(nil)))
		c.SetEnableSwagger(false)
	}, "1:${web.server.port}")

	SpringBoot.Config(func(db *gorm.DB, maxIdleConns, maxOpenConns int, logMode bool) {
		db.DB().SetMaxIdleConns(maxIdleConns)
		db.DB().SetMaxOpenConns(maxOpenConns)
		db.LogMode(logMode)
		//DBTables(db)
	}, "", "${db.max-idle-conns}", "${db.max-open-conns}", "${db.log-mode}")

	// 2. 启动 SpringWeb 服务器
	SpringBoot.RunApplication()
}
