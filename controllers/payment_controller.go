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
// @Security BearerAuth
// @Produce json
// @Param payment body models.PaymentRequest true "Data Pembayaran"
// @Success 201 {object} models.Payment
// @Failure 400 {object} utils.ErrorResponse "Data tidak valid"
// @Failure 500 {object} utils.ErrorResponse "Kesalahan server"
// @Router /payments [post]
func CreatePayment(c *gin.Context) {
	var req models.PaymentRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		utils.HandleValidationError(c, err)
		return
	}

	installmentFile, err := os.ReadFile("data/installment.json")
	if err != nil && !os.IsNotExist(err) {
		utils.HandleDatabaseError(c, err)
		return
	}

	var installments []models.Installment
	if len(installmentFile) > 0 {
		if err = json.Unmarshal(installmentFile, &installments); err != nil {
			utils.HandleDatabaseError(c, err)
			return
		}
	}

	var totalAmount float64
	for _, inst := range installments {
		if inst.ID == req.TransactionID {
			totalAmount = inst.TotalAmount
			break
		}
	}
	if totalAmount == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Transaksi tidak ditemukan"})
		return
	}

	paymentFile, err := os.ReadFile("data/payment.json")
	if err != nil && !os.IsNotExist(err) {
		utils.HandleDatabaseError(c, err)
		return
	}

	var payments []models.Payment
	if len(paymentFile) > 0 {
		if err = json.Unmarshal(paymentFile, &payments); err != nil {
			utils.HandleDatabaseError(c, err)
			return
		}
	}

	totalPaid := 0.0
	for _, payment := range payments {
		if payment.TransactionID == req.TransactionID {
			totalPaid += payment.Amount
		}
	}

	if totalPaid+req.Amount > totalAmount {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Pembayaran melebihi total tagihan (sudah lunas)"})
		return
	}

	newPayment := models.Payment{
		TransactionID: req.TransactionID,
		Amount:        req.Amount,
		PaymentMethod: req.PaymentMethod,
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
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

	if err = os.WriteFile("data/payment.json", updatedData, 0644); err != nil {
		utils.HandleDatabaseError(c, err)
		return
	}

	c.JSON(http.StatusCreated, newPayment)
}



// GetAllPayment godoc
// @Summary Get all payment
// @Description Mengambil daftar semua pembayaran
// @Tags payment
// @Security BearerAuth
// @Produce json
// @Success 200 {array} models.Payment
// @Failure 500 {object} utils.ErrorResponse "Kesalahan server"
// @Router /payments [get]
func GetAllPayment(c *gin.Context) {
	file, err := os.ReadFile("data/payment.json")
	if err != nil {
		if os.IsNotExist(err) {
			c.JSON(http.StatusOK, []models.Payment{})
			return
		}
		utils.HandleDatabaseError(c, err)
		return
	}

	var payment []models.Payment
	if len(file) == 0 {
		c.JSON(http.StatusOK, []models.Payment{})
		return
	}

	if err = json.Unmarshal(file, &payment); err != nil {
		utils.HandleDatabaseError(c, err)
		return
	}

	c.JSON(http.StatusOK, payment)
}
