package main

import (
	"log"

	"github.com/Arasy41/go-gin-quiz-api/config"
	"github.com/Arasy41/go-gin-quiz-api/internal/delivery/router"
	"github.com/Arasy41/go-gin-quiz-api/pkg/db"
	"github.com/Arasy41/go-gin-quiz-api/pkg/logger"
	"github.com/gin-gonic/gin"
)

func main() {
	// Inisialisasi logger dengan file log di dalam folder logs
	if err := logger.InitLogger("logs/app.log"); err != nil {
		log.Fatalf("Failed to initialize logger: %v", err)
	}
	defer logger.CloseLogger()

	// Setup close handler untuk menangani signal SIGINT/SIGTERM
	logger.SetupCloseHandler()

	// Inisialisasi konfigurasi dan database
	cfg := config.InitConfig()
	db.InitDB(cfg)

	// Initialize Environment
	environment := cfg.Environment

	// Initialize router and start the server
	r := router.InitRouter(db.DB)
	if environment == "production" {
		gin.SetMode(gin.ReleaseMode)
		if err := r.RunTLS(":"+cfg.Port, "cert.pem", "key.pem"); err != nil {
			log.Fatal("Failed to run server with TLS:", err)
		}
	} else if environment == "development" {
		gin.SetMode(gin.DebugMode)
		if err := r.Run(":" + cfg.Port); err != nil {
			log.Fatal("Failed to run server:", err)
		}
	}
}
