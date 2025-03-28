package controllers

import (
	"encoding/json"
	"net/http"
	"os"
	"test-mnc/models"

	"github.com/gin-gonic/gin"
)

func GetAllUsers(c *gin.Context) {
	file, err := os.ReadFile("data/user.json")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error get JSON"})
		return
	}

	var users []models.User
	err = json.Unmarshal(file, &users)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Eror get JSON"})
		return
	}

	c.JSON(http.StatusOK, users)
}
