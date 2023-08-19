package routers

import (
	"github.com/aifuxi/aifuxi_cool_api/controller"
	"github.com/gin-gonic/gin"
)

func initAdminPublicRouterGroup(r *gin.Engine) {
	adminPublicGroup := r.Group("/admin/api/public")

	// 注册
	adminPublicGroup.POST("/sign_up", controller.SignUp)

	// 登录
	adminPublicGroup.POST("/sign_in", controller.SignIn)
}
