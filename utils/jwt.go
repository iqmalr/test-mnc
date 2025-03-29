package utils

import (
	"time"

	"test-mnc/models"

	"github.com/golang-jwt/jwt/v5"
)

var jwtKey = []byte("rahasia")

func GenerateJWT(user *models.AuthUser) (string, error) {
	claims := jwt.MapClaims{
		"id":    user.ID,
		"name":  user.Name,
		"email": user.Email,
		"exp":   time.Now().Add(5 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}
