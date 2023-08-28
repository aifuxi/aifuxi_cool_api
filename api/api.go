package api

import (
	"path/filepath"

	"api.aifuxi.cool/middleware"
	"api.aifuxi.cool/settings"
	"github.com/gin-gonic/gin"
)

type Server struct {
	router *gin.Engine
}

func NewServer() *Server {
	server := &Server{}
	r := gin.New()

	gin.SetMode(settings.AppConfig.Mode)

	r.Use(middleware.GinLogger(), middleware.GinRecovery(true))

	rootPath := filepath.Join("uploads")
	r.Static("/uploads", rootPath)

	server.router = r

	return server
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}
