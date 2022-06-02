package models

type Cart struct {
	CartID string    `db:"cartId" json:"cartId"`
	UserID string    `db:"userId" json:"userId"`
	Items  []Product `db:"items" json:"items"`
}
