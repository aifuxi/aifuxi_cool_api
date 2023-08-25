package controller

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type ResponseCode int

const (
	Ok = 200_00

	InvalidParams    ResponseCode = 400_00
	NoAuthorized     ResponseCode = 400_01
	TokenExpired     ResponseCode = 400_08
	InvalidToken     ResponseCode = 400_09
	ParseTokenFailed ResponseCode = 400_10

	ServerError ResponseCode = 500_00
)

var codeMsg = map[ResponseCode]string{
	InvalidParams:    "request parameter error",
	NoAuthorized:     "no authorized",
	TokenExpired:     "token expired",
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
	c.JSON(http.StatusOK, gin.H{
		"code": Ok,
		"msg":  "ok",
		"data": data,
	})
}

func ResponseOkWithTotal(c *gin.Context, data any, total int64) {
	c.JSON(http.StatusOK, gin.H{
		"code":  Ok,
		"msg":   "ok",
		"data":  data,
		"total": total,
	})
}

func ResponseErr(c *gin.Context, code ResponseCode, arg any) {
	msg := code.Msg()

	if msg2, ok := arg.(string); ok {
		zap.L().Error(c.Request.RequestURI, zap.Error(errors.New(msg2)))
		msg = msg2
	}

	if err, ok := arg.(error); ok {
		zap.L().Error(c.Request.RequestURI, zap.Error(err))
		msg = err.Error()
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  msg,
		"data": nil,
	})
}
