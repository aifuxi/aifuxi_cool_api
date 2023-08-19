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

func GetTags(c *gin.Context) {
	tags, err := service.GetTags()
	if err != nil {
		zap.L().Error("controller.GetTags: get tags error", zap.Error(err))
		ResponseErr(c, ServerError)
		return
	}

	ResponseOk(c, tags)
}

func CreateTag(c *gin.Context) {
	createTagDTO := new(dto.CreateTagDTO)

	if err := c.ShouldBindJSON(createTagDTO); err != nil {
		// 获取validator.ValidationErrors类型的errors
		errs, ok := err.(validator.ValidationErrors)
		if ok {
			zap.L().Error("controller.CreateTag: validation params failed", zap.Error(errs))
			ResponseErrWithMsg(c, InvalidParams, errs.Error())
			return
		}

		zap.L().Error("controller.CreateTag: invalid params", zap.Error(err))
		ResponseErr(c, InvalidParams)
		return
	}

	tag, err := service.CreateTag(createTagDTO)
	if err != nil {
		zap.L().Error("controller.CreateTag: create tag error", zap.Error(err))
		ResponseErrWithMsg(c, InvalidParams, err.Error())
		return
	}

	ResponseOk(c, tag)
}

func GetTagByID(c *gin.Context) {
	idStr := c.Param("id")

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		zap.L().Error("controller.GetTagByID: invalid tag id", zap.String("error", "invalid tag id"))
		ResponseErrWithMsg(c, InvalidParams, "invalid tag id")
		return
	}

	tag, err := service.GetTagByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			zap.L().Error("controller.GetTagByID: tag not found", zap.Error(gorm.ErrRecordNotFound))
			ResponseErrWithMsg(c, InvalidParams, gorm.ErrRecordNotFound.Error())
			return
		}

		zap.L().Error("controller.GetTagByID: get tag error", zap.Error(err))
		ResponseErr(c, ServerError)
		return
	}

	ResponseOk(c, tag)
}

func UpdateTagByID(c *gin.Context) {
	idStr := c.Param("id")

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		zap.L().Error("controller.UpdateTagByID: invalid tag id", zap.String("error", "invalid tag id"))
		ResponseErrWithMsg(c, InvalidParams, "invalid tag id")
		return
	}

	updateTagDTO := new(dto.UpdateTagDTO)
	if err := c.ShouldBindJSON(updateTagDTO); err != nil {
		// 获取validator.ValidationErrors类型的errors
		errs, ok := err.(validator.ValidationErrors)
		if ok {
			zap.L().Error("controller.UpdateTagByID: validation params failed", zap.Error(errs))
			ResponseErrWithMsg(c, InvalidParams, errs.Error())
			return
		}

		zap.L().Error("controller.UpdateTagByID: invalid params", zap.Error(err))
		ResponseErr(c, InvalidParams)
		return
	}

	err = service.UpdateTagByID(id, updateTagDTO)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			zap.L().Error("controller.UpdateTagByID: tag not found", zap.Error(gorm.ErrRecordNotFound))
			ResponseErrWithMsg(c, InvalidParams, gorm.ErrRecordNotFound.Error())
			return
		}

		zap.L().Error("controller.UpdateTagByID: get tag error", zap.Error(err))
		ResponseErr(c, ServerError)
		return
	}

	ResponseOk(c, nil)
}

func DeleteTagByID(c *gin.Context) {
	idStr := c.Param("id")

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		zap.L().Error("controller.DeleteTagByID: invalid tag id", zap.String("error", "invalid tag id"))
		ResponseErrWithMsg(c, InvalidParams, "invalid tag id")
		return
	}

	err = service.DeleteTagByID(id)
	if err != nil {
		if errors.Is(err, myerror.ErrorTagNotFound) {
			zap.L().Error("controller.DeleteTagByID: tag not found", zap.Error(myerror.ErrorTagNotFound))
			ResponseErrWithMsg(c, InvalidParams, myerror.ErrorTagNotFound)
			return
		}

		zap.L().Error("controller.DeleteTagByID: delete tag error", zap.Error(err))
		ResponseErr(c, ServerError)
		return
	}

	ResponseOk(c, nil)
}
