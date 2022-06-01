package controllers

import (
	"github.com/gin-gonic/gin"
)

func GetCart(c *gin.Context) {
	return
}
func DelCartItem(c *gin.Context) {
	return
}
func AddCartItem(c *gin.Context) {
	return
}

type Product struct {
	ProductID string `json:"productId"`
	Name      string `json:"name"`
	Price     string `json:"price"`
}
