package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type responseCode int

type Response struct {
	Code responseCode `json:"code"`
	Msg  string       `json:"msg"`
	Data any          `json:"data"`
}

type TotalResponse struct {
	Response
	Total int64 `json:"total"`
}

const (
	ResponseCodeOk       responseCode = 0
	ResponseCodeFail     responseCode = -1
	ResponseCodeNotFound responseCode = 404

	ResponseCodeBadRequest             responseCode = 1001
	ResponseCodeNoAuthorized           responseCode = 1002
	ResponseCodeInvalidToken           responseCode = 1003
	ResponseCodeTokenExpired           responseCode = 1004
	ResponseCodeInvalidEmailOrPassword              = 1005
	ResponseCodeUploadFileFailed                    = 1006
)

var responseMsgMap = map[responseCode]string{
	ResponseCodeOk:                     "ok",
	ResponseCodeFail:                   "服务器错误",
	ResponseCodeBadRequest:             "请求参数错误，请检查",
	ResponseCodeNotFound:               "404 not found",
	ResponseCodeNoAuthorized:           "未登录",
	ResponseCodeInvalidToken:           "token 错误",
	ResponseCodeTokenExpired:           "token 已过期，请检查",
	ResponseCodeInvalidEmailOrPassword: "邮箱或密码错误",
	ResponseCodeUploadFileFailed:       "上传文件失败",
}

func (c responseCode) GetMsg() string {
	if msg, ok := responseMsgMap[c]; ok {
		return msg
	}

	return responseMsgMap[ResponseCodeFail]
}

func responseFail(c *gin.Context, code responseCode) {
	c.JSON(http.StatusOK, Response{Code: code, Msg: code.GetMsg(), Data: nil})
}

func responseFailWithErr(c *gin.Context, code responseCode, err error) {
	c.JSON(http.StatusOK, Response{Code: code, Msg: err.Error(), Data: nil})
}

func responseOk(c *gin.Context, data any) {
	c.JSON(http.StatusOK, Response{Code: ResponseCodeOk, Msg: ResponseCodeOk.GetMsg(), Data: data})
}

func responseOkWithTotal(c *gin.Context, data any, total int64) {
	c.JSON(http.StatusOK, TotalResponse{Response: Response{Code: ResponseCodeOk, Msg: ResponseCodeOk.GetMsg(), Data: data}, Total: total})
}
