package dto

type CreateTagDTO struct {
	Name        string `json:"name" binding:"required"`
	FriendlyUrl string `json:"friendly_url" binding:"required"`
}

type UpdateTagDTO struct {
	Name        string `json:"name"`
	FriendlyUrl string `json:"friendly_url"`
}

type QueryTagDTO struct {
	PaginationDTO
	OrderDTO
}
