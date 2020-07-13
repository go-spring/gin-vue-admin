package response

import (
	"net/http"

	"github.com/go-spring/go-spring-web/spring-web"
)

type Response struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

const (
	ERROR   = 7
	SUCCESS = 0
)

func Result(code int, data interface{}, msg string, webCtx SpringWeb.WebContext) {
	// 开始时间
	webCtx.JSON(http.StatusOK, Response{
		code,
		data,
		msg,
	})
}

func Ok(webCtx SpringWeb.WebContext) {
	Result(SUCCESS, map[string]interface{}{}, "操作成功", webCtx)
}

func OkWithMessage(message string, webCtx SpringWeb.WebContext) {
	Result(SUCCESS, map[string]interface{}{}, message, webCtx)
}

func OkWithData(data interface{}, webCtx SpringWeb.WebContext) {
	Result(SUCCESS, data, "操作成功", webCtx)
}

func OkDetailed(data interface{}, message string, webCtx SpringWeb.WebContext) {
	Result(SUCCESS, data, message, webCtx)
}

func Fail(webCtx SpringWeb.WebContext) {
	Result(ERROR, map[string]interface{}{}, "操作失败", webCtx)
}

func FailWithMessage(message string, webCtx SpringWeb.WebContext) {
	Result(ERROR, map[string]interface{}{}, message, webCtx)
}

func FailWithDetailed(code int, data interface{}, message string, webCtx SpringWeb.WebContext) {
	Result(code, data, message, webCtx)
}
