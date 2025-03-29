package controllers

import (
	"encoding/json"
	"net/http"
	"os"
	"test-mnc/models"
	"test-mnc/utils"

	"github.com/gin-gonic/gin"
)

// GetInstallmentRecap godoc
// @Summary Get installment recap
// @Description Mengambil rekap cicilan dan status pembayaran
// @Tags installment
// @Security BearerAuth
// @Produce json
// @Success 200 {array} models.InstallmentRecap
// @Failure 500 {object} utils.ErrorResponse "Kesalahan server"
// @Router /recap [get]
func GetInstallmentRecap(c *gin.Context) {
	installmentFile, err := os.ReadFile("data/installment.json")
	if err != nil {
		utils.HandleDatabaseError(c, err)
		return
	}

	var installments []models.Installment
	if err = json.Unmarshal(installmentFile, &installments); err != nil {
		utils.HandleDatabaseError(c, err)
		return
	}

	paymentFile, err := os.ReadFile("data/payment.json")
	if err != nil {
		utils.HandleDatabaseError(c, err)
		return
	}

	var payments []models.Payment
	if err = json.Unmarshal(paymentFile, &payments); err != nil {
		utils.HandleDatabaseError(c, err)
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
