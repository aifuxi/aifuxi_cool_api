package routers

import (
	"github.com/aifuxi/aifuxi_cool_api/controller"
	"github.com/gin-gonic/gin"
)

func initAdminAuthRouterGroup(r *gin.Engine) {
	adminAuthGroup := r.Group("/admin/api/auth")

	// ====================== 文章标签 ======================
	// 获取文章标签列表
	adminAuthGroup.GET("/tags", controller.GetTags)

	// 创建文章标签
	adminAuthGroup.POST("/tags", controller.CreateTag)

	// 根据 ID 获取文章标签
	adminAuthGroup.GET("/tags/:id", controller.GetTagByID)

	// 根据 ID 更新文章标签
	adminAuthGroup.PUT("/tags/:id", controller.UpdateTagByID)

	// 根据 ID 删除文章标签
	adminAuthGroup.DELETE("/tags/:id", controller.DeleteTagByID)

	// ====================== 用户 ======================
	// 获取用户列表
	adminAuthGroup.GET("/users", controller.GetUsers)

	// 创建用户
	adminAuthGroup.POST("/users", controller.CreateUser)

	// 根据 ID 获取用户
	adminAuthGroup.GET("/users/:id", controller.GetUserByID)

	// 根据 ID 更新用户
	adminAuthGroup.PUT("/users/:id", controller.UpdateUserByID)

	// 根据 ID 删除用户
	adminAuthGroup.DELETE("/users/:id", controller.DeleteUserByID)
}
