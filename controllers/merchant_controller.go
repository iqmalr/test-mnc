package controllers

import (
	"encoding/json"
	"net/http"
	"os"
	"test-mnc/models"

	"github.com/gin-gonic/gin"
)

func GetAllMerchants(c *gin.Context) {
	file, err := os.ReadFile("data/merchant.json")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal membaca file"})
		return
	}

	var merchants []models.Merchant
	err = json.Unmarshal(file, &merchants)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengurai JSON"})
		return
	}

	c.JSON(http.StatusOK, merchants)
}
