package errorx

import "fmt"

// 错误码定义
const (
	CodeSuccess      = 0
	CodeParamError   = 400
	CodeUnauthorized = 401
	CodeForbidden    = 403
	CodeNotFound     = 404
	CodeServerError  = 500
)

// 错误消息
var codeMsg = map[int]string{
	CodeSuccess:      "success",
	CodeParamError:   "参数错误",
	CodeUnauthorized: "未授权",
	CodeForbidden:    "禁止访问",
	CodeNotFound:     "资源不存在",
	CodeServerError:  "服务器内部错误",
}

type CodeError struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

func NewCodeError(code int, msg string) *CodeError {
	return &CodeError{
		Code: code,
		Msg:  msg,
	}
}

func NewDefaultError(msg string) *CodeError {
	return &CodeError{
		Code: CodeServerError,
		Msg:  msg,
	}
}

func NewParamError(msg string) *CodeError {
	return &CodeError{
		Code: CodeParamError,
		Msg:  msg,
	}
}

func NewUnauthorizedError(msg string) *CodeError {
	return &CodeError{
		Code: CodeUnauthorized,
		Msg:  msg,
	}
}

func NewNotFoundError(msg string) *CodeError {
	return &CodeError{
		Code: CodeNotFound,
		Msg:  msg,
	}
}

func (e *CodeError) Error() string {
	return fmt.Sprintf("code: %d, msg: %s", e.Code, e.Msg)
}

func (e *CodeError) GetCode() int {
	return e.Code
}

func (e *CodeError) GetMsg() string {
	return e.Msg
}

// GetCodeMsg 根据错误码获取默认错误消息
func GetCodeMsg(code int) string {
	if msg, ok := codeMsg[code]; ok {
		return msg
	}
	return "未知错误"
}

