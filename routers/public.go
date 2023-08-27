package routers

import (
	"api.aifuxi.cool/controller"
	"github.com/gin-gonic/gin"
)

func initPublicRouterGroup(r *gin.Engine) {
	publicGroup := r.Group("/public/api")

	// 获取文章列表
	publicGroup.GET("/articles", controller.PublicGetArticles)

	// 根据friendlyUrl获取文章
	publicGroup.GET("/articles/:friendly_url", controller.PublicGetArticleByFriendlyUrl)

	// 获取标签列表
	publicGroup.GET("/tags", controller.PublicGetTags)
}
