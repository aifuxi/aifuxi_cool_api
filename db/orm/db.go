package orm

import "gorm.io/gorm"

type Queries struct {
	db *gorm.DB
}

func New(db *gorm.DB) *Queries {
	return &Queries{db: db}
}
