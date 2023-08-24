package dto

type CreateUserDTO struct {
	Nickname   string `json:"nickname" binding:"required"`
	Avatar     string `json:"avatar"`
	Email      string `json:"email" binding:"required"`
	Password   string `json:"password" binding:"required"`
	RePassword string `json:"re_password" binding:"required,eqfield=Password"`
}

type UpdateUserDTO struct {
	Nickname string `json:"nickname"`
	Avatar   string `json:"avatar"`
	Password string `json:"password"`
}

type GetUsersDTO struct {
	Nickname string `form:"nickname"`
	Email    string `form:"email"`

	PaginationDTO
	OrderDTO
}
