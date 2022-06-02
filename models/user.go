package models

type User struct {
	UserID   string `db:"userId" json:"userId"`
	Email    string `db:"email" json:"email"`
	Password string `db:"password" json:"password"`
	Role     string `db:"role" json:"role"`
}
