package filter

import (
	"gin-vue-admin/global"
	"gin-vue-admin/global/response"
	"gin-vue-admin/model/request"
	"gin-vue-admin/service"

	"github.com/gin-gonic/gin"

	"github.com/go-spring/go-spring-web/spring-web"
	"github.com/go-spring/go-spring/spring-boot"
)

func init() {
	SpringBoot.RegisterBean(new(CasbinRcbaFilter))
}

type CasbinRcbaFilter struct {
	_                SpringWeb.Filter          `export:""`
	SysCasbinService *service.SysCasbinService `autowire:""`
}

func (filter *CasbinRcbaFilter) Invoke(webCtx SpringWeb.WebContext, chain SpringWeb.FilterChain) {
	ctx := webCtx.NativeContext().(*gin.Context)
	claims := webCtx.Get("claims")
	waitUse := claims.(*request.CustomClaims)

	// 获取请求的URI
	obj := webCtx.Request().RequestURI
	// 获取请求方法
	act := webCtx.Request().Method
	// 获取用户的角色
	sub := waitUse.AuthorityId
	e := filter.SysCasbinService.Casbin()

	if global.GVA_CONFIG.System.Env == "develop" || e.Enforce(sub, obj, act) {
		chain.Next(webCtx)
	} else {
		response.Result(response.ERROR, map[string]interface{}{}, "权限不足", webCtx)
		ctx.Abort()
		return
	}
}
