package mysql

import "gorm.io/gorm"

type Store interface {
	Querier
}

type SQLStore struct {
	*Queries
}

func NewStore(db *gorm.DB) Store {
	return &SQLStore{
		Queries: New(db),
	}
}
