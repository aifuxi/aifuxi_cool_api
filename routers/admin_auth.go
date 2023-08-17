package routers

import (
	"github.com/aifuxi/aifuxi_cool_api/controller"
	"github.com/gin-gonic/gin"
)

func initAdminAuthRouterGroup(r *gin.Engine) {
	adminAuthGroup := r.Group("/admin/api/auth")

	adminAuthGroup.GET("/tags", controller.GetTags)
}
