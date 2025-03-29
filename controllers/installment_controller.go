package controllers

import (
	"encoding/json"
	"net/http"
	"os"
	"test-mnc/models"
	"test-mnc/utils"
	"time"

	"github.com/gin-gonic/gin"
)

// CreateInstallment godoc
// @Summary Create new installment
// @Description Membuat transaksi baru
// @Tags installment
// @Accept json
// @Security BearerAuth
// @Produce json
// @Param installment body models.InstallmentRequest true "Data Cicilan"
// @Success 201 {object} models.Installment
// @Failure 400 {object} utils.ErrorResponse "Data tidak valid"
// @Failure 500 {object} utils.ErrorResponse "Kesalahan server"
// @Router /installment [post]
func CreateInstallment(c *gin.Context) {
	var request models.InstallmentRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		utils.HandleValidationError(c, err)
		return
	}

	newInstallment := models.Installment{
		UserID:      request.UserID,
		MerchantID:  request.MerchantID,
		TotalAmount: request.TotalAmount,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	const filePath = "data/installment.json"
	file, err := os.ReadFile(filePath)
	if err != nil && !os.IsNotExist(err) {
		utils.HandleDatabaseError(c, err)
		return
	}

	var installments []models.Installment
	if len(file) > 0 {
		if err = json.Unmarshal(file, &installments); err != nil {
			utils.HandleDatabaseError(c, err)
			return
		}
	}

	newInstallment.ID = 1
	if len(installments) > 0 {
		newInstallment.ID = installments[len(installments)-1].ID + 1
	}

	installments = append(installments, newInstallment)

	updatedData, err := json.MarshalIndent(installments, "", "  ")
	if err != nil {
		utils.HandleDatabaseError(c, err)
		return
	}

	if err = os.WriteFile(filePath, updatedData, 0644); err != nil {
		utils.HandleDatabaseError(c, err)
		return
	}

	c.JSON(http.StatusCreated, newInstallment)
}


// GetAllInstallment godoc
// @Summary Get all installments
// @Description Menampilkan daftar semua cicilan
// @Tags installment
// @Security BearerAuth
// @Produce json
// @Success 200 {array} models.Installment
// @Failure 500 {object} utils.ErrorResponse
// @Router /installment [get]
func GetAllInstallment(c *gin.Context) {
	const filePath = "data/installment.json"
	file, err := os.ReadFile(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			c.JSON(http.StatusOK, []models.Installment{})
			return
		}
		utils.HandleDatabaseError(c, err)
		return
	}

	if len(file) == 0 {
		c.JSON(http.StatusOK, []models.Installment{})
		return
	}

	var installments []models.Installment
	if err = json.Unmarshal(file, &installments); err != nil {
		utils.HandleDatabaseError(c, err)
		return
	}

	c.JSON(http.StatusOK, installments)
}
