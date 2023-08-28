package routers

import (
	"api.aifuxi.cool/middleware"
	"github.com/gin-gonic/gin"
)

func initAdminAuthRouterGroup(r *gin.Engine) {
	adminAuthGroup := r.Group("/admin/api/auth")

	adminAuthGroup.Use(middleware.JwtAuth())

	// ====================== 文章标签 ======================
	// 获取文章标签列表
	adminAuthGroup.GET("/tags")

	// 创建文章标签
	adminAuthGroup.POST("/tags")

	// 根据 ID 获取文章标签
	adminAuthGroup.GET("/tags/:id")

	// 根据 ID 更新文章标签
	adminAuthGroup.PUT("/tags/:id")

	// 根据 ID 删除文章标签
	adminAuthGroup.DELETE("/tags/:id")

	// ====================== 文章 ======================
	// 获取文章列表
	adminAuthGroup.GET("/articles")

	// 创建文章
	adminAuthGroup.POST("/articles")

	// 根据 ID 获取文章
	adminAuthGroup.GET("/articles/:id")

	// 根据 ID 更新文章
	adminAuthGroup.PUT("/articles/:id")

	// 根据 ID 删除文章
	adminAuthGroup.DELETE("/articles/:id")

	// ====================== 用户 ======================
	// 获取用户列表
	adminAuthGroup.GET("/users")

	// 创建用户
	adminAuthGroup.POST("/users")

	// 根据 ID 获取用户
	adminAuthGroup.GET("/users/:id")

	// 根据 ID 更新用户
	adminAuthGroup.PUT("/users/:id")

	// 根据 ID 删除用户
	adminAuthGroup.DELETE("/users/:id")

	// ====================== 上传文件 ======================
	// 上传文件
	adminAuthGroup.POST("/uploads")
}
