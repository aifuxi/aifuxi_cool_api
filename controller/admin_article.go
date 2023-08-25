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

func GetArticles(c *gin.Context) {
	var getArticlesDTO dto.GetArticlesDTO

	// 解析分页参数
	if err := c.ShouldBindQuery(&getArticlesDTO); err != nil {
		// 获取validator.ValidationErrors类型的errors
		errs, ok := err.(validator.ValidationErrors)
		if ok {
			zap.L().Error("controller.GetArticles: validation params failed", zap.Error(errs))
			ResponseErrWithMsg(c, InvalidParams, errs.Error())
			return
		}

		zap.L().Error("controller.GetArticles: invalid params", zap.Error(err))
		ResponseErr(c, InvalidParams)
		return
	}

	articles, total, err := service.GetArticles(getArticlesDTO)
	if err != nil {
		zap.L().Error("controller.GetArticles: get articles error", zap.Error(err))
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
			zap.L().Error("controller.CreateArticle: validation params failed", zap.Error(errs))
			ResponseErrWithMsg(c, InvalidParams, errs.Error())
			return
		}

		zap.L().Error("controller.CreateArticle: invalid params", zap.Error(err))
		ResponseErr(c, InvalidParams)
		return
	}

	article, err := service.CreateArticle(createArticleDTO)
	if err != nil {
		zap.L().Error("controller.CreateArticle: create article error", zap.Error(err))
		ResponseErrWithMsg(c, InvalidParams, err.Error())
		return
	}

	ResponseOk(c, article)
}

func GetArticleByID(c *gin.Context) {
	idStr := c.Param("id")

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		zap.L().Error("controller.GetArticleByID: invalid article id", zap.String("error", "invalid article id"))
		ResponseErrWithMsg(c, InvalidParams, "invalid article id")
		return
	}

	article, err := service.GetArticleByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			zap.L().Error("controller.GetArticleByID: article not found", zap.Error(gorm.ErrRecordNotFound))
			ResponseErrWithMsg(c, InvalidParams, gorm.ErrRecordNotFound.Error())
			return
		}

		zap.L().Error("controller.GetArticleByID: get article error", zap.Error(err))
		ResponseErr(c, ServerError)
		return
	}

	ResponseOk(c, article)
}

func UpdateArticleByID(c *gin.Context) {
	idStr := c.Param("id")

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		zap.L().Error("controller.UpdateArticleByID: invalid article id", zap.String("error", "invalid article id"))
		ResponseErrWithMsg(c, InvalidParams, "invalid article id")
		return
	}

	var updateArticleDTO dto.UpdateArticleDTO
	if err := c.ShouldBindJSON(&updateArticleDTO); err != nil {
		// 获取validator.ValidationErrors类型的errors
		errs, ok := err.(validator.ValidationErrors)
		if ok {
			zap.L().Error("controller.UpdateArticleByID: validation params failed", zap.Error(errs))
			ResponseErrWithMsg(c, InvalidParams, errs.Error())
			return
		}

		zap.L().Error("controller.UpdateArticleByID: invalid params", zap.Error(err))
		ResponseErr(c, InvalidParams)
		return
	}

	err = service.UpdateArticleByID(id, updateArticleDTO)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			zap.L().Error("controller.UpdateArticleByID: article not found", zap.Error(gorm.ErrRecordNotFound))
			ResponseErrWithMsg(c, InvalidParams, gorm.ErrRecordNotFound.Error())
			return
		}

		zap.L().Error("controller.UpdateArticleByID: get article error", zap.Error(err))
		ResponseErr(c, ServerError)
		return
	}

	ResponseOk(c, nil)
}

func DeleteArticleByID(c *gin.Context) {
	idStr := c.Param("id")

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		zap.L().Error("controller.DeleteArticleByID: invalid article id", zap.String("error", "invalid article id"))
		ResponseErrWithMsg(c, InvalidParams, "invalid article id")
		return
	}

	err = service.DeleteArticleByID(id)
	if err != nil {
		if errors.Is(err, myerror.ErrorArticleNotFound) {
			zap.L().Error("controller.DeleteArticleByID: article not found", zap.Error(myerror.ErrorArticleNotFound))
			ResponseErrWithMsg(c, InvalidParams, myerror.ErrorArticleNotFound)
			return
		}

		zap.L().Error("controller.DeleteArticleByID: delete article error", zap.Error(err))
		ResponseErr(c, ServerError)
		return
	}

	ResponseOk(c, nil)
}
