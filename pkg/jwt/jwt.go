package jwt

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Arasy41/go-gin-quiz-api/config"
	"github.com/Arasy41/go-gin-quiz-api/pkg/helper"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

var API_SECRET = config.AppConfig.SecretKey

func GenerateToken(user_id, role_id uint) (string, error) {
	token_lifespan, err := strconv.Atoi(helper.Getenv("TOKEN_HOUR_LIFESPAN", "1"))

	if err != nil {
		return "", err
	}

	claims := jwt.MapClaims{
		"authorized": true,
		"user_id":    user_id,
		"role":       role_id,
		"exp":        time.Now().Add(time.Hour * time.Duration(token_lifespan)).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(API_SECRET))

}

func TokenValid(c *gin.Context) error {
	tokenString := ExtractToken(c)
	_, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(API_SECRET), nil
	})
	if err != nil {
		return err
	}
	return nil
}

func ExtractToken(c *gin.Context) string {
	token := c.Query("token")

	if token != "" {
		return token
	}

	bearerToken := c.Request.Header.Get("Authorization")
	if len(strings.Split(bearerToken, " ")) == 2 {
		return strings.Split(bearerToken, " ")[1]
	}

	return ""
}

func ExtractTokenID(c *gin.Context) (uint, error) {

	tokenString := ExtractToken(c)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(API_SECRET), nil
	})
	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		uid, err := strconv.ParseUint(fmt.Sprintf("%.0f", claims["user_id"]), 10, 32)
		if err != nil {
			return 0, err
		}
		return uint(uid), nil
	}

	return 0, nil
}
