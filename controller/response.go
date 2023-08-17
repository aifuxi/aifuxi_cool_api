package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ResponseCode int

const (
	Ok = 200_00

	InvalidParams ResponseCode = 400_00
	InvalidToken  ResponseCode = 400_01

	ServerError ResponseCode = 500_00
)

var codeMsg = map[ResponseCode]string{
	InvalidParams: "请求参数错误",
	InvalidToken:  "token 不合法",
	ServerError:   "服务器内部错误",
}

func (r ResponseCode) Msg() string {
	if msg, ok := codeMsg[r]; ok {
		return msg
	}

	return codeMsg[ServerError]
}

func ResponseOk(c *gin.Context, data any) {
	c.AbortWithStatusJSON(http.StatusOK, gin.H{
		"code": Ok,
		"msg":  "ok",
		"data": data,
	})
}

func ResponseErr(c *gin.Context, code ResponseCode) {
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  code.Msg(),
		"data": nil,
	})
}

func ResponseErrWithMsg(c *gin.Context, code ResponseCode, msg any) {
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  msg,
		"data": nil,
	})
}
