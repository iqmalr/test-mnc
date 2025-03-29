package models

import "time"

type Installment struct {
	ID          int       `json:"id"`
	UserID      int       `json:"user_id"`
	MerchantID  int       `json:"merchant_id"`
	TotalAmount float64   `json:"total_amount"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
