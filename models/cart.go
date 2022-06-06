package models

type Cart struct {
	CartID string `db:"cartId" json:"cartId" gorm:"primaryKey; column:cartId"`
	UserID string `db:"userId" json:"userId" gorm:"column: userId"`
	Item   string `db:"item" json:"item"`
}

func (Cart) TableName() string {
	return "cart"
}
