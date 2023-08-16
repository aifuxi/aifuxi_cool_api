package models

import (
	"fmt"

	"github.com/aifuxi/aifuxi_cool_api/settings"
	"gorm.io/driver/mysql"

	"gorm.io/gorm"
)

var DB *gorm.DB

func Setup() error {
	var err error

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		settings.MySQLConfig.Username,
		settings.MySQLConfig.Password,
		settings.MySQLConfig.Host,
		settings.MySQLConfig.Port,
		settings.MySQLConfig.DBName,
	)

	DB, err = gorm.Open(mysql.New(mysql.Config{
		DSN: dsn,
	}), &gorm.Config{})

	return err
}
