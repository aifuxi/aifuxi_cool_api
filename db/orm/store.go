package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Store interface {
	Querier
}

type SQLStore struct {
	*Queries
}

func NewStore() (Store, error) {
	dsn := "root:123456@tcp(127.0.0.1:3306)/my_website?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	store := &SQLStore{
		Queries: NewQueries(db),
	}

	return store, nil
}
