package main

import (
	"api.aifuxi.cool/api"
	"api.aifuxi.cool/settings"
	"log"
)

func main() {
	err := settings.Init()
	if err != nil {
		log.Fatalln("init settings error: ", err)
	}

	server, err := api.NewServer()
	if err != nil {
		log.Fatalln("new server error: ", err)
	}

	err = server.Start("localhost:9003")
	if err != nil {
		log.Fatalln("start server error: ", err)
	}
}
