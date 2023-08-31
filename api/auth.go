package api

import (
	db "api.aifuxi.cool/db/orm"
	"api.aifuxi.cool/util"
	"errors"
	"github.com/gin-gonic/gin"
)

func (s *Server) SignUp(c *gin.Context) {
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

type signInRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type signInResponse struct {
	User        db.User `json:"user"`
	AccessToken string  `json:"access_token"`
}

func (s *Server) SignIn(c *gin.Context) {
	var req signInRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		responseFailWithErr(c, ResponseCodeBadRequest, err)
		return
	}

	user, err := s.store.GetUserByEmail(req.Email)
	if errors.Is(err, db.ErrUserExist) {
		responseFailWithErr(c, ResponseCodeInvalidEmailOrPassword, err)
		return
	}

	err = util.CheckPassword(user.Password, req.Password)
	if err != nil {
		responseFailWithErr(c, ResponseCodeInvalidEmailOrPassword, err)
		return
	}

	// 3. 生成 token
	token, err := util.GenToken(req.Email)
	if err != nil {
		responseFailWithErr(c, ResponseCodeInvalidEmailOrPassword, err)
		return
	}

	resp := signInResponse{
		User:        user,
		AccessToken: token,
	}

	responseOk(c, resp)
}
