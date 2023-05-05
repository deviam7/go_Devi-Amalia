package models

type Admin struct {
	ID       int    `json:"id"`
	Role     string `json:"role"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
