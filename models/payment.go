package models

import "time"

type Payment struct {
	ID            int       `json:"id"`
	TransactionID int       `json:"transaction_id"`
	UserID        int       `json:"user_id"`
	MerchantID    int       `json:"merchant_id"`
	Amount        float64   `json:"amount"`
	PaymentMethod string    `json:"payment_method"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

type PaymentRequest struct {
	TransactionID int     `json:"transaction_id" example:"1"`
	UserID        int     `json:"user_id" example:"1"`
	MerchantID    int     `json:"merchant_id" example:"2"`
	Amount        float64 `json:"amount" example:"200000"`
	PaymentMethod string  `json:"payment_method" example:"bank_transfer"`
}