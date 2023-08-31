package api

import (
	"api.aifuxi.cool/util"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"strings"
)

func JwtAuth() func(c *gin.Context) {
	return func(c *gin.Context) {
		bearerToken := c.Request.Header.Get("Authorization")

		if bearerToken == "" {
			responseFail(c, ResponseCodeNoAuthorized)
			c.Abort()
			return
		}

		// Bearer xxxxxx
		parts := strings.SplitN(bearerToken, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			responseFail(c, ResponseCodeInvalidToken)
			c.Abort()
			return
		}

		mc, err := util.ParseToken(parts[1])
		if err != nil {

			if errors.Is(err, jwt.ErrTokenExpired) {
				responseFail(c, ResponseCodeTokenExpired)

				c.Abort()
				return
			}

			responseFail(c, ResponseCodeInvalidToken)
			c.Abort()
			return
		}

		// 把 token 中存放的 email 挂到 context 上
		// 后面的 handler 就可以通过 c.Get("email") 拿到挂载的信息
		c.Set("email", mc.Email)
		c.Next()
	}
}
