package cart

import (
	"net/http"
)

func GetCart(res http.ResponseWriter, req *http.Request) {
	return
}
func DelItem(res http.ResponseWriter, req *http.Request) {
	return
}
func AddItem(res http.ResponseWriter, req *http.Request) {
	return
}

type Cart struct {
	CartID string    `json:"cartId"`
	UserID string    `json:"userId"`
	Items  []Product `json:"items"`
}
type Product struct {
	ProductID string `json:"productId"`
	Name      string `json:"name"`
	Price     string `json:"price"`
}
