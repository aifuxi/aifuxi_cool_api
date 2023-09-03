package db

import (
	"api.aifuxi.cool/util"
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID        int64     `gorm:"column:id;type:bigint;primaryKey" json:"id,string"`
	Nickname  string    `gorm:"column:nickname;type:varchar" json:"nickname"`
	Avatar    string    `gorm:"column:avatar;type:varchar" json:"avatar"`
	Email     string    `gorm:"column:email;type:varchar" json:"email"`
	Password  string    `gorm:"column:password;type:varchar" json:"-"`
	CreatedAt time.Time `gorm:"column:created_at;type:datetime" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at;type:datetime" json:"updated_at"`
}

func (u *User) TableName() string {
	return "user"
}

func (u *User) BeforeCreate(tx *gorm.DB) error {

	id, err := util.NewSnowflakeID()
	if err != nil {
		return err
	}

	password, err := util.HashPassword(u.Password)
	if err != nil {
		return err
	}

	u.ID = id
	u.Password = password

	return nil
}

type Tag struct {
	ID          int64     `gorm:"column:id;type:bigint;primaryKey" json:"id,string"`
	Name        string    `gorm:"column:name;type:varchar" json:"name"`
	FriendlyURL string    `gorm:"column:friendly_url;type:varchar" json:"friendly_url"`
	CreatedAt   time.Time `gorm:"column:created_at;type:datetime" json:"created_at"`
	UpdatedAt   time.Time `gorm:"column:updated_at;type:datetime" json:"updated_at"`

	// 非数据字段，代表标签上有多少文章
	ArticleCount int `gorm:"-" json:"article_count"`
}

func (t *Tag) TableName() string {
	return "tag"
}

func (t *Tag) BeforeCreate(tx *gorm.DB) error {

	id, err := util.NewSnowflakeID()
	if err != nil {
		return err
	}

	t.ID = id

	return nil
}

type Article struct {
	ID          int64     `gorm:"column:id;type:bigint;primaryKey" json:"id,string"`
	CreatedAt   time.Time `gorm:"column:created_at;type:datetime" json:"created_at"`
	UpdatedAt   time.Time `gorm:"column:updated_at;type:datetime" json:"updated_at"`
	Title       string    `gorm:"column:title;type:varchar" json:"title"`
	Description string    `gorm:"column:description;type:text" json:"description"`
	Cover       string    `gorm:"column:cover;type:varchar" json:"cover"`
	Content     string    `gorm:"column:content;type:text" json:"content"`
	FriendlyURL string    `gorm:"column:friendly_url;type:varchar" json:"friendly_url"`
	IsTop       bool      `gorm:"column:is_top;type:bool" json:"is_top"`
	TopPriority int       `gorm:"column:top_priority;type:int" json:"top_priority"`
	IsPublished bool      `gorm:"column:is_published;type:bool" json:"is_published"`

	// 非数据字段，表示文章下面的所有标签
	Tags []Tag `gorm:"-" json:"tags"`
}

func (a *Article) TableName() string {
	return "article"
}

func (a *Article) BeforeCreate(tx *gorm.DB) error {
	id, err := util.NewSnowflakeID()

	if err != nil {
		return err
	}

	a.ID = id

	return nil
}

type ArticleTag struct {
	ArticleID int64 `gorm:"column:article_id;type:bigint;" json:"article_id,string"`
	TagID     int64 `gorm:"column:tag_id;type:bigint;" json:"tag_id,string"`
}

func (at *ArticleTag) TableName() string {
	return "article_tag"
}
