package database

import (
	"github.com/BerIincat/shopapi/models"
)

//var Order orderControl

type orderControl struct {
}

func Order() *orderControl {
	return &orderControl{}
}

func (u orderControl) Create(cartid string, total string) error {
	order := models.Order{CartID: cartid, SubTotal: total}
	result := DB.Select("CartID", "SubTotal").Create(&order)
	return result.Error
}
