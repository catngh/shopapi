package database

import (
	"github.com/BerIincat/shopapi/models"
)

//var Product productControl

type productControl struct {
}

func Product() *productControl {
	return &productControl{}
}
func (u productControl) GetByID(pid string) (models.Product, error) {
	product := models.Product{}
	result := DB.Where("productId=?", pid).First(&product)
	return product, result.Error
}

func (u productControl) GetAll() ([]models.Product, error) {
	products := []models.Product{}
	result := DB.Find(&products)
	return products, result.Error
}

func (u productControl) GetAllByUserID(uid string) ([]models.Product, error) {
	products := []models.Product{}
	result := DB.Where("vendorId=?", uid).Find(&products)
	return products, result.Error
}

func (u productControl) IdExist(pid string) bool {
	product := models.Product{}
	res := DB.Where("productId=?", pid).First(&product)
	if res.Error != nil {
		return false
	}
	return true
}
