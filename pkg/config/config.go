package config

import (
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	Host        string
	Environment string
	GinMode     string
	Port        string

	DBProvider string
	DBHost     string
	DBUser     string
	DBPassword string
	DBName     string
	DBPort     string

	SecretKey string
}

var AppConfig *Config

func InitConfig() {
	viper.SetConfigFile(".env")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	AppConfig = &Config{
		Host:        viper.GetString("HOST"),
		Environment: viper.GetString("ENVIRONMENT"),
		GinMode:     viper.GetString("GIN_MODE"),
		Port:        viper.GetString("PORT"),

		DBProvider: viper.GetString("DB_PROVIDER"),
		DBHost:     viper.GetString("DB_HOST"),
		DBUser:     viper.GetString("DB_USER"),
		DBPassword: viper.GetString("DB_PASSWORD"),
		DBName:     viper.GetString("DB_NAME"),
		DBPort:     viper.GetString("DB_PORT"),

		SecretKey: viper.GetString("SECRET_KEY"),
	}
}
