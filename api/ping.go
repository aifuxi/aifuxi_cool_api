package api

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (server *Server) Pong(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, errorResponse(errors.New("test error")))
}
