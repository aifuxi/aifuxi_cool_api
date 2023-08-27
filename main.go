package main

import (
	"fmt"
	"log"

	"api.aifuxi.cool/api"
	"api.aifuxi.cool/dao/mysql"
	"api.aifuxi.cool/logger"
	"api.aifuxi.cool/settings"
)

func main() {
	err := Init()
	if err != nil {
		log.Fatalf("初始化失败: %v\n", err)
	}

	address := fmt.Sprintf("localhost:%d", settings.AppConfig.Port)
	server := api.NewServer()
	server.Start(address)
}

func Init() error {
	err := settings.Init()
	if err != nil {
		fmt.Printf("初始化配置失败: %v\n", err)
		return err
	}

	logger.Init()

	err = mysql.Init()
	if err != nil {
		fmt.Printf("连接 MySQL 失败: %v\n", err)
		return err
	}

	return nil
}
