package filter

import (
	"gin-vue-admin/global/response"
	"gin-vue-admin/middleware"
	"gin-vue-admin/model"
	"gin-vue-admin/service"

	"github.com/gin-gonic/gin"

	"github.com/go-spring/go-spring-web/spring-web"
	"github.com/go-spring/go-spring/spring-boot"
)

func init() {
	SpringBoot.RegisterBean(new(JwtFilter))
}

type JwtFilter struct {
	_                   SpringWeb.Filter             `export:""`
	JwtBlackListService *service.JwtBlackListService `autowire:""`
}

func (filter *JwtFilter) Invoke(webCtx SpringWeb.WebContext, chain SpringWeb.FilterChain) {
	ctx := webCtx.NativeContext().(*gin.Context)
	token := webCtx.GetHeader("x-token")
	if token == "" || len(token) == 0 {
		response.Result(response.ERROR, map[string]interface{}{
			"reload": true,
		}, "未登录或非法访问", webCtx)
		ctx.Abort()
		return
	}

	modelToken := model.JwtBlacklist{
		Jwt: token,
	}

	if filter.JwtBlackListService.IsBlacklist(token, modelToken) {
		response.Result(response.ERROR, map[string]interface{}{
			"reload": true,
		}, "未登录或非法访问", webCtx)
		ctx.Abort()
		return
	}

	j := middleware.NewJWT()
	// parseToken 解析token包含的信息
	claims, err := j.ParseToken(token)
	if err != nil {
		if err == middleware.TokenExpired {
			response.Result(response.ERROR, map[string]interface{}{
				"reload": true,
			}, "授权已过期", webCtx)
			ctx.Abort()
			return
		}
		response.Result(response.ERROR, map[string]interface{}{
			"reload": true,
		}, err.Error(), webCtx)
		ctx.Abort()
		return
	}
	webCtx.Set("claims", claims)
	chain.Next(webCtx)
}
