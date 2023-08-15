package models

import "time"

type Tag struct {
	ID          int64      `gorm:"column:id;type:bigint;primaryKey" json:"id,string"`
	Name        string     `gorm:"column:name;type:varchar" json:"name"`
	FriendlyUrl string     `gorm:"column:friendly_url;type:varchar" json:"friendly_url"`
	CreatedAt   time.Time  `gorm:"column:created_at;type:datetime" json:"created_at"`
	UpdatedAt   time.Time  `gorm:"column:updated_at;type:datetime" json:"updated_at"`
	DeletedAt   *time.Time `gorm:"column:deleted_at;type:datetime" json:"deleted_at,omitempty"`
}

// GORM 自定义表名
func (Tag) TableName() string {
	return "tag"
}
