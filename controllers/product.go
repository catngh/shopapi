package controllers

import (
	"database/sql"

	"github.com/BerIincat/shopapi/models"
	"github.com/BerIincat/shopapi/utils"
	"github.com/gin-gonic/gin"
)

func (h *Handler) GetProducts(c *gin.Context) {
	products, err := h.db.Product.GetAll()
	if err != nil {
		c.JSON(500, gin.H{"error": "database error"})
		return
	}
	c.IndentedJSON(200, products)
}
func (h *Handler) GetUserInventory(c *gin.Context) {
	id := c.Param("userid")
	products := []models.Product{}
	user := models.User{}

	// Get user from db
	user, err := h.db.User.GetById(id)

	// Check user id and role
	if err == sql.ErrNoRows {
		c.JSON(400, gin.H{"error": "user not found"})
		return
	}
	if utils.PrintErrIfAny(err, 500, gin.H{"error": "database error"}, c) {
		return
	}
	if user.Role != "vendor" {
		c.JSON(400, gin.H{"error": "user is not a vendor"})
		return
	}

	// Get user inventory and bind to products

	products, err = h.db.Product.GetAllByUserID(id)
	if utils.PrintErrIfAny(err, 500, gin.H{"error": "database error"}, c) {
		return
	}
	c.IndentedJSON(200, products)
}
