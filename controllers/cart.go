package controllers

import (
	"github.com/BerIincat/shopapi/database"
	"github.com/BerIincat/shopapi/models"
	"github.com/BerIincat/shopapi/utils"
	"github.com/gin-gonic/gin"
)

func GetCart(c *gin.Context) {
	userId := c.Param("userid")
	product := models.Product{}
	var products []models.Product
	cart := []models.Cart{}

	// Check for user id
	if !database.User().IdExist(userId) {
		c.JSON(400, gin.H{"error": "user not found"})
		return
	}

	// Get items in cart
	cart, err := database.Cart().GetAllByUserID(userId)

	if utils.PrintErrIfAny(err, 500, gin.H{"error": "database error"}, c) {
		return
	}

	for _, ele := range cart {
		product, _ = database.Product().GetByID(ele.Item)
		products = append(products, product)
	}

	// products hold list of product object
	c.JSON(200, products)
	return
}
func DelCartItem(c *gin.Context) {
	userId := c.Param("userid")
	body := ReqBody{} //To parse req body
	err := c.BindJSON(&body)

	if utils.PrintErrIfAny(err, 400, gin.H{"error": "invalid request form"}, c) {
		return
	}

	if !validateInput(userId, body.ProdId, c) {
		return
	}

	// Prepare query
	err = database.Cart().Delete(userId, body.ProdId)
	if utils.PrintErrIfAny(err, 500, gin.H{"error": "database error"}, c) {
		return
	}

	c.JSON(201, gin.H{"status": "item deleted succesfully"})

}
func AddCartItem(c *gin.Context) {
	userId := c.Param("userid")
	body := ReqBody{} //To parse req body

	err := c.BindJSON(&body)
	if utils.PrintErrIfAny(err, 400, gin.H{"error": "invalid request form"}, c) {
		return
	}

	if !validateInput(userId, body.ProdId, c) {
		return
	}

	// Insert new item
	err = database.Cart().Create(userId, body.ProdId)

	if utils.PrintErrIfAny(err, 500, gin.H{"error": "database error"}, c) {
		return
	}

	c.JSON(201, gin.H{"status": "item added succesfully"})
}

type ReqBody struct {
	ProdId string `json:"productId"`
}

func validateInput(userid string, prodid string, c *gin.Context) bool {
	// Check for user id
	if !database.User().IdExist(userid) {
		c.JSON(400, gin.H{"error": "user not found"})
		return false
	}

	// Check for product id
	if !database.Product().IdExist(prodid) {
		c.JSON(400, gin.H{"error": "product not found"})
		return false
	}
	return true
}
