package filter

import (
	"gin-vue-admin/global"
	"gin-vue-admin/global/response"
	"gin-vue-admin/model/request"
	"gin-vue-admin/service"

	"github.com/gin-gonic/gin"

	"github.com/go-spring/spring-boot"
	"github.com/go-spring/spring-web"
)

func init() {
	SpringBoot.RegisterBean(new(CasbinRcbaFilter))
}

type CasbinRcbaFilter struct {
	_                SpringWeb.Filter          `export:""`
	SysCasbinService *service.SysCasbinService `autowire:""`
}

func (filter *CasbinRcbaFilter) Invoke(webCtx SpringWeb.WebContext, chain SpringWeb.FilterChain) {
	c := webCtx.NativeContext().(*gin.Context)

	claims, _ := c.Get("claims")
	waitUse := claims.(*request.CustomClaims)

	// 获取请求的URI
	obj := c.Request.URL.RequestURI()
	// 获取请求方法
	act := c.Request.Method
	// 获取用户的角色
	sub := waitUse.AuthorityId

	e := filter.SysCasbinService.Casbin()
	// 判断策略中是否存在
	if global.GVA_CONFIG.System.Env == "develop" || e.Enforce(sub, obj, act) {
		c.Next()
	} else {
		response.Result(response.ERROR, gin.H{}, "权限不足", webCtx)
		c.Abort()
		return
	}
}
