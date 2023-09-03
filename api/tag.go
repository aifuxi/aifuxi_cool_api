package api

import (
	db "api.aifuxi.cool/db/orm"
	"errors"
	"github.com/gin-gonic/gin"
	"strconv"
)

type listTagsRequest struct {
	Name        string `form:"name"`
	FriendlyURL string `form:"friendly_url"`

	Page     int    `form:"page" binding:"required,gte=1"`
	PageSize int    `form:"page_size" binding:"required,gte=10,lte=500"`
	Order    string `form:"order" binding:"required,oneof=asc desc"`
	OrderBy  string `form:"order_by" binding:"required,oneof=created_at updated_at"`
}

func (s *Server) ListTags(c *gin.Context) {
	var req listTagsRequest

	if err := c.ShouldBindQuery(&req); err != nil {
		responseFailWithErr(c, ResponseCodeBadRequest, err)
		return
	}

	arg := db.ListTagsParams{
		Name:        req.Name,
		FriendlyURL: req.FriendlyURL,
		Page:        req.Page,
		PageSize:    req.PageSize,
		Order:       req.Order,
		OrderBy:     req.OrderBy,
	}

	tags, total, err := s.store.ListTags(arg)
	if err != nil {
		responseFailWithErr(c, ResponseCodeFail, err)
		return
	}

	responseOkWithTotal(c, tags, total)
}

type createTagRequest struct {
	Name        string `json:"name" binding:"required"`
	FriendlyURL string `json:"friendly_url" binding:"required"`
}

func (s *Server) CreateTag(c *gin.Context) {
	var req createTagRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		responseFailWithErr(c, ResponseCodeBadRequest, err)
		return
	}

	arg := db.CreateTagParams{
		Name:        req.Name,
		FriendlyURL: req.FriendlyURL,
	}

	tag, err := s.store.CreateTag(arg)
	if err != nil {
		if errors.Is(err, db.ErrTagExist) {
			responseFailWithErr(c, ResponseCodeBadRequest, err)
			return
		}

		responseFailWithErr(c, ResponseCodeFail, err)
		return
	}

	responseOk(c, tag)
}

func (s *Server) GetTag(c *gin.Context) {

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		responseFailWithErr(c, ResponseCodeBadRequest, err)
		return
	}

	var tag db.Tag
	tag, err = s.store.GetTagByID(id)
	if err != nil {
		if errors.Is(err, db.ErrTagNotFound) {
			responseFailWithErr(c, ResponseCodeBadRequest, err)
			return
		}
		responseFailWithErr(c, ResponseCodeFail, err)
		return
	}

	responseOk(c, tag)
}

type updateTagRequest struct {
	Name        string `json:"name"`
	FriendlyURL string `json:"friendly_url"`
}

func (s *Server) UpdateTag(c *gin.Context) {
	var req updateTagRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		responseFailWithErr(c, ResponseCodeBadRequest, err)
		return
	}

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		responseFailWithErr(c, ResponseCodeBadRequest, err)
		return
	}

	arg := db.UpdateTagParams{
		Name:        req.Name,
		FriendlyURL: req.FriendlyURL,
	}

	err = s.store.UpdateTag(id, arg)
	if err != nil {
		if errors.Is(err, db.ErrTagNotFound) {
			responseFailWithErr(c, ResponseCodeBadRequest, err)
			return
		}
		responseFailWithErr(c, ResponseCodeFail, err)
		return
	}

	responseOk(c, nil)
}

func (s *Server) DeleteTag(c *gin.Context) {

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		responseFailWithErr(c, ResponseCodeBadRequest, err)
		return
	}

	err = s.store.DeleteTagByID(id)
	if err != nil {
		if errors.Is(err, db.ErrTagNotFound) {
			responseFailWithErr(c, ResponseCodeBadRequest, err)
			return
		}
		responseFailWithErr(c, ResponseCodeFail, err)
		return
	}

	responseOk(c, nil)
}
