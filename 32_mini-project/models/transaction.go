package models

type Transaction struct {
	ID       int     `json:"id"`
	Total    int     `json:"total"`
	DateTime string  `json:"date_time"`
	Menus    []*Menu `json:"menus" gorm:"-"`
}

type NewTransaction struct {
	Items []int `json:"items"`
}
