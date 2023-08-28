package orm

type PaginationParams struct {
	Page     int
	PageSize int
}

type OrderParams struct {
	Order   string
	OrderBy string
}
