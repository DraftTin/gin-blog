package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	Success = iota
	Error
)

type Response struct {
	Code    int    `json:"code"`
	Data    any    `json:"data"`
	Message string `json:"message"`
}

func Result(code int, data any, message string, ctx *gin.Context) {
	ctx.JSON(http.StatusOK, Response{
		Code:    code,
		Data:    data,
		Message: message,
	})
}

func OK(data any, msg string, ctx *gin.Context) {
	Result(Success, data, msg, ctx)
}

func OKWithData(data any, ctx *gin.Context) {
	Result(Success, data, "Success", ctx)
}

func OKWithMsg(msg string, ctx *gin.Context) {
	Result(Success, struct{}{}, msg, ctx)
}

func Fail(data any, msg string, ctx *gin.Context) {
	Result(Error, data, msg, ctx)
}

func FailWithMsg(msg string, ctx *gin.Context) {
	Result(Error, struct{}{}, msg, ctx)
}

func FailWithCode(code ErrorCode, ctx *gin.Context) {
	msg, ok := ErrorMsg[code]
	if ok == false {
		FailWithMsg("unknow error", ctx)
		return
	}
	FailWithMsg(msg, ctx)
}
