package database

import (
	"github.com/BerIincat/shopapi/models"
)

//var Cart cartControl

type cartControl struct {
}

func Cart() *cartControl {
	return &cartControl{}
}

func (u cartControl) GetAllByUserID(uid string) ([]models.Cart, error) {
	cart := []models.Cart{}
	result := DB.Where("userId=?", uid).Find(&cart)
	return cart, result.Error
}
func (u cartControl) Delete(uid string, pid string) error {
	cart := models.Cart{}
	result := DB.Where("userId=? AND item=?", uid, pid).Delete(&cart)
	return result.Error
}
func (u cartControl) Create(uid string, pid string) error {
	cart := models.Cart{UserID: uid, Item: pid}
	result := DB.Select("UserID", "Item").Create(&cart)
	return result.Error
}
