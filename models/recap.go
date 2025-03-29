package models

import "time"

type InstallmentRecap struct {
	ID          int       `json:"id"`
	UserID      int       `json:"user_id"`
	MerchantID  int       `json:"merchant_id"`
	TotalAmount float64   `json:"total_amount"`
	Remaining   float64   `json:"remaining"`
	Status      string    `json:"status"`
	Payments    []Payment `json:"payments"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
