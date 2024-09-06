package middleware

import (
	"net/http"
	"strings"

	"github.com/Arasy41/go-gin-quiz-api/internal/domain/models"
	"github.com/Arasy41/go-gin-quiz-api/pkg/constant"
	"github.com/Arasy41/go-gin-quiz-api/pkg/jwt"
	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
)

func JWTAuthMiddleware(db *gorm.DB, allowedRoles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader(constant.AuthorizationKey)
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header required"})
			c.Abort()
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		claims, err := jwt.ParseToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		// Fetch user from database to get the role
		var user models.User
		if err := db.Preload("Role").Where("id = ?", claims.UserID).First(&user).Error; err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
			c.Abort()
			return
		}

		// Set user data in context BEFORE calling c.Next()
		c.Set("user_id", user.ID)
		c.Set("user_role", user.Role.Name)

		// Check if the user has one of the allowed roles
		for _, role := range allowedRoles {
			if user.Role.Name == role || user.Role.Name == "admin" {
				c.Next()
				return
			}
		}

		c.JSON(http.StatusForbidden, gin.H{"error": "You don't have permission to access this resource"})
		c.Abort()
	}
}
