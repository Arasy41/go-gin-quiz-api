package jwt

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var jwtKey = []byte(os.Getenv("SECRET_KEY"))

type Claims struct {
	ID       uint
	Username string
	RoleID   uint
	jwt.StandardClaims
}

func GenerateJWT(id uint, username string, roleID uint) (string, error) {
	expirationTime := time.Now().Add(1 * time.Hour)

	claims := &Claims{
		ID:       id,
		Username: username,
		RoleID:   roleID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(jwtKey)
}

func ValidateToken(signedToken string) (*jwt.Token, error) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&Claims{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(jwtKey), nil
		},
	)

	if err != nil {
		return nil, err
	}

	return token, nil
}
