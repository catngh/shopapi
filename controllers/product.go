package controllers

import (
	"database/sql"

	"github.com/BerIincat/shopapi/database"
	"github.com/BerIincat/shopapi/models"
	"github.com/BerIincat/shopapi/utils"
	"github.com/gin-gonic/gin"
)

func GetProducts(c *gin.Context) {
	db := database.DB
	products := []models.Product{}

	// Query and bind to products
	err := db.Select(&products, "SELECT * FROM product")
	if err != nil {
		c.JSON(500, gin.H{"error": "database error"})
		return
	}
	c.IndentedJSON(200, products)
}
func GetUserInventory(c *gin.Context) {
	id := c.Param("userid")
	db := database.DB
	products := []models.Product{}
	user := models.User{}

	// Query and bind to user
	err := db.Get(&user, "SELECT * FROM usr WHERE userId=?", id)

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
	err = db.Select(&products, "SELECT * FROM product WHERE vendorId='"+id+"'")
	if utils.PrintErrIfAny(err, 500, gin.H{"error": "database error"}, c) {
		return
	}
	c.IndentedJSON(200, products)
}
