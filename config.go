package config

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
)

var Config _Config

func InitConfig() {
	viper.SetConfigName("dev")
	viper.SetConfigType("toml")
	viper.AddConfigPath("./conf/")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("read config failed: %v", err)
	}

	viper.Unmarshal(&Config)

	fmt.Println(Config.MySQL)
}

type _Config struct {
	AppName  string `mapstructure:"app_name"`
	LogLevel string `mapstructure:"log_Level"`

	MySQL _MySQLConfig
	Redis _RedisConfig
}

type _MySQLConfig struct {
	IP       string
	Port     int
	User     string
	Password string
	Database string
}

type _RedisConfig struct {
	IP   string
	Port int
}

