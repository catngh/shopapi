package models

type Order struct {
	OrderID     string `db:"orderId" json:"orderId" gorm:"primaryKey; column:orderId"`
	CartID      string `db:"cartId" json:"cartId" gorm:"column:cartId"`
	SubTotal    string `db:"subTotal" json:"subTotal" gorm:"column:subTotal"`
	TimeCreated string `db:"timeCreated" json:"timeCreated" gorm:"column:timeCreated"`
}

func (Order) TableName() string {
	return "order"
}
