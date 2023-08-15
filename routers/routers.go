package routers

import (
	"github.com/aifuxi/aifuxi_cool_api/models"
	"github.com/gin-gonic/gin"
)

func Setup() *gin.Engine {
	r := gin.New()

	r.Use(gin.Recovery(), gin.Logger())

	// adminPublicGroup := r.Group("/admin/api/public")
	adminAuthGroup := r.Group("/admin/api/auth")

	adminAuthGroup.GET("/tags", func(c *gin.Context) {
		var tags []models.Tag
		models.DB.Model(models.Tag{}).Find(&tags)

		c.JSON(200, gin.H{
			"data": tags,
			"msg":  "success",
		})
	})

	return r
}
