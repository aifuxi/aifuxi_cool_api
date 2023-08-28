package models

import (
	"time"

	"api.aifuxi.cool/internal"
	"gorm.io/gorm"
)

type Article struct {
	ID          int64      `gorm:"column:id;type:bigint;primaryKey" json:"id,string"`
	CreatedAt   time.Time  `gorm:"column:created_at;type:datetime" json:"created_at"`
	UpdatedAt   time.Time  `gorm:"column:updated_at;type:datetime" json:"updated_at"`
	DeletedAt   *time.Time `gorm:"column:deleted_at;type:datetime" json:"deleted_at,omitempty"`
	Title       string     `gorm:"column:title;type:varchar" json:"title"`
	Description string     `gorm:"column:description;type:text" json:"description"`
	Cover       string     `gorm:"column:cover;type:varchar" json:"cover"`
	Content     string     `gorm:"column:content;type:text" json:"content"`
	FriendlyUrl string     `gorm:"column:friendly_url;type:varchar" json:"friendly_url"`
	IsTop       bool       `gorm:"column:is_top;type:bool" json:"is_top"`
	TopPriority int        `gorm:"column:top_priority;type:int" json:"top_priority"`
	IsPublished bool       `gorm:"column:is_published;type:bool" json:"is_published"`
	Tags        []Tag      `gorm:"-" json:"tags"`
}

// GORM 自定义表名
func (Article) TableName() string {
	return "article"
}

func (article *Article) BeforeCreate(tx *gorm.DB) error {
	id, err := internal.GenSnowflakeID()

	if err != nil {
		return err
	}

	article.ID = id

	return nil
}
