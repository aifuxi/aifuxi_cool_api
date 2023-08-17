package dto

type CreateTagDTO struct {
	Name        string `json:"name" binding:"required"`
	FriendlyUrl string `json:"friendly_url" binding:"required"`
}
