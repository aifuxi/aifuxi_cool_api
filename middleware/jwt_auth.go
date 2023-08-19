package middleware

import (
	"strings"

	"api.aifuxi.cool/controller"
	"api.aifuxi.cool/internal"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func JwtAuth() func(c *gin.Context) {
	return func(c *gin.Context) {
		bearerToken := c.Request.Header.Get("Authorization")

		if bearerToken == "" {
			controller.ResponseErr(c, controller.NoAuthorized)
			c.Abort()
			return
		}

		// Bearer xxxxxx
		parts := strings.SplitN(bearerToken, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			zap.L().Debug("走到这，token格式不对")
			controller.ResponseErr(c, controller.InvalidToken)
			c.Abort()
			return
		}

		mc, err := internal.ParseToken(parts[1])
		if err != nil {
			zap.L().Debug("走到这，解析出错了")
			controller.ResponseErr(c, controller.InvalidToken)
			c.Abort()
			return
		}

		// 把 token 中存放的 email 挂到 context 上
		// 后面的 handler 就可以通过 c.Get("email") 拿到挂载的信息
		c.Set("email", mc.Email)
		c.Next()
	}
}
