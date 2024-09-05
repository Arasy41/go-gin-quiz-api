package jwt

import (
	"errors"
	"log"
	"time"

	"github.com/Arasy41/go-gin-quiz-api/config"

	"github.com/golang-jwt/jwt/v5"
)

var secretKey = []byte(config.InitConfig().SecretKey)

type Claims struct {
	UserID   uint   `json:"user_id"`
	UserRole string `json:"user_role"`
	jwt.RegisteredClaims
}

func GenerateToken(userID uint, userRole string) (string, error) {
	log.Println("Generating token with UserRole:", userRole)
	claims := &Claims{
		UserID:   userID,
		UserRole: userRole,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * time.Duration(config.InitConfig().TokenLifespan))),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func ParseToken(tokenString string) (*Claims, error) {
	claims := &Claims{}
	_, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		return nil, err
	}

	if claims.ExpiresAt.Before(time.Now()) {
		return nil, errors.New("token expired")
	}

	return claims, nil
}
