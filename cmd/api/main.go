package main

import (
	"log"

	"github.com/Arasy41/go-gin-quiz-api/config"
	"github.com/Arasy41/go-gin-quiz-api/internal/delivery/router"
	"github.com/Arasy41/go-gin-quiz-api/pkg/db"
	"github.com/Arasy41/go-gin-quiz-api/pkg/helper"
	"github.com/gin-gonic/gin"
)

// @title API Culinary Review
// @version 1.0
// @description This is a sample server for culinary review API.
// @termsOfService http://swagger.io/terms/

// @contact.name Arasy41
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /api

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	config.InitConfig()
	db.InitDB(config.AppConfig)

	environment := helper.Getenv("ENVIRONMENT", "development")
	if environment == "development" {
		gin.SetMode(gin.DebugMode)
	}

	if environment == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	// Initialize router and start the server
	r := router.InitRouter()

	if err := r.Run(":" + config.AppConfig.Port); err != nil {
		log.Fatal("Failed to run server:", err)
	}
}
