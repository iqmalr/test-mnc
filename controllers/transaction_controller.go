package controllers

import (
	"encoding/json"
	"net/http"
	"os"
	"test-mnc/models"
	"time"

	"github.com/gin-gonic/gin"
)

func CreateTransaction(c *gin.Context) {
	var newTransaction models.Transaction

	if err := c.ShouldBindJSON(&newTransaction); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Data tidak valid"})
		return
	}

	newTransaction.CreatedAt = time.Now()
	newTransaction.UpdatedAt = time.Now()

	file, err := os.ReadFile("data/transaction.json")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal membaca file transaksi"})
		return
	}

	var transactions []models.Transaction
	if len(file) > 0 {
		err = json.Unmarshal(file, &transactions)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengurai JSON transaksi"})
			return
		}
	}

	newTransaction.ID = 1
	if len(transactions) > 0 {
		newTransaction.ID = transactions[len(transactions)-1].ID + 1
	}

	transactions = append(transactions, newTransaction)

	updatedData, err := json.MarshalIndent(transactions, "", "  ")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengonversi JSON transaksi"})
		return
	}
	err = os.WriteFile("data/transaction.json", updatedData, 0644)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menyimpan transaksi"})
		return
	}

	c.JSON(http.StatusCreated, newTransaction)
}
