package config

import (
	"fmt"
	"strings"
	"test-mnc/utils"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		fmt.Println("Authorization Header:", authHeader) // Debug log

		if authHeader == "" {
			c.JSON(401, gin.H{"message": "Token tidak ditemukan"})
			c.Abort()
			return
		}

		tokenParts := strings.Split(authHeader, " ")
		if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
			fmt.Println("Invalid Token Format") // Debug log
			c.JSON(401, gin.H{"message": "Format token tidak valid"})
			c.Abort()
			return
		}

		token := tokenParts[1]
		fmt.Println("Extracted Token:", token) // Debug log

		claims, err := utils.VerifyJWT(token)
		if err != nil {
			fmt.Println("Token Verification Error:", err) // Debug log
			c.JSON(401, gin.H{"message": "Token tidak valid atau telah kadaluarsa"})
			c.Abort()
			return
		}

		c.Set("user", claims)
		c.Next()
	}
}
