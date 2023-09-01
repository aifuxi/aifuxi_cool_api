package api

import (
	"api.aifuxi.cool/util"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"go.uber.org/zap"
	"net"
	"net/http"
	"net/http/httputil"
	"os"
	"runtime/debug"
	"strings"
	"time"
)

const authKey = "Authorization"
const authPrefix = "Bearer"
const customCtxKey = "email"

func JwtAuth() func(c *gin.Context) {
	return func(c *gin.Context) {
		bearerToken := c.Request.Header.Get(authKey)

		if bearerToken == "" {
			responseFail(c, ResponseCodeNoAuthorized)
			c.Abort()
			return
		}

		// Bearer xxxxxx
		parts := strings.SplitN(bearerToken, " ", 2)
		if !(len(parts) == 2 && parts[0] == authPrefix) {
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
		c.Set(customCtxKey, mc.Email)
		c.Next()
	}
}

func GinLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		query := c.Request.URL.RawQuery
		c.Next()

		cost := time.Since(start)
		zap.L().Info(path,
			zap.Int("status", c.Writer.Status()),
			zap.String("method", c.Request.Method),
			zap.String("path", path),
			zap.String("query", query),
			zap.String("ip", c.ClientIP()),
			zap.String("errors", c.Errors.ByType(gin.ErrorTypePrivate).String()),
			zap.Duration("cost", cost),
		)
	}
}

// GinRecovery recover掉项目可能出现的panic，并使用zap记录相关日志
func GinRecovery(stack bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				// Check for a broken connection, as it is not really a
				// condition that warrants a panic stack trace.
				var brokenPipe bool
				if ne, ok := err.(*net.OpError); ok {
					if se, ok := ne.Err.(*os.SyscallError); ok {
						if strings.Contains(strings.ToLower(se.Error()), "broken pipe") || strings.Contains(strings.ToLower(se.Error()), "connection reset by peer") {
							brokenPipe = true
						}
					}
				}

				httpRequest, _ := httputil.DumpRequest(c.Request, false)
				if brokenPipe {
					zap.L().Error(c.Request.URL.Path,
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
					)
					// If the connection is dead, we can't write a status to it.
					c.Error(err.(error)) // nolint: errcheck
					c.Abort()
					return
				}

				if stack {
					zap.L().Error("[Recovery from panic]",
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
						zap.String("stack", string(debug.Stack())),
					)
				} else {
					zap.L().Error("[Recovery from panic]",
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
					)
				}
				c.AbortWithStatus(http.StatusInternalServerError)
			}
		}()
		c.Next()
	}
}
