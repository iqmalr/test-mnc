package models

import "time"

type Transaction struct {
	ID         int       `json:"id"`
	UserID     int       `json:"user_id"`
	MerchantID int       `json:"merchant_id"`
	Amount     float64   `json:"amount"`
	Status     string    `json:"status"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
