package models

type Cart struct {
	CartID string `db:"cartId" json:"cartId"`
	UserID string `db:"userId" json:"userId"`
	Item   string `db:"item" json:"item"`
}
