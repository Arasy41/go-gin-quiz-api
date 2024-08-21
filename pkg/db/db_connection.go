package db

import (
	"fmt"
	"log"

	"github.com/Arasy41/go-gin-quiz-api/pkg/config"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB(cfg *config.Config) (*gorm.DB, error) {
	var err error
	var dsn string

	switch cfg.DBProvider {
	case "sqlite":
		dsn = cfg.DBName
		DB, err = gorm.Open(sqlite.Open(dsn), &gorm.Config{})
	case "mysql":
		dsn = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName)
		DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	case "postgres":
		dsn = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta",
			cfg.DBHost, cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.DBPort)
		DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	default:
		log.Fatal("Invalid database provider")
		return nil, fmt.Errorf("invalid database provider: %s", cfg.DBProvider)
	}

	if err != nil {
		log.Fatal("Failed to connect to database:", err)
		return nil, err
	}

	log.Println("Database connection established")
	return DB, nil
}
