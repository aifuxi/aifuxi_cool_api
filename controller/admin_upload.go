package controller

import (
	"fmt"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

var baseUploadDir = "uploads"

func UploadFile(c *gin.Context) {
	// 单文件
	file, _ := c.FormFile("file")
	dst := filepath.Join(baseUploadDir, file.Filename)
	c.SaveUploadedFile(file, dst)

	fileUrl := fmt.Sprintf("/%s/%s", baseUploadDir, file.Filename)

	ResponseOk(c, fileUrl)
}
