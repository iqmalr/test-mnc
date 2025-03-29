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
		fmt.Println("Authorization Header:", authHeader)

		if authHeader == "" {
			c.JSON(401, gin.H{"message": "Token tidak ditemukan di header Authorization"})
			c.Abort()
			return
		}
		if !strings.HasPrefix(authHeader, "Bearer ") {
			authHeader = "Bearer " + authHeader
		}
		tokenParts := strings.Split(authHeader, " ")
		if len(tokenParts) != 2 {
			c.JSON(401, gin.H{
				"message":  "Format token salah. Gunakan format 'Bearer <token>'",
				"received": authHeader,
			})
			c.Abort()
			return
		}

		if tokenParts[0] != "Bearer" {
			c.JSON(401, gin.H{
				"message":  "Token harus diawali dengan 'Bearer'",
				"received": tokenParts[0],
			})
			c.Abort()
			return
		}

		token := tokenParts[1]
		fmt.Println("Extracted Token:", token)

		claims, err := utils.VerifyJWT(token)
		if err != nil {
			c.JSON(401, gin.H{
				"message": "Token tidak valid atau sudah expired",
				"error":   err.Error(),
			})
			c.Abort()
			return
		}

		c.Set("user", claims)
		c.Next()
	}
}
