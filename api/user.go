package api

import (
	db "api.aifuxi.cool/db/orm"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

type listUsersRequest struct {
	Email    string `form:"email"`
	Nickname string `form:"nickname"`

	Page     int    `form:"page" binding:"required,gte=1"`
	PageSize int    `form:"page_size" binding:"required,gte=10,lte=50"`
	Order    string `form:"order" binding:"required,oneof=asc desc"`
	OrderBy  string `form:"order_by" binding:"required,oneof=created_at updated_at"`
}

func (s *Server) ListUsers(c *gin.Context) {
	var req listUsersRequest

	if err := c.ShouldBindQuery(&req); err != nil {
		responseFailWithErr(c, ResponseCodeBadRequest, err)
		return
	}

	arg := db.ListUsersParams{
		Email:    req.Email,
		Nickname: req.Nickname,
		Page:     req.Page,
		PageSize: req.PageSize,
		Order:    req.Order,
		OrderBy:  req.OrderBy,
	}

	users, total, err := s.store.ListUsers(arg)
	if err != nil {
		responseFailWithErr(c, ResponseCodeFail, err)
		return
	}

	responseOkWithTotal(c, users, total)
}

type createUserRequest struct {
	Nickname   string `json:"nickname" binding:"required"`
	Avatar     string `json:"avatar"`
	Email      string `json:"email" binding:"required,email"`
	Password   string `json:"password" binding:"required"`
	RePassword string `json:"re_password" binding:"required,eqfield=Password"`
}

func (s *Server) CreateUser(c *gin.Context) {
	var req createUserRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		responseFailWithErr(c, ResponseCodeBadRequest, err)
		return
	}

	arg := db.CreateUserParams{
		Nickname: req.Nickname,
		Avatar:   req.Avatar,
		Email:    req.Email,
		Password: req.Password,
	}

	user, err := s.store.CreateUser(arg)
	if err != nil {
		if errors.Is(err, db.ErrUserExist) {
			responseFailWithErr(c, ResponseCodeBadRequest, err)
			return
		}

		responseFailWithErr(c, ResponseCodeFail, err)
		return
	}

	responseOk(c, user)
}

func (s *Server) GetUser(c *gin.Context) {

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		responseFailWithErr(c, ResponseCodeBadRequest, err)
		return
	}

	var user db.User
	user, err = s.store.GetUserByID(id)
	if err != nil {
		if errors.Is(err, db.ErrUserNotFound) {
			responseFailWithErr(c, ResponseCodeBadRequest, err)
			return
		}
		responseFailWithErr(c, ResponseCodeFail, err)
		return
	}

	responseOk(c, user)
}

type updateUserRequest struct {
	Nickname   string `json:"nickname"`
	Avatar     string `json:"avatar"`
	Email      string `json:"email" binding:"email"`
	Password   string `json:"password"`
	RePassword string `json:"re_password" binding:"required_with=Password"`
}

func (s *Server) UpdateUser(c *gin.Context) {
	var req updateUserRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		responseFailWithErr(c, ResponseCodeBadRequest, err)
		return
	}

	if len(req.Password) > 0 && req.Password != req.RePassword {
		responseFailWithErr(c, ResponseCodeBadRequest, fmt.Errorf("两次输入密码不一致"))
		return
	}

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		responseFailWithErr(c, ResponseCodeBadRequest, err)
		return
	}

	arg := db.UpdateUserParams{
		Nickname: req.Nickname,
		Avatar:   req.Avatar,
		Email:    req.Email,
		Password: req.Password,
	}

	err = s.store.UpdateUser(id, arg)
	if err != nil {
		if errors.Is(err, db.ErrUserNotFound) {
			responseFailWithErr(c, ResponseCodeBadRequest, err)
			return
		}
		responseFailWithErr(c, ResponseCodeFail, err)
		return
	}

	responseOk(c, nil)
}

func (s *Server) DeleteUser(c *gin.Context) {

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		responseFailWithErr(c, ResponseCodeBadRequest, err)
		return
	}

	err = s.store.DeleteUserByID(id)
	if err != nil {
		if errors.Is(err, db.ErrUserNotFound) {
			responseFailWithErr(c, ResponseCodeBadRequest, err)
			return
		}
		responseFailWithErr(c, ResponseCodeFail, err)
		return
	}

	responseOk(c, nil)
}
