package routers

import (
	"github.com/gin-gonic/gin"
)

func initAdminPublicRouterGroup(r *gin.Engine) {
	adminPublicGroup := r.Group("/admin/api/public")

	// 注册
	adminPublicGroup.GET("/sign_up", func(c *gin.Context) {})

	// 登录
	adminPublicGroup.GET("/sign_in", func(c *gin.Context) {})
}
