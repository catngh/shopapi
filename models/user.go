package models

type User struct {
	UserID   string `json:"userId"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
}
