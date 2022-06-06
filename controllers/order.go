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
	db := database.DB
	product := models.Product{}
	response := ResponseBody{}
	cartId := ""
	var total float64
	var products []models.Product
	//products := []models.Product{}
	cart := []models.Cart{} // Need to be dynamic !!!

	// Check for user id
	if !utils.UserIdExist(userId) {
		c.JSON(400, gin.H{"error": "user not found"})
		return
	}

	// Get items in cart
	//rows, err := db.Queryx("SELECT * FROM cart WHERE userId=?", userId)
	result := db.Where("userId=?", userId).Find(&cart)
	if utils.PrintErrIfAny(result.Error, 500, gin.H{"error": "database error"}, c) {
		return
	}

	for _, ele := range cart {
		// Get product by id -> append to product list -> sum into total price
		db.Where("productId=?", ele.Item).Find(&product)
		products = append(products, product)
		price, _ := strconv.ParseFloat(product.Price, 64)
		total += price
		cartId = ele.CartID
	}
	subTotal := strconv.FormatFloat(total, 'f', -1, 64)
	order := models.Order{CartID: cartId, SubTotal: subTotal}

	result = db.Select("CartID", "SubTotal").Create(&order)
	//res, err := db.Exec("INSERT INTO `order`(cartId, subTotal) VALUES (?,?)", cartId, subTotal)
	if utils.PrintErrIfAny(result.Error, 400, gin.H{"error": "order error"}, c) {
		return
	}

	response = ResponseBody{OrdId: order.CartID, Total: subTotal, Items: products}
	c.JSON(201, response)
}

type ResponseBody struct {
	OrdId string
	Items []models.Product
	Total string
}
