package routers

import (
	"api.aifuxi.cool/controller"
	"api.aifuxi.cool/middleware"
	"github.com/gin-gonic/gin"
)

func initAdminAuthRouterGroup(r *gin.Engine) {
	adminAuthGroup := r.Group("/admin/api/auth")

	adminAuthGroup.Use(middleware.JwtAuth())

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

	// ====================== 文章 ======================
	// 获取文章列表
	adminAuthGroup.GET("/articles", controller.GetArticles)

	// 创建文章
	adminAuthGroup.POST("/articles", controller.CreateArticle)

	// 根据 ID 获取文章
	adminAuthGroup.GET("/articles/:id", controller.GetArticleByID)

	// 根据 ID 更新文章
	adminAuthGroup.PUT("/articles/:id", controller.UpdateArticleByID)

	// 根据 ID 删除文章
	adminAuthGroup.DELETE("/articles/:id", controller.DeleteArticleByID)

	// ====================== 用户 ======================
	// 获取用户列表
	adminAuthGroup.GET("/users", controller.GetUsers)

	// 获取当前登录用户的信息
	adminAuthGroup.GET("/users/profile", controller.GetUserProfile)

	// 创建用户
	adminAuthGroup.POST("/users", controller.CreateUser)

	// 根据 ID 获取用户
	adminAuthGroup.GET("/users/:id", controller.GetUserByID)

	// 根据 ID 更新用户
	adminAuthGroup.PUT("/users/:id", controller.UpdateUserByID)

	// 根据 ID 删除用户
	adminAuthGroup.DELETE("/users/:id", controller.DeleteUserByID)
}
