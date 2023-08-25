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

func GetTags(c *gin.Context) {
	var getTagsDTO dto.GetTagsDTO

	// 解析分页参数
	if err := c.ShouldBindQuery(&getTagsDTO); err != nil {
		// 获取validator.ValidationErrors类型的errors
		errs, ok := err.(validator.ValidationErrors)
		if ok {
			ResponseErr(c, InvalidParams, errs.Error())
			return
		}

		ResponseErr(c, InvalidParams, nil)
		return
	}

	tags, total, err := service.GetTags(getTagsDTO)
	if err != nil {
		ResponseErr(c, ServerError, nil)
		return
	}

	ResponseOkWithTotal(c, tags, total)
}

func CreateTag(c *gin.Context) {
	var createTagDTO dto.CreateTagDTO

	if err := c.ShouldBindJSON(&createTagDTO); err != nil {
		// 获取validator.ValidationErrors类型的errors
		errs, ok := err.(validator.ValidationErrors)
		if ok {
			ResponseErr(c, InvalidParams, errs.Error())
			return
		}

		ResponseErr(c, InvalidParams, nil)
		return
	}

	tag, err := service.CreateTag(createTagDTO)
	if err != nil {
		ResponseErr(c, InvalidParams, err.Error())
		return
	}

	ResponseOk(c, tag)
}

func GetTagByID(c *gin.Context) {
	idStr := c.Param("id")

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		ResponseErr(c, InvalidParams, "invalid tag id")
		return
	}

	tag, err := service.GetTagByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ResponseErr(c, InvalidParams, gorm.ErrRecordNotFound.Error())
			return
		}

		ResponseErr(c, ServerError, nil)
		return
	}

	ResponseOk(c, tag)
}

func UpdateTagByID(c *gin.Context) {
	idStr := c.Param("id")

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		ResponseErr(c, InvalidParams, "invalid tag id")
		return
	}

	var updateTagDTO dto.UpdateTagDTO
	if err := c.ShouldBindJSON(&updateTagDTO); err != nil {
		// 获取validator.ValidationErrors类型的errors
		errs, ok := err.(validator.ValidationErrors)
		if ok {
			ResponseErr(c, InvalidParams, errs.Error())
			return
		}

		ResponseErr(c, InvalidParams, nil)
		return
	}

	err = service.UpdateTagByID(id, updateTagDTO)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ResponseErr(c, InvalidParams, gorm.ErrRecordNotFound.Error())
			return
		}

		ResponseErr(c, ServerError, nil)
		return
	}

	ResponseOk(c, nil)
}

func DeleteTagByID(c *gin.Context) {
	idStr := c.Param("id")

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		ResponseErr(c, InvalidParams, "invalid tag id")
		return
	}

	err = service.DeleteTagByID(id)
	if err != nil {
		if errors.Is(err, myerror.ErrorTagNotFound) {
			ResponseErr(c, InvalidParams, myerror.ErrorTagNotFound.Error())
			return
		}

		ResponseErr(c, ServerError, nil)
		return
	}

	ResponseOk(c, nil)
}
