package middleware

import (
	"net/http"

	"github.com/Arasy41/go-gin-quiz-api/internal/domain/models"
	"github.com/Arasy41/go-gin-quiz-api/pkg/jwt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func JwtAuthMiddleware(allowedRoles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		err := jwt.TokenValid(c)
		if err != nil {
			c.String(http.StatusUnauthorized, err.Error())
			c.Abort()
			return
		}

		userId, err := jwt.ExtractTokenID(c)
		if err != nil {
			c.String(http.StatusUnauthorized, err.Error())
			c.Abort()
			return
		}

		var user models.User
		db := c.MustGet("db").(*gorm.DB)
		findUserErr := db.Preload("Role").Where("id = ?", userId).First(&user).Error
		if findUserErr != nil {
			c.String(http.StatusUnauthorized, findUserErr.Error())
			c.Abort()
			return
		}

		// Check if the user's role is allowed to access the route
		for _, role := range allowedRoles {
			if role == user.Role.Name || user.Role.Name == "admin" {
				c.Set("user_id", user.ID)
				c.Set("user_role", user.Role.Name)
				c.Next()
				return
			}
		}

		c.JSON(http.StatusForbidden, gin.H{"error": "Forbidden access"})
		c.Abort()
	}
}
