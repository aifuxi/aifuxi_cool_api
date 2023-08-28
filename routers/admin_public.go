package routers

import (
	"github.com/gin-gonic/gin"
)

func initAdminPublicRouterGroup(r *gin.Engine) {
	adminPublicGroup := r.Group("/admin/api/public")

	// 注册
	adminPublicGroup.POST("/sign_up")

	// 登录
	adminPublicGroup.POST("/sign_in")
}
