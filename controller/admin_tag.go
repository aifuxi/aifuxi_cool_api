package controller

import (
	"errors"

	"github.com/aifuxi/aifuxi_cool_api/dao/mysql"
	"github.com/aifuxi/aifuxi_cool_api/dto"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func GetTags(c *gin.Context) {
	tags, err := mysql.GetTags()
	if err != nil {
		zap.L().Error("获取文章标签列表失败", zap.Error(err))
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
			zap.L().Error("参数校验失败", zap.Error(errs))
			ResponseErrWithMsg(c, InvalidParams, errs.Error())
			return
		}
		zap.L().Error("创建文章标签失败", zap.Error(err))
		ResponseErr(c, InvalidParams)
		return
	}

	tag, err := mysql.CreateTag(createTagDTO)
	if err != nil {
		zap.L().Error("创建文章标签失败", zap.Error(err))
		ResponseErr(c, ServerError)
		return
	}

	ResponseOk(c, tag)
}

func GetTagByID(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		zap.L().Error("获取文章标签失败", zap.String("error", "缺少id"))
		ResponseErrWithMsg(c, InvalidParams, "缺少id")
		return
	}

	tag, err := mysql.GetTagByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			zap.L().Error("文章标签不存在", zap.Error(gorm.ErrRecordNotFound))
			ResponseErrWithMsg(c, InvalidParams, gorm.ErrRecordNotFound.Error())
			return
		}

		zap.L().Error("获取文章标签失败", zap.Error(err))
		ResponseErr(c, ServerError)
		return
	}

	ResponseOk(c, tag)
}

func DeleteTagByID(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		zap.L().Error("删除文章标签失败", zap.String("error", "缺少id"))
		ResponseErrWithMsg(c, InvalidParams, "缺少id")
		return
	}

	err := mysql.DeleteTagByID(id)
	if err != nil {
		if errors.Is(err, mysql.ErrorTagNotFound) {
			zap.L().Error("文章标签不存在", zap.Error(mysql.ErrorTagNotFound))
			ResponseErrWithMsg(c, InvalidParams, mysql.ErrorTagNotFound)
			return
		}

		zap.L().Error("删除文章标签失败", zap.Error(err))
		ResponseErr(c, ServerError)
		return
	}

	ResponseOk(c, nil)
}
