package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"test-mnc/models"
	"test-mnc/utils"

	"github.com/gin-gonic/gin"
)

// GetAllMerchants godoc
// @Summary Get all merchants
// @Description Mengambil daftar semua merchant
// @Tags merchant
// @Security BearerAuth
// @Produce json
// @Success 200 {array} models.Merchant
// @Failure 500 {object} utils.ErrorResponse "Kesalahan server"
// @Router /merchants [get]
func GetAllMerchants(c *gin.Context) {

	file, err := os.ReadFile("data/merchant.json")
	if err != nil {
		utils.HandleDatabaseError(c, err)
		return
	}

	var merchants []models.Merchant
	if err = json.Unmarshal(file, &merchants); err != nil {
		utils.HandleDatabaseError(c, err)
		return
	}

	log.Printf("Successfully retrieved %d merchants", len(merchants))
	c.JSON(http.StatusOK, merchants)
}
