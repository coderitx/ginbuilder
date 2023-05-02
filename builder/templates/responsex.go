package templates

var ResponsexTemplate = `package responsex

import (
	"{{.pkgname}}/common/errorx"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Code    int    #json:"code"#
	Data    any    #json:"data"#
	Message string #json:"msg"#
}

func Result(code int, data any, msg string, c *gin.Context) {
	c.JSON(http.StatusOK, Response{
		Code:    code,
		Data:    data,
		Message: msg,
	})
}

func Ok(data any, msg string, c *gin.Context) {
	Result(int(errorx.SuccessCode), data, msg, c)
}

func OkWithData(data any, c *gin.Context) {
	Result(int(errorx.SuccessCode), data, "成功", c)
}

func OkWithMessage(msg string, c *gin.Context) {
	Result(int(errorx.SuccessCode), map[string]any{}, msg, c)
}

func OkWith(c *gin.Context) {
	Result(int(errorx.SuccessCode), map[string]any{}, "成功", c)
}

func Fail(data any, msg string, c *gin.Context) {
	Result(int(errorx.FailedCode), data, msg, c)
}

func FailWithMessage(msg string, c *gin.Context) {
	Result(int(errorx.FailedCode), map[string]any{}, msg, c)
}

func FailWithCode(code errorx.ErrorCode, msg string, c *gin.Context) {
	msg, ok := errorx.ErrorMap[code]
	if ok {
		Result(int(code), map[string]any{}, msg, c)
	}
	Result(int(errorx.FailedCode), map[string]any{}, msg, c)
}
`
