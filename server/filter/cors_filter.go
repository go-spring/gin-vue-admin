package filter

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/go-spring/go-spring-web/spring-web"
	"github.com/go-spring/go-spring/spring-boot"
)

func init() {
	SpringBoot.RegisterBean(new(CorsFilter))
}

type CorsFilter struct {
	_ SpringWeb.Filter `export:""`
}

func (filter *CorsFilter) Invoke(webCtx SpringWeb.WebContext, chain SpringWeb.FilterChain) {
	ctx := webCtx.NativeContext().(*gin.Context)
	method := webCtx.Request().Method
	webCtx.Header("Access-Control-Allow-Origin", "*")
	webCtx.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token,Authorization,Token,x-token")
	webCtx.Header("Access-Control-Allow-Methods", "POST, GET, DELETE, PUT, OPTIONS")
	webCtx.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
	webCtx.Header("Access-Control-Allow-Credentials", "true")

	if method == "OPTIONS" {
		webCtx.NoContent(http.StatusNoContent)
		ctx.Abort()
		return
	}

	chain.Next(webCtx)
}
