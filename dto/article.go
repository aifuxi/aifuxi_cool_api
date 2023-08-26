package dto

type CreateArticleDTO struct {
	Title       string   `json:"title" binding:"required"`
	Description string   `json:"description" binding:"required"`
	Cover       string   `json:"cover"`
	Content     string   `json:"content" binding:"required"`
	FriendlyUrl string   `json:"friendly_url" binding:"required"`
	IsTop       bool     `json:"is_top"`
	TopPriority int      `json:"top_priority"`
	IsPublished bool     `json:"is_published"`
	TagIDs      []string `json:"tag_ids" validate:"dive"`
}

type UpdateArticleDTO struct {
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Cover       string   `json:"cover"`
	Content     string   `json:"content"`
	FriendlyUrl string   `json:"friendly_url"`
	IsTop       bool     `json:"is_top"`
	TopPriority int      `json:"top_priority"`
	IsPublished bool     `json:"is_published"`
	TagIDs      []string `json:"tag_ids" validate:"dive"`
}

type GetArticlesDTO struct {
	Title       string `form:"title"`
	FriendlyUrl string `form:"friendly_url"`

	PaginationDTO
	OrderDTO
}
