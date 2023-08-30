package api

import (
	db "api.aifuxi.cool/db/orm"
	"github.com/gin-gonic/gin"
)

type Server struct {
	router *gin.Engine
	store  db.Store
}

func NewServer() (*Server, error) {
	store, err := db.NewStore()
	if err != nil {
		return nil, err
	}

	server := &Server{
		store: store,
	}

	server.setupRouter()

	return server, nil
}

func (s *Server) setupRouter() {
	router := gin.Default()

	adminAuthApi := router.Group("/adminapi/auth")
	{
		adminAuthApi.GET("/users", s.ListUsers)
		adminAuthApi.POST("/users", s.CreateUser)
		adminAuthApi.GET("/users/:id", s.GetUser)
		adminAuthApi.PUT("/users/:id", s.UpdateUser)
		adminAuthApi.DELETE("/users/:id", s.DeleteUser)
	}

	s.router = router
}

func (s *Server) Start(addr string) error {

	return s.router.Run(addr)
}
