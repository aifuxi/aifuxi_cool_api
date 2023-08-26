package models

import (
	"time"

	"api.aifuxi.cool/internal"
	"gorm.io/gorm"
)

type User struct {
	ID        int64      `gorm:"column:id;type:bigint;primaryKey" json:"id,string"`
	Nickname  string     `gorm:"column:nickname;type:varchar" json:"nickname"`
	Avatar    string     `gorm:"column:avatar;type:varchar" json:"avatar"`
	Email     string     `gorm:"column:email;type:varchar" json:"email"`
	Password  string     `gorm:"column:password;type:varchar" json:"-"`
	CreatedAt time.Time  `gorm:"column:created_at;type:datetime" json:"created_at"`
	UpdatedAt time.Time  `gorm:"column:updated_at;type:datetime" json:"updated_at"`
	DeletedAt *time.Time `gorm:"column:deleted_at;type:datetime" json:"deleted_at,omitempty"`
}

// GORM 自定义表名
func (User) TableName() string {
	return "user"
}

func (user *User) BeforeCreate(tx *gorm.DB) error {
	id, err := internal.GenSnowflakeID()

	if err != nil {
		return err
	}

	user.ID = id

	return nil
}
