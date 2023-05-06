package models

type Category struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Menus []*Menu `json:"menu" gorm:"-"`
}
