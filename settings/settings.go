package settings

import (
	"github.com/spf13/viper"
)

type MySQLSettings struct {
	Host     string `mapstructure:"host"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	Port     int    `mapstructure:"port"`
	DBName   string `mapstructure:"dbname"`
}

var MySQLConfig MySQLSettings

func Init() (err error) {
	viper.SetConfigFile("./configs/dev.toml")

	err = viper.ReadInConfig() // 查找并读取配置文件
	if err != nil {
		return err
	}

	mysqlCfg := viper.Sub("mysql")
	err = mysqlCfg.Unmarshal(&MySQLConfig)
	if err != nil {
		return err
	}

	return err
}
