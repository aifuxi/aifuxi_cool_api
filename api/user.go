package api

import (
	"api.aifuxi.cool/db/orm"
	"github.com/gin-gonic/gin"
)

type listUsersRequest struct {
	Page     int `form:"page" binding:"required"`
	PageSize int `form:"page_size" binding:"required"`
}

func (server *Server) ListUsers(ctx *gin.Context) {
	var req listUsersRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(400, gin.H{
			"error": err,
		})
		return
	}

	arg := orm.ListUsersParams{
		PaginationParams: orm.PaginationParams{
			Page:     req.Page,
			PageSize: req.PageSize,
		},
	}

	users, total, err := server.store.ListUsers(arg)
	if err != nil {
		ctx.JSON(500, gin.H{
			"users": nil,
			"total": 0,
			"error": err,
		})
		return
	}

	ctx.JSON(200, gin.H{
		"users": users,
		"total": total,
	})
}
