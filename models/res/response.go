package res

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data any    `json:"data"`
}

const (
	Success = 0
	Err     = 7
)

func Result(code int, msg string, data any, c *gin.Context) {
	c.JSON(http.StatusOK, Response{
		Code: code,
		Msg:  msg,
		Data: data,
	})
}

func OK(data any, msg string, c *gin.Context) {
	Result(Success, msg, data, c)
}

func OKWithMsg(msg string, c *gin.Context) {
	Result(Success, msg, map[string]any{}, c)
}

func OKWithData(data any, c *gin.Context) {
	Result(Success, "成功", data, c)
}
func Fail(data any, msg string, c *gin.Context) {
	Result(Err, msg, data, c)
}

func FailWithMsg(msg string, c *gin.Context) {
	Result(Err, msg, map[string]any{}, c)
}

func FailWithCode(code ErrorCode, c *gin.Context) {
	msg, ok := ErrorMap[code]
	if ok {
		Result(int(code), msg, map[string]any{}, c)
		return
	}
	Result(Err, "未知错误", map[string]any{}, c)
}
