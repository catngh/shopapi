package user

import (
	"net/http"
)

func Login(res http.ResponseWriter, req *http.Request) {
	return
}
func Register(res http.ResponseWriter, req *http.Request) {
	return
}

type User struct {
	UserID   string `json:"userId"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
}
