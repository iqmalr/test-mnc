package controllers

import (
	"encoding/json"
	"net/http"
	"os"
	"test-mnc/models"

	"github.com/gin-gonic/gin"
)

func GetInstallmentRecap(c *gin.Context) {
	installmentFile, err := os.ReadFile("data/installment.json")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal get data installment"})
		return
	}

	var installments []models.Installment
	err = json.Unmarshal(installmentFile, &installments)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal decode data installment"})
		return
	}

	paymentFile, err := os.ReadFile("data/payment.json")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal membaca data payment"})
		return
	}

	var payments []models.Payment
	err = json.Unmarshal(paymentFile, &payments)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal decode data payment"})
		return
	}

	var history []models.InstallmentRecap
	for _, installment := range installments {
		var relatedPayments []models.Payment
		totalPaid := 0.0

		for _, payment := range payments {
			if payment.TransactionID == installment.ID {
				relatedPayments = append(relatedPayments, payment)
				totalPaid += payment.Amount
			}
		}

		remaining := installment.TotalAmount - totalPaid
		status := "Belum Lunas"
		if remaining <= 0 {
			status = "Lunas"
		}

		history = append(history, models.InstallmentRecap{
			ID:          installment.ID,
			UserID:      installment.UserID,
			MerchantID:  installment.MerchantID,
			TotalAmount: installment.TotalAmount,
			Remaining:   remaining,
			Status:      status,
			Payments:    relatedPayments,
			CreatedAt:   installment.CreatedAt,
			UpdatedAt:   installment.UpdatedAt,
		})
	}

	c.JSON(http.StatusOK, history)
}
