package models

type Category struct {
	ID    int     `json:"id"`
	Name  string  `json:"Name"`
	Menus []*Menu `json:"menu" gorm:"-"`
}
