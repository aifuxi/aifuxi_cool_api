package models

import "time"

type Model struct {
	ID        int64      `gorm:"column:id;type:bigint;primaryKey" json:"id,string"`
	CreatedAt time.Time  `gorm:"column:created_at;type:datetime" json:"created_at"`
	UpdatedAt time.Time  `gorm:"column:updated_at;type:datetime" json:"updated_at"`
	DeletedAt *time.Time `gorm:"column:deleted_at;type:datetime" json:"deleted_at,omitempty"`
}
