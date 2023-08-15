package test

import (
	"fmt"
	"log"
	"testing"

	"github.com/aifuxi/aifuxi_cool_api/models"
)

func Test(t *testing.T) {
	err := models.Setup()

	if err != nil {
		log.Fatalf("连接 MySQL 失败: %v", err)
	}

	var tags []models.Tag
	models.DB.Model(models.Tag{}).Find(&tags)

	fmt.Printf("tags: %v\n", tags)
}
