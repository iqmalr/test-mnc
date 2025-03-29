package controllers

import (
	"encoding/json"
	"fmt"
	"os"
	"test-mnc/models"
	"test-mnc/utils"

	"github.com/gin-gonic/gin"
)

// Login godoc
// @Summary Login user
// @Description Autentikasi user dengan email: "iqmalr@gmail.com" dan password: "password"
// @Tags auth
// @Accept json
// @Produce json
// @Param credentials body models.LoginRequest true "Email dan Password"
// @Success 200 {object} map[string]string "Token JWT"
// @Failure 400 {object} utils.ErrorResponse "Input tidak valid"
// @Failure 401 {object} utils.ErrorResponse "Login gagal"
// @Router /login [post]
func Login(c *gin.Context) {
	var req models.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.HandleValidationError(c, err)
		return
	}

	userFile, err := os.ReadFile("data/user.json")
	if err != nil {
		utils.HandleDatabaseError(c, err)
		return
	}

	var users []models.AuthUser
	if err = json.Unmarshal(userFile, &users); err != nil {
		utils.HandleDatabaseError(c, err)
		return
	}

	fmt.Println("Raw user file content:", string(userFile))

	fmt.Printf("Request email: %q, password: %q\n", req.Email, req.Password)
	for i, user := range users {
		fmt.Printf("User %d: email=%q, password=%q\n", i, user.Email, user.Password)
		fmt.Printf("Match email? %v, Match password? %v\n",
			user.Email == req.Email,
			user.Password == req.Password)
	}

	var loggedInUser *models.AuthUser
	for _, user := range users {
		fmt.Printf("Mencocokkan: %s dengan %s\n", req.Email, user.Email)
		if user.Email == req.Email && user.Password == req.Password {
			loggedInUser = &user
			break
		}
	}

	if loggedInUser == nil {
		utils.HandleError(c, 401, "Email atau password salah", nil)
		return
	}

	token, err := utils.GenerateJWT(loggedInUser)
	if err != nil {
		utils.HandleError(c, 500, "Gagal membuat token", err)
		return
	}

	tokenData := map[string]string{"token": token}
	tokenJSON, _ := json.MarshalIndent(tokenData, "", "  ")
	_ = os.WriteFile("data/token.json", tokenJSON, 0644)

	c.JSON(200, gin.H{"token": token})
}

// Logout godoc
// @Summary Logout user
// @Description Menghapus token JWT
// @Tags auth
// @Success 200 {string} string "Logout berhasil"
// @Router /logout [post]
func Logout(c *gin.Context) {
	_ = os.WriteFile("data/token.json", []byte("{}"), 0644)
	c.JSON(200, gin.H{"message": "Logout berhasil"})
}
