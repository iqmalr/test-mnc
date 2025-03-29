package controllers

import (
	"encoding/json"
	"net/http"
	"os"
	"test-mnc/models"
	"test-mnc/utils"

	"github.com/gin-gonic/gin"
)

// GetAllUsers godoc
// @Summary Get all users
// @Description Menampilkan daftar semua pengguna
// @Tags users
// @Security BearerAuth
// @Produce json
// @Success 200 {array} models.User
// @Failure 401 {object} utils.ErrorResponse "Unauthorized"
// @Failure 500 {object} utils.ErrorResponse "Internal Server Error"
// @Router /users [get]
func GetAllUsers(c *gin.Context) {
	file, err := os.ReadFile("data/user.json")
	if err != nil {
		utils.HandleDatabaseError(c, err)
		return
	}

	var users []models.User
	err = json.Unmarshal(file, &users)
	if err != nil {
		utils.HandleError(c, http.StatusInternalServerError, "Error unmarshal data", err)
		return
	}

	c.JSON(http.StatusOK, users)
}
