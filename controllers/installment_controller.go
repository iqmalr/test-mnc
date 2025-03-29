package controllers

import (
	"encoding/json"
	"net/http"
	"os"
	"test-mnc/models"
	"time"

	"github.com/gin-gonic/gin"
)

func CreateInstallment(c *gin.Context) {
	var newInstallment models.Installment

	if err := c.ShouldBindJSON(&newInstallment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Data tidak valid"})
		return
	}

	newInstallment.CreatedAt = time.Now()
	newInstallment.UpdatedAt = time.Now()

	file, err := os.ReadFile("data/installment.json")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error transaksi"})
		return
	}

	var installment []models.Installment
	if len(file) > 0 {
		err = json.Unmarshal(file, &installment)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error transaksi"})
			return
		}
	}

	newInstallment.ID = 1
	if len(installment) > 0 {
		newInstallment.ID = installment[len(installment)-1].ID + 1
	}

	installment = append(installment, newInstallment)

	updatedData, err := json.MarshalIndent(installment, "", "  ")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengonversi JSON transaksi"})
		return
	}
	err = os.WriteFile("data/transaction.json", updatedData, 0644)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menyimpan transaksi"})
		return
	}

	c.JSON(http.StatusCreated, newInstallment)
}

func GetAllInstallment(c *gin.Context) {
	file, err := os.ReadFile("data/installment.json")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error get data"})
		return
	}

	var installment []models.Installment
	err = json.Unmarshal(file, &installment)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Eror umarshal data"})
		return
	}

	c.JSON(http.StatusOK, installment)
}
