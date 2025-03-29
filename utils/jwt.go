package utils

import (
	"errors"
	"fmt"
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
func VerifyJWT(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			fmt.Println("Invalid Signing Method")
			return nil, errors.New("metode tanda tangan tidak valid")
		}
		return jwtKey, nil
	})

	if err != nil {
		fmt.Println("Token Parsing Error:", err)
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		exp, ok := claims["exp"].(float64)
		if !ok {
			fmt.Println("Token Exp Field Missing")
			return nil, errors.New("token tidak memiliki exp")
		}

		if time.Now().Unix() > int64(exp) {
			fmt.Println("Token Expired")
			return nil, errors.New("token telah kadaluarsa")
		}

		fmt.Println("Token Verified:", claims)
		return claims, nil
	}

	fmt.Println("Invalid Token Claims")
	return nil, errors.New("token tidak valid")
}
