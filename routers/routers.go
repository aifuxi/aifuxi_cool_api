package routers

import (
	"log"

	"github.com/aifuxi/aifuxi_cool_api/models"
	"github.com/aifuxi/aifuxi_cool_api/zlog"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func Setup() *gin.Engine {
	r := gin.New()

	r.Use(gin.Recovery(), gin.Logger())

	// adminPublicGroup := r.Group("/admin/api/public")
	adminAuthGroup := r.Group("/admin/api/auth")

	adminAuthGroup.GET("/tags", func(c *gin.Context) {
		var tags []models.Tag
		models.DB.Model(models.Tag{}).Find(&tags)

		log.Println(
			"get tags嘻嘻 ✔✔")
		zlog.L.Info("嘻嘻嘻",
			zap.String("reason", "这是测试日志"),
		)

		c.JSON(200, gin.H{
			"data": tags,
			"msg":  "success",
		})
	})

	return r
}
