package main

import (
	"log"

	"github.com/aifuxi/aifuxi_cool_api/models"
	"github.com/aifuxi/aifuxi_cool_api/routers"
	"github.com/aifuxi/aifuxi_cool_api/settings"
	"github.com/aifuxi/aifuxi_cool_api/zlog"
)

var err error

func main() {
	err = settings.Init()
	if err != nil {
		log.Fatalf("初始化配置失败: %v", err)
	}

	zlog.Setup()

	defer zlog.L.Sync()

	err = models.Setup()
	if err != nil {
		log.Fatalf("连接 MySQL 失败: %v", err)
	}

	r := routers.Setup()
	r.Run("localhost:9003")
}
