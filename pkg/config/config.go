package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	Host        string
	Environment string
	Port        string

	DBProvider string
	DBHost     string
	DBUser     string
	DBPassword string
	DBName     string
	DBPort     string

	SecretKey     string
	TokenLifespan int
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
		Port:        viper.GetString("PORT"),

		DBProvider: viper.GetString("DB_PROVIDER"),
		DBHost:     viper.GetString("DB_HOST"),
		DBUser:     viper.GetString("DB_USER"),
		DBPassword: viper.GetString("DB_PASSWORD"),
		DBName:     viper.GetString("DB_NAME"),
		DBPort:     viper.GetString("DB_PORT"),

		SecretKey:     viper.GetString("SECRET_KEY"),
		TokenLifespan: viper.GetInt("TOKEN_HOUR_LIFESPAN"),
	}
}
