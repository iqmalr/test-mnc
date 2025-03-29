package controllers

import (
	"encoding/json"
	"net/http"
	"os"
	"test-mnc/models"

	"github.com/gin-gonic/gin"
)

// GetAllUsers godoc
// @Summary Get all users
// @Description Menampilkan daftar semua pengguna
// @Tags users
// @Produce json
// @Success 200 {array} models.User
// @Failure 500 {object} map[string]interface{}
// @Router /users [get]
func GetAllUsers(c *gin.Context) {
	file, err := os.ReadFile("data/user.json")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error get data"})
		return
	}

	var users []models.User
	err = json.Unmarshal(file, &users)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "eror umarshal data"})
		return
	}

	c.JSON(http.StatusOK, users)
}
