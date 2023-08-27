package controller

import (
	"api.aifuxi.cool/dto"
	"api.aifuxi.cool/service"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func PublicGetTags(c *gin.Context) {
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
