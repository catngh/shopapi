package controllers

import (
	"github.com/BerIincat/shopapi/models"
	"github.com/BerIincat/shopapi/utils"
	"github.com/gin-gonic/gin"
)

func (h *Handler) GetCart(c *gin.Context) {
	userId := c.Param("userid")
	product := models.Product{}
	var products []models.Product
	cart := []models.Cart{}

	// Check for user id
	if !h.db.User.IdExist(userId) {
		c.JSON(400, gin.H{"error": "user not found"})
		return
	}

	// Get items in cart
	cart, err := h.db.Cart.GetAllByUserID(userId)

	if utils.PrintErrIfAny(err, 500, gin.H{"error": "database error"}, c) {
		return
	}

	for _, ele := range cart {
		product, _ = h.db.Product.GetByID(ele.Item)
		products = append(products, product)
	}

	// products hold list of product object
	c.JSON(200, products)
	return
}
func (h *Handler) DelCartItem(c *gin.Context) {
	userId := c.Param("userid")
	body := ReqBody{} //To parse req body
	err := c.BindJSON(&body)

	if utils.PrintErrIfAny(err, 400, gin.H{"error": "invalid request form"}, c) {
		return
	}

	if !h.validateInput(userId, body.ProdId, c) {
		return
	}

	// Prepare query
	err = h.db.Cart.Delete(userId, body.ProdId)
	if utils.PrintErrIfAny(err, 500, gin.H{"error": "database error"}, c) {
		return
	}

	c.JSON(201, gin.H{"status": "item deleted succesfully"})

}
func (h *Handler) AddCartItem(c *gin.Context) {
	userId := c.Param("userid")
	body := ReqBody{} //To parse req body

	err := c.BindJSON(&body)
	if utils.PrintErrIfAny(err, 400, gin.H{"error": "invalid request form"}, c) {
		return
	}

	if !h.validateInput(userId, body.ProdId, c) {
		return
	}

	// Insert new item
	err = h.db.Cart.Create(userId, body.ProdId)

	if utils.PrintErrIfAny(err, 500, gin.H{"error": "database error"}, c) {
		return
	}

	c.JSON(201, gin.H{"status": "item added succesfully"})
}

type ReqBody struct {
	ProdId string `json:"productId"`
}

func (h *Handler) validateInput(userid string, prodid string, c *gin.Context) bool {
	// Check for user id
	if !h.db.User.IdExist(userid) {
		c.JSON(400, gin.H{"error": "user not found"})
		return false
	}

	// Check for product id
	if !h.db.Product.IdExist(prodid) {
		c.JSON(400, gin.H{"error": "product not found"})
		return false
	}
	return true
}
