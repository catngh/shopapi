package product

import (
	"net/http"

	"github.com/gorilla/mux"
)

func GetAll(res http.ResponseWriter, req *http.Request) {
	return
}
func GetInventory(res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	userId := vars["userId"]
}

type Product struct {
	ProductID string `json:"productId"`
	Name      string `json:"name"`
	Price     string `json:"price"`
}
