package controller

import (
	"fmt"
	"path/filepath"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

var baseUploadDir = "uploads"

func UploadFile(c *gin.Context) {
	// 单文件
	file, _ := c.FormFile("file")
	filename := fmt.Sprintf("%d-%s", time.Now().Unix(), strings.ToLower(file.Filename))
	dst := filepath.Join(baseUploadDir, filename)
	c.SaveUploadedFile(file, dst)

	fileUrl := fmt.Sprintf("http://%s/%s/%s", c.Request.Host, baseUploadDir, filename)
	ResponseOk(c, fileUrl)
}
