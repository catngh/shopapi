package models

type Product struct {
	ProductID string `db:"productId" json:"productId" gorm:"primaryKey; column:productId"`
	VendorID  string `db:"vendorId" json:"vendorId" gorm:"column:vendorId"`
	Name      string `db:"name" json:"name" `
	Price     string `db:"price" json:"price"`
}

func (Product) TableName() string {
	return "product"
}
