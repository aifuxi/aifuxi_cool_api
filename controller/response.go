package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ResponseCode int

const (
	Ok = 200_00

	InvalidParams    ResponseCode = 400_00
	NoAuthorized     ResponseCode = 400_01
	InvalidToken     ResponseCode = 400_09
	ParseTokenFailed ResponseCode = 400_10

	ServerError ResponseCode = 500_00
)

var codeMsg = map[ResponseCode]string{
	InvalidParams:    "request parameter error",
	NoAuthorized:     "no authorized",
	InvalidToken:     "invalid token",
	ParseTokenFailed: "parse token failed",
	ServerError:      "internal server error",
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
