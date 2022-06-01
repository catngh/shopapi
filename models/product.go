package models

type Product struct {
	ProductID string `json:"productId"`
	VendorID  string `json:"vendorId"`
	Name      string `json:"name"`
	Price     string `json:"price"`
}
