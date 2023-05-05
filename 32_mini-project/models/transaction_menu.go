package models

type TransactionMenu struct {
	ID            int `json:"id"`
	TransactionID int `json:"transaction_id"`
	MenuID        int `json:"menu_id"`
}
