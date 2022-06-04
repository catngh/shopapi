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
	var ProdIdList []string
	cart := [20]models.Cart{} // Need to be dynamic !!!

	// Check for user id
	if !utils.UserIdExist(userId) {
		c.JSON(400, gin.H{"error": "user not found"})
		return
	}

	// Get items in cart
	rows, err := db.Queryx("SELECT * FROM cart WHERE userId=?", userId)
	if utils.PrintErrIfAny(err, 500, gin.H{"error": "database error"}, c) {
		return
	}

	// Parse result into struct list
	i := 0
	for rows.Next() {
		_ = rows.StructScan(&cart[i])
		ProdIdList = append(ProdIdList, cart[i].Item)
		i++
	}

	// Use item ids in cart to query for full product info => Inefficient
	for _, ele := range ProdIdList {
		db.Get(&product, "SELECT * FROM product WHERE productId=?", ele)
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
	err := c.BindJSON(&body)

	if utils.PrintErrIfAny(err, 400, gin.H{"error": "invalid request form"}, c) {
		return
	}

	if !validateInput(userId, body.ProdId, c) {
		return
	}

	// Prepare query
	_, err = db.Query("DELETE FROM cart WHERE userId=? AND item=?", userId, body.ProdId)
	if utils.PrintErrIfAny(err, 500, gin.H{"error": "database error"}, c) {
		return
	}

	c.JSON(201, gin.H{"status": "item deleted succesfully"})

}
func AddCartItem(c *gin.Context) {
	db := database.DB
	userId := c.Param("userid")
	body := ReqBody{} //To parse req body
	product := models.Product{}

	err := c.BindJSON(&body)
	if utils.PrintErrIfAny(err, 400, gin.H{"error": "invalid request form"}, c) {
		return
	}

	if !validateInput(userId, body.ProdId, c) {
		return
	}

	// Insert new item
	_, err = db.Query("INSERT INTO cart(userId,item) VALUES (?,?)", userId, product.ProductID)
	if utils.PrintErrIfAny(err, 400, gin.H{"error": "product not found"}, c) {
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
