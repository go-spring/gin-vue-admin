package initialize

import (
	_ "gin-vue-admin/docs"
	"gin-vue-admin/global"
	"gin-vue-admin/middleware"
	"gin-vue-admin/router"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-spring/go-spring-web/spring-gin"
	"github.com/go-spring/go-spring-web/spring-web"
	"github.com/go-spring/go-spring/spring-boot"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

const Host = "http://127.0.0.1:8887"

// 初始化总路由
func Routers() *gin.Engine {
	var Router = gin.Default()
	// Router.Use(middleware.LoadTls())  // 打开就能玩https了
	global.GVA_LOG.Debug("use middleware logger")

	redirect := map[string]string{ // 已经迁移的地址
		"/base/login":   				"",
		"/base/captcha": 				"",
		"/base/register":				"",
		"/base/captcha/:captchaId": 	"",

		"/user/changePassword":   		"",
		"/user/uploadHeaderImg":  		"",
		"/user/getUserList":	  		"",
		"/user/setUserAuthority": 		"",
		"/user/deleteUser":		  		"",

		"/menu/getMenu":          		"",
		"/menu/getMenuList":      		"",
		"/menu/addBaseMenu":      		"",
		"/menu/getBaseMenuTree":  		"",
		"/menu/addMenuAuthority": 		"",
		"/menu/getMenuAuthority": 		"",
		"/menu/deleteBaseMenu":   		"",
		"/menu/updateBaseMenu":   		"",
		"/menu/getBaseMenuById":  		"",

		"/authority/createAuthority":   "",
		"/authority/deleteAuthority":   "",
		"/authority/updateAuthority":   "",
		"/authority/copyAuthority":     "",
		"/authority/getAuthorityList":  "",
		"/authority/setDataAuthority":  "",

		"/customer/customer":     		"",	// POST, PUT, DELETE, GET 路由同名，因此只在map中注册一个
		"/customer/customerList": 		"",

		"/api/createApi":  				"",
		"/api/deleteApi":  				"",
		"/api/getApiList": 				"",
		"/api/getApiById": 				"",
		"/api/updateApi":  				"",
		"/api/getAllApis": 				"",

		"/fileUploadAndDownload/upload": 					"",
		"/fileUploadAndDownload/getFileList": 				"",
		"/fileUploadAndDownload/deleteFile": 				"",
		"/fileUploadAndDownload/breakpointContinue": 		"",
		"/fileUploadAndDownload/findFile": 					"",
		"/fileUploadAndDownload/breakpointContinueFinish":  "",
		"/fileUploadAndDownload/removeChunk": 				"",

		"/autoCode/createTemp":					"",

		"/casbin/updateCasbin":					"",
		"/casbin/getPolicyPathByAuthorityId":	"",
		"/casbin/casbinTest/:pathParam":		"",

		"/jwt/jsonInBlacklist":					"",

		"/system/getSystemConfig":				"",
		"/system/setSystemConfig":				"",

		"/workflow/createWorkFlow":				"",
	}

	// 3. 增加一个重定向中间件，对于可迁移的接口使用重定向机制进行迁移
	Router.Use(func(ctx *gin.Context) {
		if _, ok := redirect[ctx.FullPath()]; ok {
			ctx.Redirect(http.StatusTemporaryRedirect, Host + ctx.Request.RequestURI)
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
