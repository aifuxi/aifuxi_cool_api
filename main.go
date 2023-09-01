package main

import (
	"api.aifuxi.cool/api"
	"api.aifuxi.cool/logger"
	"api.aifuxi.cool/settings"
	"go.uber.org/zap"
	"log"
)

func main() {
	err := settings.Init()
	if err != nil {
		log.Fatalln("init settings error: ", err)
	}

	logger.Init()

	server, err := api.NewServer()
	if err != nil {
		zap.L().Fatal("new server error: ", zap.Error(err))
	}

	err = server.Start(settings.AppConfig.Addr)
	if err != nil {
		zap.L().Fatal("start server error: ", zap.Error(err))
	}
}
