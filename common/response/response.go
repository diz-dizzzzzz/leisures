package response

import (
	"net/http"

	"acupofcoffee/common/errorx"

	"github.com/zeromicro/go-zero/rest/httpx"
)

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}

// Success 成功响应
func Success(w http.ResponseWriter, data interface{}) {
	httpx.OkJson(w, Response{
		Code: errorx.CodeSuccess,
		Msg:  "success",
		Data: data,
	})
}

// Error 错误响应
func Error(w http.ResponseWriter, err error) {
	var code int
	var msg string

	switch e := err.(type) {
	case *errorx.CodeError:
		code = e.GetCode()
		msg = e.GetMsg()
	default:
		code = errorx.CodeServerError
		msg = err.Error()
	}

	httpx.OkJson(w, Response{
		Code: code,
		Msg:  msg,
	})
}

// ParamError 参数错误响应
func ParamError(w http.ResponseWriter, err error) {
	httpx.OkJson(w, Response{
		Code: errorx.CodeParamError,
		Msg:  "参数错误: " + err.Error(),
	})
}

// Unauthorized 未授权响应
func Unauthorized(w http.ResponseWriter, msg string) {
	if msg == "" {
		msg = "未授权"
	}
	httpx.OkJson(w, Response{
		Code: errorx.CodeUnauthorized,
		Msg:  msg,
	})
}

// Forbidden 禁止访问响应
func Forbidden(w http.ResponseWriter, msg string) {
	if msg == "" {
		msg = "禁止访问"
	}
	httpx.OkJson(w, Response{
		Code: errorx.CodeForbidden,
		Msg:  msg,
	})
}

// NotFound 资源不存在响应
func NotFound(w http.ResponseWriter, msg string) {
	if msg == "" {
		msg = "资源不存在"
	}
	httpx.OkJson(w, Response{
		Code: errorx.CodeNotFound,
		Msg:  msg,
	})
}

