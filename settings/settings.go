package settings

import (
	"github.com/spf13/viper"
)

type appSettings struct {
	Mode           string `mapstructure:"mode"`
	Addr           string `mapstructure:"addr"`
	JwtSecret      string `mapstructure:"jwt_secret"`
	JwtExpiredHour int    `mapstructure:"jwt_expired_hour"`
}

type mySQLSettings struct {
	Host     string `mapstructure:"host"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	Port     int    `mapstructure:"port"`
	DBName   string `mapstructure:"dbname"`
}

type logSettings struct {
	Filename   string `mapstructure:"filename"`
	MaxSize    int    `mapstructure:"max_size"`
	MaxBackups int    `mapstructure:"max_backups"`
	MaxAge     int    `mapstructure:"max_age"`
}

var AppConfig appSettings
var MySQLConfig mySQLSettings
var LogConfig logSettings

func Init() (err error) {
	viper.SetConfigFile("./configs/dev.toml")

	err = viper.ReadInConfig() // 查找并读取配置文件
	if err != nil {
		return err
	}

	appCfg := viper.Sub("app")
	err = appCfg.Unmarshal(&AppConfig)
	if err != nil {
		return err
	}

	mysqlCfg := viper.Sub("mysql")
	err = mysqlCfg.Unmarshal(&MySQLConfig)
	if err != nil {
		return err
	}

	logCfg := viper.Sub("log")
	err = logCfg.Unmarshal(&LogConfig)
	if err != nil {
		return err
	}

	return err
}
