package initialize

import (
	"net/http"

	_ "gin-vue-admin/docs"
	"gin-vue-admin/global"
	"gin-vue-admin/middleware"
	"gin-vue-admin/router"

	"github.com/gin-gonic/gin"
	"github.com/go-spring/go-spring-web/spring-gin"
	"github.com/go-spring/go-spring-web/spring-web"
	"github.com/go-spring/go-spring/spring-boot"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

// 初始化总路由
func Routers() *gin.Engine {
	var Router = gin.Default()
	// Router.Use(middleware.LoadTls())  // 打开就能玩https了
	global.GVA_LOG.Debug("use middleware logger")

	redirect := map[string]string{ // 已经迁移的地址
		//"/base/login": "http://127.0.0.1:8887/base/login",
		"/base/captcha": "http://127.0.0.1:8887/base/captcha",
	}

	// 3. 增加一个重定向中间件，对于可迁移的接口使用重定向机制进行迁移
	Router.Use(func(ctx *gin.Context) {
		if redirectPath, ok := redirect[ctx.FullPath()]; ok {
			ctx.Redirect(http.StatusTemporaryRedirect, redirectPath)
			ctx.Abort()
		}
	})

	// 4. 给 go-spring 也配置跨域组件
	SpringBoot.Config(func(c SpringWeb.WebContainer) {
		c.AddFilter(SpringGin.Filter(middleware.Cors()))
	})

	// 跨域
	Router.Use(middleware.Cors())
	global.GVA_LOG.Debug("use middleware cors")

	Router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	global.GVA_LOG.Debug("register swagger handler")

	// 方便统一添加路由组前缀 多服务器上线使用
	ApiGroup := Router.Group("")
	router.InitUserRouter(ApiGroup)                  // 注册用户路由
	router.InitBaseRouter(ApiGroup)                  // 注册基础功能路由 不做鉴权
	router.InitMenuRouter(ApiGroup)                  // 注册menu路由
	router.InitAuthorityRouter(ApiGroup)             // 注册角色路由
	router.InitApiRouter(ApiGroup)                   // 注册功能api路由
	router.InitFileUploadAndDownloadRouter(ApiGroup) // 文件上传下载功能路由
	router.InitWorkflowRouter(ApiGroup)              // 工作流相关路由
	router.InitCasbinRouter(ApiGroup)                // 权限相关路由
	router.InitJwtRouter(ApiGroup)                   // jwt相关路由
	router.InitSystemRouter(ApiGroup)                // system相关路由
	router.InitCustomerRouter(ApiGroup)              // 客户路由
	router.InitAutoCodeRouter(ApiGroup)              // 创建自动化代码
	global.GVA_LOG.Info("router register success")
	return Router
}
