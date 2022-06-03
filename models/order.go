package models

type Order struct {
	OrderID     string `db:"orderId" json:"orderId"`
	CartID      string `db:"cartId" json:"cartId"`
	SubTotal    string `db:"subTotal" json:"subTotal"`
	TimeCreated string `db:"timeCreated" json:"timeCreated"`
}
