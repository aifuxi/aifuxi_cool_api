package controller

import (
	"errors"
	"strconv"

	"api.aifuxi.cool/dto"
	"api.aifuxi.cool/myerror"
	"api.aifuxi.cool/service"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func GetUsers(c *gin.Context) {
	users, err := service.GetUsers()
	if err != nil {
		zap.L().Error("controller.GetUsers: get users error", zap.Error(err))
		ResponseErr(c, ServerError)
		return
	}

	ResponseOk(c, users)
}

func CreateUser(c *gin.Context) {
	createUserDTO := new(dto.CreateUserDTO)

	if err := c.ShouldBindJSON(createUserDTO); err != nil {
		// 获取validator.ValidationErrors类型的errors
		errs, ok := err.(validator.ValidationErrors)
		if ok {
			zap.L().Error("controller.CreateUser: validation params failed", zap.Error(errs))
			ResponseErrWithMsg(c, InvalidParams, errs.Error())
			return
		}

		zap.L().Error("controller.CreateUser: invalid params", zap.Error(err))
		ResponseErr(c, InvalidParams)
		return
	}

	user, err := service.CreateUser(createUserDTO)
	if err != nil {
		zap.L().Error("controller.CreateUser: create user error", zap.Error(err))
		ResponseErrWithMsg(c, InvalidParams, err.Error())
		return
	}

	ResponseOk(c, user)
}

func GetUserByID(c *gin.Context) {
	idStr := c.Param("id")

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		zap.L().Error("controller.GetUserByID: invalid user id", zap.String("error", "invalid user id"))
		ResponseErrWithMsg(c, InvalidParams, "invalid user id")
		return
	}

	user, err := service.GetUserByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			zap.L().Error("controller.GetUserByID: user not found", zap.Error(gorm.ErrRecordNotFound))
			ResponseErrWithMsg(c, InvalidParams, gorm.ErrRecordNotFound.Error())
			return
		}

		zap.L().Error("controller.GetUserByID: get user error", zap.Error(err))
		ResponseErr(c, ServerError)
		return
	}

	ResponseOk(c, user)
}

func UpdateUserByID(c *gin.Context) {
	idStr := c.Param("id")

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		zap.L().Error("controller.UpdateUserByID: invalid user id", zap.String("error", "invalid user id"))
		ResponseErrWithMsg(c, InvalidParams, "invalid user id")
		return
	}

	updateUserDTO := new(dto.UpdateUserDTO)
	if err := c.ShouldBindJSON(updateUserDTO); err != nil {
		// 获取validator.ValidationErrors类型的errors
		errs, ok := err.(validator.ValidationErrors)
		if ok {
			zap.L().Error("controller.UpdateUserByID: validation params failed", zap.Error(errs))
			ResponseErrWithMsg(c, InvalidParams, errs.Error())
			return
		}

		zap.L().Error("controller.UpdateUserByID: invalid params", zap.Error(err))
		ResponseErr(c, InvalidParams)
		return
	}

	err = service.UpdateUserByID(id, updateUserDTO)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			zap.L().Error("controller.UpdateUserByID: user not found", zap.Error(gorm.ErrRecordNotFound))
			ResponseErrWithMsg(c, InvalidParams, gorm.ErrRecordNotFound.Error())
			return
		}

		zap.L().Error("controller.UpdateUserByID: get user error", zap.Error(err))
		ResponseErr(c, ServerError)
		return
	}

	ResponseOk(c, nil)
}

func DeleteUserByID(c *gin.Context) {
	idStr := c.Param("id")

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		zap.L().Error("controller.DeleteUserByID: invalid user id", zap.String("error", "invalid user id"))
		ResponseErrWithMsg(c, InvalidParams, "invalid user id")
		return
	}

	err = service.DeleteUserByID(id)
	if err != nil {
		if errors.Is(err, myerror.ErrorUserNotFound) {
			zap.L().Error("controller.DeleteUserByID: user not found", zap.Error(myerror.ErrorUserNotFound))
			ResponseErrWithMsg(c, InvalidParams, myerror.ErrorUserNotFound)
			return
		}

		zap.L().Error("controller.DeleteUserByID: delete user error", zap.Error(err))
		ResponseErr(c, ServerError)
		return
	}

	ResponseOk(c, nil)
}
