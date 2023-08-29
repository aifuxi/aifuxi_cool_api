package main

import (
	"api.aifuxi.cool/api"
	"log"
)

func main() {
	server, err := api.NewServer()
	if err != nil {
		log.Fatalln("new server error: ", err)
	}

	err = server.Start("localhost:9003")
	if err != nil {
		log.Fatalln("start server error: ", err)
	}
}
