package routers

import (
	"github.com/aifuxi/aifuxi_cool_api/controller"
	"github.com/gin-gonic/gin"
)

func initAdminAuthRouterGroup(r *gin.Engine) {
	adminAuthGroup := r.Group("/admin/api/auth")

	// 获取文章标签列表
	adminAuthGroup.GET("/tags", controller.GetTags)

	// 创建文章标签
	adminAuthGroup.POST("/tags", controller.CreateTag)

	// 根据 ID 获取文章标签
	adminAuthGroup.GET("/tags/:id", controller.GetTagByID)

	// 根据 ID 删除文章标签
	adminAuthGroup.DELETE("/tags/:id", controller.DeleteTagByID)
}
