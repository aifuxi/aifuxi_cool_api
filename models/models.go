package models

import (
	"gorm.io/driver/mysql"

	"gorm.io/gorm"
)

var DB *gorm.DB

func Setup() error {
	var err error
	DB, err = gorm.Open(mysql.New(mysql.Config{
		DSN: "root:123456@tcp(localhost:3306)/aifuxi_cool?charset=utf8mb4&parseTime=True&loc=Local",
	}), &gorm.Config{})

	return err
}
