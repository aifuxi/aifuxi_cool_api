package api

import (
	db "api.aifuxi.cool/db/orm"
	"errors"
	"github.com/gin-gonic/gin"
	"strconv"
)

type listArticlesRequest struct {
	Title       string `form:"title"`
	FriendlyURL string `form:"friendly_url"`

	Page     int    `form:"page" binding:"required,gte=1"`
	PageSize int    `form:"page_size" binding:"required,gte=10,lte=50"`
	Order    string `form:"order" binding:"required,oneof=asc desc"`
	OrderBy  string `form:"order_by" binding:"required,oneof=created_at updated_at"`
}

func (s *Server) ListArticles(c *gin.Context) {
	var req listArticlesRequest

	if err := c.ShouldBindQuery(&req); err != nil {
		responseFailWithErr(c, ResponseCodeBadRequest, err)
		return
	}

	arg := db.ListArticlesParams{
		Title:       req.Title,
		FriendlyURL: req.FriendlyURL,
		Page:        req.Page,
		PageSize:    req.PageSize,
		Order:       req.Order,
		OrderBy:     req.OrderBy,
	}

	articles, total, err := s.store.ListArticles(arg)
	if err != nil {
		responseFailWithErr(c, ResponseCodeFail, err)
		return
	}

	responseOkWithTotal(c, articles, total)
}

type createArticleRequest struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
	Cover       string `json:"cover"`
	Content     string `json:"content" binding:"required"`
	FriendlyURL string `json:"friendly_url" binding:"required"`
	IsTop       bool   `json:"is_top"`
	TopPriority int    `json:"top_priority"`
	IsPublished bool   `json:"is_published"`
}

func (s *Server) CreateArticle(c *gin.Context) {
	var req createArticleRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		responseFailWithErr(c, ResponseCodeBadRequest, err)
		return
	}

	arg := db.CreateArticleParams{
		Title:       req.Title,
		Description: req.Description,
		Cover:       req.Cover,
		Content:     req.Content,
		FriendlyURL: req.FriendlyURL,
		IsTop:       req.IsTop,
		TopPriority: req.TopPriority,
		IsPublished: req.IsPublished,
	}

	user, err := s.store.CreateArticle(arg)
	if err != nil {
		if errors.Is(err, db.ErrArticleExist) {
			responseFailWithErr(c, ResponseCodeBadRequest, err)
			return
		}

		responseFailWithErr(c, ResponseCodeFail, err)
		return
	}

	responseOk(c, user)
}

func (s *Server) GetArticle(c *gin.Context) {

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		responseFailWithErr(c, ResponseCodeBadRequest, err)
		return
	}

	var user db.Article
	user, err = s.store.GetArticleByID(id)
	if err != nil {
		if errors.Is(err, db.ErrArticleNotFound) {
			responseFailWithErr(c, ResponseCodeBadRequest, err)
			return
		}
		responseFailWithErr(c, ResponseCodeFail, err)
		return
	}

	responseOk(c, user)
}

type updateArticleRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Cover       string `json:"cover"`
	Content     string `json:"content"`
	FriendlyURL string `json:"friendly_url"`
	IsTop       bool   `json:"is_top"`
	TopPriority int    `json:"top_priority"`
	IsPublished bool   `json:"is_published"`
}

func (s *Server) UpdateArticle(c *gin.Context) {
	var req updateArticleRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		responseFailWithErr(c, ResponseCodeBadRequest, err)
		return
	}

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		responseFailWithErr(c, ResponseCodeBadRequest, err)
		return
	}

	arg := db.UpdateArticleParams{
		Title:       req.Title,
		Description: req.Description,
		Cover:       req.Cover,
		Content:     req.Content,
		FriendlyURL: req.FriendlyURL,
		IsTop:       req.IsTop,
		TopPriority: req.TopPriority,
		IsPublished: req.IsPublished,
	}

	err = s.store.UpdateArticle(id, arg)
	if err != nil {
		if errors.Is(err, db.ErrArticleNotFound) {
			responseFailWithErr(c, ResponseCodeBadRequest, err)
			return
		}
		responseFailWithErr(c, ResponseCodeFail, err)
		return
	}

	responseOk(c, nil)
}

func (s *Server) DeleteArticle(c *gin.Context) {

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		responseFailWithErr(c, ResponseCodeBadRequest, err)
		return
	}

	err = s.store.DeleteArticleByID(id)
	if err != nil {
		if errors.Is(err, db.ErrArticleNotFound) {
			responseFailWithErr(c, ResponseCodeBadRequest, err)
			return
		}
		responseFailWithErr(c, ResponseCodeFail, err)
		return
	}

	responseOk(c, nil)
}
