package main

import (
	"fmt"
	"log"

	"api.aifuxi.cool/api"
	"api.aifuxi.cool/db/orm"
	"api.aifuxi.cool/settings"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func main() {
	err := settings.Init()
	if err != nil {
		log.Fatalf("初始化配置失败: %v\n", err)
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		settings.MySQLConfig.Username,
		settings.MySQLConfig.Password,
		settings.MySQLConfig.Host,
		settings.MySQLConfig.Port,
		settings.MySQLConfig.DBName,
	)
	var db *gorm.DB
	db, err = gorm.Open(mysql.New(mysql.Config{
		DSN: dsn,
	}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatalf("初始化数据库: %v\n", err)
	}

	store := orm.NewStore(db)
	server := api.NewServer(store)
	address := fmt.Sprintf("localhost:%d", settings.AppConfig.Port)
	server.Start(address)
}
