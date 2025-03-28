package models

type Merchant struct {
	ID          int    `json:"id"`
	Name        string `json:"merchant_name"`
	BankAccount string `json:"bank_account"`
}
