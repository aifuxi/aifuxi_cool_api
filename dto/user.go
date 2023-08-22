package dto

type CreateUserDTO struct {
	Nickname   string `json:"nickname" binding:"required"`
	Email      string `json:"email" binding:"required"`
	Password   string `json:"password" binding:"required"`
	RePassword string `json:"re_password" binding:"required,eqfield=Password"`
}

type UpdateUserDTO struct {
	Nickname string `json:"nickname"`
	Password string `json:"password"`
}

type GetUsersDTO struct {
	Nickname string `form:"nickname"`
	Email    string `form:"email"`

	PaginationDTO
	OrderDTO
}
