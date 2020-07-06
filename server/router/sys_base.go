package router

import (
	"gin-vue-admin/api/v1"

	"github.com/gin-gonic/gin"
	"github.com/go-spring/go-spring-web/spring-gin"
	"github.com/go-spring/go-spring/spring-boot"
)

func InitBaseRouter(Router *gin.RouterGroup) (R gin.IRoutes) {

	// 5. 迁移接口到 go-spring，老接口先保留
	r := SpringBoot.Route("/base")
	r.HandlePost("/login", SpringGin.Gin(v1.Login))
	r.HandlePost("/captcha", SpringGin.Gin(v1.Captcha))
	r.HandlePost("/register", SpringGin.Gin(v1.Register))

	BaseRouter := Router.Group("base")
	{
		BaseRouter.POST("register", v1.Register)
		BaseRouter.POST("login", v1.Login)
		BaseRouter.POST("captcha", v1.Captcha)
		BaseRouter.GET("captcha/:captchaId", v1.CaptchaImg)
	}
	return BaseRouter
}
