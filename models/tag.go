package models

import (
	"time"

	"api.aifuxi.cool/internal"
	"gorm.io/gorm"
)

type Tag struct {
	ID          int64      `gorm:"column:id;type:bigint;primaryKey" json:"id,string"`
	CreatedAt   time.Time  `gorm:"column:created_at;type:datetime" json:"created_at"`
	UpdatedAt   time.Time  `gorm:"column:updated_at;type:datetime" json:"updated_at"`
	DeletedAt   *time.Time `gorm:"column:deleted_at;type:datetime" json:"deleted_at,omitempty"`
	Name        string     `gorm:"column:name;type:varchar" json:"name"`
	FriendlyUrl string     `gorm:"column:friendly_url;type:varchar" json:"friendly_url"`
}

// GORM 自定义表名
func (Tag) TableName() string {
	return "tag"
}

func (tag *Tag) BeforeCreate(tx *gorm.DB) error {
	id, err := internal.GenSnowflakeID()

	if err != nil {
		return err
	}

	tag.ID = id

	return nil
}
