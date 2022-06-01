package models

import "time"

type Order struct {
	OrderID     string    `json:"orderId"`
	CartID      string    `json:"cartId"`
	SubTotal    string    `json:"subTotal"`
	TimeCreated time.Time `json:"timeCreated"`
}
