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

func GetArticles(c *gin.Context) {
	var getArticlesDTO dto.GetArticlesDTO

	// 解析分页参数
	if err := c.ShouldBindQuery(&getArticlesDTO); err != nil {
		// 获取validator.ValidationErrors类型的errors
		errs, ok := err.(validator.ValidationErrors)
		if ok {
			ResponseErrWithMsg(c, InvalidParams, errs.Error())
			return
		}

		ResponseErr(c, InvalidParams)
		return
	}

	articles, total, err := service.GetArticles(getArticlesDTO)
	if err != nil {
		ResponseErr(c, ServerError)
		return
	}

	ResponseOkWithTotal(c, articles, total)
}

func CreateArticle(c *gin.Context) {
	var createArticleDTO dto.CreateArticleDTO

	if err := c.ShouldBindJSON(&createArticleDTO); err != nil {
		// 获取validator.ValidationErrors类型的errors
		errs, ok := err.(validator.ValidationErrors)
		if ok {
			ResponseErrWithMsg(c, InvalidParams, errs.Error())
			return
		}

		ResponseErr(c, InvalidParams)
		return
	}

	article, err := service.CreateArticle(createArticleDTO)
	if err != nil {
		ResponseErrWithMsg(c, InvalidParams, err.Error())
		return
	}

	ResponseOk(c, article)
}

func GetArticleByID(c *gin.Context) {
	idStr := c.Param("id")

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		ResponseErrWithMsg(c, InvalidParams, "invalid article id")
		return
	}

	article, err := service.GetArticleByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ResponseErrWithMsg(c, InvalidParams, gorm.ErrRecordNotFound.Error())
			return
		}

		ResponseErr(c, ServerError)
		return
	}

	ResponseOk(c, article)
}

func UpdateArticleByID(c *gin.Context) {
	idStr := c.Param("id")

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		ResponseErrWithMsg(c, InvalidParams, "invalid article id")
		return
	}

	var updateArticleDTO dto.UpdateArticleDTO
	if err := c.ShouldBindJSON(&updateArticleDTO); err != nil {
		// 获取validator.ValidationErrors类型的errors
		errs, ok := err.(validator.ValidationErrors)
		if ok {
			ResponseErrWithMsg(c, InvalidParams, errs.Error())
			return
		}

		ResponseErr(c, InvalidParams)
		return
	}

	err = service.UpdateArticleByID(id, updateArticleDTO)
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

func DeleteArticleByID(c *gin.Context) {
	idStr := c.Param("id")

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		ResponseErrWithMsg(c, InvalidParams, "invalid article id")
		return
	}

	err = service.DeleteArticleByID(id)
	if err != nil {
		if errors.Is(err, myerror.ErrorArticleNotFound) {
			ResponseErrWithMsg(c, InvalidParams, myerror.ErrorArticleNotFound)
			return
		}

		ResponseErr(c, ServerError)
		return
	}

	ResponseOk(c, nil)
}
