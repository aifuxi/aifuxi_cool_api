package middleware

import (
	"errors"
	"strings"

	"api.aifuxi.cool/internal"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"go.uber.org/zap"
)

func JwtAuth() func(c *gin.Context) {
	return func(c *gin.Context) {
		bearerToken := c.Request.Header.Get("Authorization")

		if bearerToken == "" {
			// TODO： 错误处理
			c.Abort()
			return
		}

		// Bearer xxxxxx
		parts := strings.SplitN(bearerToken, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			zap.L().Debug("走到这，token格式不对")
			// TODO： 错误处理
			c.Abort()
			return
		}

		mc, err := internal.ParseToken(parts[1])
		if err != nil {

			if errors.Is(err, jwt.ErrTokenExpired) {
				// TODO： 错误处理
				c.Abort()
				return
			}

			// TODO： 错误处理
			c.Abort()
			return
		}

		// 把 token 中存放的 email 挂到 context 上
		// 后面的 handler 就可以通过 c.Get("email") 拿到挂载的信息
		c.Set("email", mc.Email)
		c.Next()
	}
}
