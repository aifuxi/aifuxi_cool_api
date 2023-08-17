package controller

import (
	"github.com/aifuxi/aifuxi_cool_api/dao/mysql"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// 获取标签列表
func GetTags(c *gin.Context) {
	tags, err := mysql.GetTags()
	if err != nil {
		zap.L().Error("获取标签列表失败", zap.Error(err))
		ResponseErr(c, ServerError)
		return
	}

	ResponseOk(c, tags)
}
