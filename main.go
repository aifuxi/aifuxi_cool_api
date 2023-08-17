package main

import (
	"fmt"
	"log"

	"github.com/aifuxi/aifuxi_cool_api/dao/mysql"
	"github.com/aifuxi/aifuxi_cool_api/logger"
	"github.com/aifuxi/aifuxi_cool_api/routers"
	"github.com/aifuxi/aifuxi_cool_api/settings"
)

var err error

func main() {
	err = settings.Init()
	if err != nil {
		log.Fatalf("初始化配置失败: %v", err)
	}

	logger.Init()

	err = mysql.Init()
	if err != nil {
		log.Fatalf("连接 MySQL 失败: %v", err)
	}

	r := routers.Init()

	addr := fmt.Sprintf("localhost:%d", settings.AppConfig.Port)
	r.Run(addr)
}
