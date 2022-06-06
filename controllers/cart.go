package controllers

import (
	"github.com/BerIincat/shopapi/database"
	"github.com/BerIincat/shopapi/models"
	"github.com/BerIincat/shopapi/utils"
	"github.com/gin-gonic/gin"
)

func GetCart(c *gin.Context) {
	userId := c.Param("userid")
	db := database.DB
	product := models.Product{}
	var products []models.Product
	cart := []models.Cart{}

	// Check for user id
	if !utils.UserIdExist(userId) {
		c.JSON(400, gin.H{"error": "user not found"})
		return
	}

	// Get items in cart
	result := db.Where("userId=?", userId).Find(&cart)
	//rows, err := db.Queryx("SELECT * FROM cart WHERE userId=?", userId)

	if utils.PrintErrIfAny(result.Error, 500, gin.H{"error": "database error"}, c) {
		return
	}

	for _, ele := range cart {
		db.Where("productId=?", ele.Item).Find(&product)
		products = append(products, product)
	}

	// products hold list of product object
	c.JSON(200, products)
	return
}
func DelCartItem(c *gin.Context) {
	db := database.DB
	userId := c.Param("userid")
	body := ReqBody{} //To parse req body
	cart := models.Cart{}
	err := c.BindJSON(&body)

	if utils.PrintErrIfAny(err, 400, gin.H{"error": "invalid request form"}, c) {
		return
	}

	if !validateInput(userId, body.ProdId, c) {
		return
	}

	// Prepare query
	result := db.Where("userId=? AND item=?", userId, body.ProdId).Delete(&cart)
	//_, err = db.Query("DELETE FROM cart WHERE userId=? AND item=?", userId, body.ProdId)
	if utils.PrintErrIfAny(result.Error, 500, gin.H{"error": "database error"}, c) {
		return
	}

	c.JSON(201, gin.H{"status": "item deleted succesfully"})

}
func AddCartItem(c *gin.Context) {
	db := database.DB
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
	cart := models.Cart{UserID: userId, Item: body.ProdId}
	result := db.Select("UserID", "Item").Create(&cart)
	//_, err = db.Query("INSERT INTO cart(userId,item) VALUES (?,?)", userId, .ProductID)

	if utils.PrintErrIfAny(result.Error, 500, gin.H{"error": "database error"}, c) {
		return
	}

	c.JSON(201, gin.H{"status": "item added succesfully"})
}

type ReqBody struct {
	ProdId string `json:"productId"`
}

func validateInput(userid string, prodid string, c *gin.Context) bool {
	// Check for user id
	if !utils.UserIdExist(userid) {
		c.JSON(400, gin.H{"error": "user not found"})
		return false
	}

	// Check for product id
	if !utils.ProductIdExist(prodid) {
		c.JSON(400, gin.H{"error": "product not found"})
		return false
	}
	return true
}
