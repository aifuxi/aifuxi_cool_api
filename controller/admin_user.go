package controller

import (
	"errors"
	"strconv"

	"api.aifuxi.cool/dto"
	"api.aifuxi.cool/myerror"
	"api.aifuxi.cool/service"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

func GetUsers(c *gin.Context) {
	var getUsersDTO dto.GetUsersDTO

	// 解析分页参数
	if err := c.ShouldBindQuery(&getUsersDTO); err != nil {
		// 获取validator.ValidationErrors类型的errors
		errs, ok := err.(validator.ValidationErrors)
		if ok {
			ResponseErrWithMsg(c, InvalidParams, errs.Error())
			return
		}

		ResponseErr(c, InvalidParams)
		return
	}

	users, total, err := service.GetUsers(getUsersDTO)
	if err != nil {
		ResponseErr(c, ServerError)
		return
	}

	ResponseOkWithTotal(c, users, total)
}

func GetUserProfile(c *gin.Context) {
	if email, exists := c.Get("email"); exists {
		user, err := service.GetUserProfile((email.(string)))
		if err != nil {
			ResponseErr(c, ServerError)
			return
		}

		ResponseOk(c, user)
	} else {
		ResponseErr(c, ServerError)
	}
}

func CreateUser(c *gin.Context) {
	var createUserDTO dto.CreateUserDTO

	if err := c.ShouldBindJSON(&createUserDTO); err != nil {
		// 获取validator.ValidationErrors类型的errors
		errs, ok := err.(validator.ValidationErrors)
		if ok {
			ResponseErrWithMsg(c, InvalidParams, errs.Error())
			return
		}

		ResponseErr(c, InvalidParams)
		return
	}

	user, err := service.CreateUser(createUserDTO)
	if err != nil {
		ResponseErrWithMsg(c, InvalidParams, err.Error())
		return
	}

	ResponseOk(c, user)
}

func GetUserByID(c *gin.Context) {
	idStr := c.Param("id")

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		ResponseErrWithMsg(c, InvalidParams, "invalid user id")
		return
	}

	user, err := service.GetUserByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ResponseErrWithMsg(c, InvalidParams, gorm.ErrRecordNotFound.Error())
			return
		}

		ResponseErr(c, ServerError)
		return
	}

	ResponseOk(c, user)
}

func UpdateUserByID(c *gin.Context) {
	idStr := c.Param("id")

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		ResponseErrWithMsg(c, InvalidParams, "invalid user id")
		return
	}

	var updateUserDTO dto.UpdateUserDTO
	if err := c.ShouldBindJSON(&updateUserDTO); err != nil {
		// 获取validator.ValidationErrors类型的errors
		errs, ok := err.(validator.ValidationErrors)
		if ok {
			ResponseErrWithMsg(c, InvalidParams, errs.Error())
			return
		}

		ResponseErr(c, InvalidParams)
		return
	}

	err = service.UpdateUserByID(id, updateUserDTO)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ResponseErrWithMsg(c, InvalidParams, gorm.ErrRecordNotFound.Error())
			return
		}

		ResponseErr(c, ServerError)
		return
	}

	ResponseOk(c, nil)
}

func DeleteUserByID(c *gin.Context) {
	idStr := c.Param("id")

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		ResponseErrWithMsg(c, InvalidParams, "invalid user id")
		return
	}

	err = service.DeleteUserByID(id)
	if err != nil {
		if errors.Is(err, myerror.ErrorUserNotFound) {
			ResponseErrWithMsg(c, InvalidParams, myerror.ErrorUserNotFound)
			return
		}

		ResponseErr(c, ServerError)
		return
	}

	ResponseOk(c, nil)
}
