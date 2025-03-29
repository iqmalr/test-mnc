package controllers

import (
	"encoding/json"
	"net/http"
	"os"
	"time"

	"test-mnc/models"
	"test-mnc/utils"

	"github.com/gin-gonic/gin"
)

// CreatePayment godoc
// @Summary Create new payment
// @Description Membuat pembayaran baru
// @Tags payment
// @Accept json
// @Produce json
// @Param payment body models.Payment true "Data Pembayaran"
// @Success 201 {object} models.Payment
// @Failure 400 {object} utils.ErrorResponse "Data tidak valid"
// @Failure 500 {object} utils.ErrorResponse "Kesalahan server"
// @Router /payments [post]
func CreatePayment(c *gin.Context) {
	var newPayment models.Payment

	if err := c.ShouldBindJSON(&newPayment); err != nil {
		utils.HandleValidationError(c, err)
		return
	}

	newPayment.CreatedAt = time.Now()
	newPayment.UpdatedAt = time.Now()

	file, err := os.ReadFile("data/payments.json")
	if err != nil && !os.IsNotExist(err) {
		utils.HandleDatabaseError(c, err)
		return
	}

	var payments []models.Payment
	if len(file) > 0 {
		if err = json.Unmarshal(file, &payments); err != nil {
			utils.HandleDatabaseError(c, err)
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
		utils.HandleDatabaseError(c, err)
		return
	}

	if err = os.WriteFile("data/payments.json", updatedData, 0644); err != nil {
		utils.HandleDatabaseError(c, err)
		return
	}

	c.JSON(http.StatusCreated, newPayment)
}

// GetAllPayments godoc
// @Summary Get all payments
// @Description Mengambil daftar semua pembayaran
// @Tags payment
// @Produce json
// @Success 200 {array} models.Payment
// @Failure 500 {object} utils.ErrorResponse "Kesalahan server"
// @Router /payments [get]
func GetAllPayments(c *gin.Context) {
	file, err := os.ReadFile("data/payments.json")
	if err != nil {
		if os.IsNotExist(err) {
			c.JSON(http.StatusOK, []models.Payment{})
			return
		}
		utils.HandleDatabaseError(c, err)
		return
	}

	var payments []models.Payment
	if len(file) == 0 {
		c.JSON(http.StatusOK, []models.Payment{})
		return
	}

	if err = json.Unmarshal(file, &payments); err != nil {
		utils.HandleDatabaseError(c, err)
		return
	}

	c.JSON(http.StatusOK, payments)
}
