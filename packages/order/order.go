package order

import (
	"net/http"
)

func NewOrder(res http.ResponseWriter, req *http.Request) {

}

type Order struct {
	CartID   string `json:"cartId"`
	SubTotal string `json:"subTotal"`
}
