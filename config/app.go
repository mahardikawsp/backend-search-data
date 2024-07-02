package config

import (
	"fmt"
	"github.com/spf13/viper"
)

type Config struct{}

type AppInfo struct {
	Name string
	Port string
}

type ElasticSearchInfo struct {
	Host     string
	Username string
	Password string
	UseSSL   string
}

type MysqlInfo struct {
	Host      string
	Port      string
	Username  string
	Password  string
	DBName    string
	IsMigrate string
}

func New() Config {
	return Config{}
}

func (c Config) SetupConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath("./config")
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Read Config from :", viper.ConfigFileUsed())
	} else {
		fmt.Println("Redirected using automatic env")
		viper.AutomaticEnv()
	}
}

func GetAppName() AppInfo {
	return AppInfo{
		Name: viper.GetString("APP_NAME"),
		Port: viper.GetString("APP_PORT"),
	}
}

func GetElasticSearchInfo() ElasticSearchInfo {
	return ElasticSearchInfo{
		Host:     viper.GetString("ES_HOST"),
		Username: viper.GetString("ES_USERNAME"),
		Password: viper.GetString("ES_PASSWORD"),
		UseSSL:   viper.GetString("USE_SSL"),
	}
}

func GetMysqlInfo() MysqlInfo {
	return MysqlInfo{
		Host:      viper.GetString("DB_HOST"),
		Port:      viper.GetString("DB_PORT"),
		Username:  viper.GetString("DB_USERNAME"),
		Password:  viper.GetString("DB_PASSWORD"),
		DBName:    viper.GetString("DB_NAME"),
		IsMigrate: viper.GetString("IS_MIGRATE"),
	}
}
