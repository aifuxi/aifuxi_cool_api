package controller

import (
	"api.aifuxi.cool/dto"
	"api.aifuxi.cool/service"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func SignUp(c *gin.Context) {
	var signUpDTO dto.SignUpDTO

	if err := c.ShouldBindJSON(&signUpDTO); err != nil {
		// 获取validator.ValidationErrors类型的errors
		errs, ok := err.(validator.ValidationErrors)
		if ok {
			ResponseErr(c, InvalidParams, errs.Error())
			return
		}

		ResponseErr(c, InvalidParams, nil)
		return
	}

	user, err := service.SignUp(signUpDTO)
	if err != nil {
		ResponseErr(c, InvalidParams, err.Error())
		return
	}

	ResponseOk(c, user)
}

func SignIn(c *gin.Context) {
	var signInDto dto.SignInDTO

	if err := c.ShouldBindJSON(&signInDto); err != nil {
		// 获取validator.ValidationErrors类型的errors
		errs, ok := err.(validator.ValidationErrors)
		if ok {
			ResponseErr(c, InvalidParams, errs.Error())
			return
		}

		ResponseErr(c, InvalidParams, nil)
		return
	}

	token, user, err := service.SignIn(signInDto)
	if err != nil {
		ResponseErr(c, InvalidParams, err.Error())
		return
	}

	ResponseOk(c, gin.H{
		"access_token": token,
		"user":         user,
	})
}
