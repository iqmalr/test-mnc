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
