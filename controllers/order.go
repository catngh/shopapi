package controllers

import (
	"strconv"

	"github.com/BerIincat/shopapi/database"
	"github.com/BerIincat/shopapi/models"
	"github.com/BerIincat/shopapi/utils"
	"github.com/gin-gonic/gin"
)

func NewOrder(c *gin.Context) {
	userId := c.Param("userid")
	product := models.Product{}
	response := ResponseBody{}
	cartId, total := "", ""
	var products []models.Product
	cart := []models.Cart{}

	// Check for user id
	if !database.User.IdExist(userId) {
		c.JSON(400, gin.H{"error": "user not found"})
		return
	}

	// Get items in cart
	cart, err := database.Cart.GetAllByUserID(userId)
	if utils.PrintErrIfAny(err, 500, gin.H{"error": "database error"}, c) {
		return
	}

	for _, ele := range cart {
		// Get product by id -> append to product list -> sum into total price
		product, _ = database.Product.GetByID(ele.Item)
		products = append(products, product)
		cartId = ele.CartID
	}
	total = getTotalPrice(products)

	err = database.Order.Create(cartId, total)
	if utils.PrintErrIfAny(err, 400, gin.H{"error": "order error"}, c) {
		return
	}
	response = ResponseBody{OrdId: cartId, Total: total, Items: products}
	c.JSON(201, response)
}

type ResponseBody struct {
	OrdId string
	Items []models.Product
	Total string
}

func getTotalPrice(products []models.Product) string {
	var sum float64 = 0
	for _, ele := range products {
		price, _ := strconv.ParseFloat(ele.Price, 64)
		sum += price
	}
	return strconv.FormatFloat(sum, 'f', -1, 64)
}
