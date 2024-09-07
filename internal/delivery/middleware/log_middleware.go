package middleware

import (
	"log"
	"time"

	"github.com/Arasy41/go-gin-quiz-api/pkg/logger"
	"github.com/gin-gonic/gin"
)

func RequestLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Rotate log file
		if err := logger.RotateLogFileIfNeeded("logs"); err != nil {
			log.Printf("Failed to rotate log file: %v", err)
		}
		// Catat waktu sebelum request
		startTime := time.Now()

		// Proses request
		c.Next()

		// Catat waktu setelah request selesai
		endTime := time.Now()
		latency := endTime.Sub(startTime)

		// Ambil status code dan informasi lainnya
		statusCode := c.Writer.Status()
		clientIP := c.ClientIP()
		method := c.Request.Method
		path := c.Request.URL.Path

		// Log informasi request
		log.Printf("| %3d | %13v | %15s | %-7s  %#v | message : %s\n",
			statusCode,
			latency,
			clientIP,
			method,
			path,
			c.Errors.ByType(gin.ErrorTypePrivate).String(),
		)

		// Jika status code menunjukkan success, log sebagai info
		if statusCode < 400 {
			log.Printf("INFO: %d - %s %s", statusCode, method, path)
		}

		// Jika status code menunjukkan error, log sebagai error
		if statusCode >= 400 {
			log.Printf("ERROR: %d - %s %s", statusCode, method, path)
		}
	}
}
