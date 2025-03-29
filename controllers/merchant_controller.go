package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"test-mnc/models"

	"github.com/gin-gonic/gin"
)

func GetAllMerchants(c *gin.Context) {
	log.Println("fetching all merchants...")
	file, err := os.ReadFile("data/merchant.json")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error get data"})
		return
	}

	var merchants []models.Merchant
	err = json.Unmarshal(file, &merchants)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Eror umarshal data"})
		return
	}
	log.Printf("successfully retrieved %d merchants", len(merchants))
	c.JSON(http.StatusOK, merchants)
}
