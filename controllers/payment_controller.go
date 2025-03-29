package controllers

import (
	"encoding/json"
	"net/http"
	"os"
	"time"

	"test-mnc/models"

	"github.com/gin-gonic/gin"
)

func CreatePayment(c *gin.Context) {
	var newPayment models.Payment

	if err := c.ShouldBindJSON(&newPayment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Data tidak valid"})
		return
	}

	newPayment.CreatedAt = time.Now()
	newPayment.UpdatedAt = time.Now()

	file, err := os.ReadFile("data/payments.json")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error membayar"})
		return
	}

	var payments []models.Payment
	if len(file) > 0 {
		err = json.Unmarshal(file, &payments)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error membayar"})
			return
		}
	}

	newPayment.ID = 1
	if len(payments) > 0 {
		newPayment.ID = payments[len(payments)-1].ID + 1
	}

	payments = append(payments, newPayment)

	updatedData, err := json.MarshalIndent(payments, "", "  ")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengonversi JSON"})
		return
	}

	err = os.WriteFile("data/payments.json", updatedData, 0644)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menyimpan"})
		return
	}

	c.JSON(http.StatusCreated, newPayment)
}

func GetAllPayments(c *gin.Context) {
	file, err := os.ReadFile("data/payments.json")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error get data"})
		return
	}

	var payments []models.Payment
	err = json.Unmarshal(file, &payments)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error unmarshal data"})
		return
	}

	c.JSON(http.StatusOK, payments)
}
