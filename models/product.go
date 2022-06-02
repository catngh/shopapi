package models

type Product struct {
	ProductID string `db:"productId" json:"productId"`
	VendorID  string `db:"vendorId" json:"vendorId"`
	Name      string `db:"name" json:"name"`
	Price     string `db:"price" json:"price"`
}
