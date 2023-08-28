package mysql

type PaginationParams struct {
	Page     int `form:"page" binding:"required"`
	PageSize int `form:"page_size" binding:"required"`
}

type OrderParams struct {
	Order   string `form:"order" binding:"required"`
	OrderBy string `form:"order_by" binding:"required"`
}
