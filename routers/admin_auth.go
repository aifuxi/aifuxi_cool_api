package routers

import (
	"github.com/aifuxi/aifuxi_cool_api/models"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func initAdminAuthRouterGroup(r *gin.Engine) {
	adminAuthGroup := r.Group("/admin/api/auth")

	adminAuthGroup.GET("/tags", func(c *gin.Context) {
		var tags []models.Tag
		models.DB.Model(models.Tag{}).Find(&tags)

		zap.L().Debug("test global zap logger", zap.String("msg", "test zap"))

		c.JSON(200, gin.H{
			"data": tags,
			"msg":  "success",
		})
	})
}
