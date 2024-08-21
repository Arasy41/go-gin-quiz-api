package db

import (
	"log"

	"github.com/Arasy41/go-gin-quiz-api/internal/domain/models"
	"github.com/Arasy41/go-gin-quiz-api/pkg/config"
)

// InitDB initializes the database connection using the configuration provided.
func InitDB(cfg *config.Config) {
	var err error
	DB, err = ConnectDB(cfg)
	if err != nil {
		log.Fatal("Could not initialize the database connection:", err)
	}

	// Add AutoMigrate or other database initialization steps here if needed
	DB.AutoMigrate(
		&models.User{},
		&models.Quiz{},
		&models.Question{},
		&models.Role{},
	)

	log.Println("Database initialized successfully")
}
