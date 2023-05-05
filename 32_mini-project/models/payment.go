package models

type Payment struct {
	ID            int    `json:"id"`
	TransactionID int    `json:"transaction_id"`
	Total         int    `json:"total"`
	Type          string `json:"type"`
	Pay           int    `json:"pay"`
	Change        int    `json:"change"`
	DateTime      string `json:"date_time"`
}
type NewPayment struct {
	TransactionID int    `json:"transaction_id"`
	Type          string `json:"type"`
	Pay           int    `json:"pay"`
}
