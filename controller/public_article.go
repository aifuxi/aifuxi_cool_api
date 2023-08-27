package controller

import (
	"errors"

	"api.aifuxi.cool/dto"
	"api.aifuxi.cool/service"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

func PublicGetArticles(c *gin.Context) {
	var getArticlesDTO dto.GetArticlesDTO

	// 解析分页参数
	if err := c.ShouldBindQuery(&getArticlesDTO); err != nil {
		// 获取validator.ValidationErrors类型的errors
		errs, ok := err.(validator.ValidationErrors)
		if ok {
			ResponseErr(c, InvalidParams, errs.Error())
			return
		}

		ResponseErr(c, InvalidParams, nil)
		return
	}

	articles, total, err := service.GetArticles(getArticlesDTO)
	if err != nil {
		ResponseErr(c, ServerError, nil)
		return
	}

	ResponseOkWithTotal(c, articles, total)
}

func PublicGetArticleByFriendlyUrl(c *gin.Context) {
	friendlyUrl := c.Param("friendly_url")

	article, err := service.GetArticleByFriendlyUrl(friendlyUrl)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ResponseErr(c, InvalidParams, gorm.ErrRecordNotFound.Error())
			return
		}

		ResponseErr(c, ServerError, nil)
		return
	}

	ResponseOk(c, article)
}
