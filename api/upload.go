package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"path/filepath"
	"strings"
	"time"
)

const uploadField = "file"

func (s *Server) UploadFile(c *gin.Context) {
	// 单文件
	file, _ := c.FormFile(uploadField)
	filename := fmt.Sprintf("%d-%s", time.Now().Unix(), strings.ToLower(file.Filename))
	dst := filepath.Join(baseUploadDir, filename)
	err := c.SaveUploadedFile(file, dst)
	if err != nil {
		responseFail(c, ResponseCodeUploadFileFailed)
	}

	fileUrl := fmt.Sprintf("http://%s/%s/%s", c.Request.Host, baseUploadDir, filename)
	responseOk(c, fileUrl)
}
