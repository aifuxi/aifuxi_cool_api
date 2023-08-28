package api

import (
	"path/filepath"

	"api.aifuxi.cool/db/orm"
	"github.com/gin-gonic/gin"
)

type Server struct {
	router *gin.Engine
	store  orm.Store
}

func NewServer(store orm.Store) *Server {
	server := &Server{
		store: store,
	}
	router := gin.Default()

	// 配置静态文件服务路径
	rootPath := filepath.Join("uploads")
	router.Static("/uploads", rootPath)

	adminAuthApi := router.Group("/adminapi/auth")
	{
		adminAuthApi.GET("/users", server.ListUsers)
	}

	server.router = router

	return server
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}
