package dto

type PaginationDTO struct {
	Page     int `form:"page" binding:"required"`
	PageSize int `form:"page_size" binding:"required"`
}

// TODO: 完成排序
type OrderDTO struct {
	Order   string `form:"order" binding:"required"`
	OrderBy string `form:"order_by" binding:"required"`
}
