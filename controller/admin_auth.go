package controller

import (
	"api.aifuxi.cool/dto"
	"api.aifuxi.cool/service"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

func SignUp(c *gin.Context) {
	signUpDTO := new(dto.SignUpDTO)

	if err := c.ShouldBindJSON(signUpDTO); err != nil {
		// 获取validator.ValidationErrors类型的errors
		errs, ok := err.(validator.ValidationErrors)
		if ok {
			zap.L().Error("controller.SignUp: validation params failed", zap.Error(errs))
			ResponseErrWithMsg(c, InvalidParams, errs.Error())
			return
		}

		zap.L().Error("controller.SignUp: invalid params", zap.Error(err))
		ResponseErr(c, InvalidParams)
		return
	}

	user, err := service.SignUp(signUpDTO)
	if err != nil {
		zap.L().Error("controller.SignUp: create user error", zap.Error(err))
		ResponseErrWithMsg(c, InvalidParams, err.Error())
		return
	}

	ResponseOk(c, user)
}

func SignIn(c *gin.Context) {
	signInDto := new(dto.SignInDTO)

	if err := c.ShouldBindJSON(signInDto); err != nil {
		// 获取validator.ValidationErrors类型的errors
		errs, ok := err.(validator.ValidationErrors)
		if ok {
			zap.L().Error("controller.SignIn: validation params failed", zap.Error(errs))
			ResponseErrWithMsg(c, InvalidParams, errs.Error())
			return
		}

		zap.L().Error("controller.SignIn: invalid params", zap.Error(err))
		ResponseErr(c, InvalidParams)
		return
	}

	token, err := service.SignIn(signInDto)
	if err != nil {
		zap.L().Error("controller.SignIn: sign in error", zap.Error(err))
		ResponseErrWithMsg(c, InvalidParams, err.Error())
		return
	}

	ResponseOk(c, token)
}
