package models

type Cart struct {
	CartID string    `json:"cartId"`
	UserID string    `json:"userId"`
	Items  []Product `json:"items"`
}
