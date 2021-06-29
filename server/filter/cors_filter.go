package filter

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/go-spring/spring-boot"
	"github.com/go-spring/spring-web"
)

func init() {
	SpringBoot.RegisterBean(new(CorsFilter))
}

type CorsFilter struct {
	_ SpringWeb.Filter `export:""`
}

func (filter *CorsFilter) Invoke(webCtx SpringWeb.WebContext, chain SpringWeb.FilterChain) {
	c := webCtx.NativeContext().(*gin.Context)

	method := c.Request.Method
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token,Authorization,Token,x-token")
	c.Header("Access-Control-Allow-Methods", "POST, GET, DELETE, PUT, OPTIONS")
	c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
	c.Header("Access-Control-Allow-Credentials", "true")

	// 放行所有OPTIONS方法
	if method == "OPTIONS" {
		c.AbortWithStatus(http.StatusNoContent)
	}

	// 处理请求
	c.Next()
}
