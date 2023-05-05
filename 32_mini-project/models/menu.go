package models

type Menu struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	CategoryID  int     `json:"category_id"`
	Description *string `json:"description"`
	Price       int     `json:"price"`
}

type NewMenu struct {
	Name        string  `json:"name"`
	CategoryID  int     `json:"category_id"`
	Description *string `json:"description"`
	Price       int     `json:"price"`
}
