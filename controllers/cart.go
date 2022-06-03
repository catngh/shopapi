package controllers

import (
	"database/sql"
	"fmt"

	"github.com/BerIincat/shopapi/database"
	"github.com/BerIincat/shopapi/models"
	"github.com/gin-gonic/gin"
)

func GetCart(c *gin.Context) {
	userId := c.Param("userid")
	db := database.DB
	product := models.Product{}
	var products []models.Product
	var ProdIdList []string
	//products := []models.Product{}
	cart := [20]models.Cart{} // Need to be dynamic !!!

	// Check for user id
	q := "SELECT * FROM usr WHERE userId=" + userId
	row := db.QueryRow(q)
	if row.Scan() == sql.ErrNoRows {
		c.JSON(400, gin.H{"error": "user not found"})
		return
	}

	// Get items in cart
	q = "SELECT * FROM cart WHERE userId=" + userId
	rows, err := db.Queryx(q)
	if err != nil {
		c.JSON(500, gin.H{"error": "database error"})
		return
	}
	i := 0
	for rows.Next() {
		_ = rows.StructScan(&cart[i])
		ProdIdList = append(ProdIdList, cart[i].Item)
		i++
	}
	for _, ele := range ProdIdList {
		q := "SELECT * FROM product WHERE productId=" + ele
		db.Get(&product, q)
		products = append(products, product)
	}
	// ProdIdList has list of id, need to use this list to out put product object
	c.JSON(200, products)
	return
}
func DelCartItem(c *gin.Context) {
	db := database.DB
	userId := c.Param("userid")
	body := ReqBody{} //To parse req body
	err := c.BindJSON(&body)
	if err != nil {
		c.JSON(400, gin.H{"error": "invalid request form"})
		return
	}

	// Check for user id
	q := "SELECT * FROM usr WHERE userId=" + userId
	row := db.QueryRow(q)
	if row.Scan() == sql.ErrNoRows {
		c.JSON(400, gin.H{"error": "user not found"})
		return
	}

	// Check for product id
	q = "SELECT * FROM product WHERE productId=" + body.ProdId
	row = db.QueryRow(q)
	if row.Scan() == sql.ErrNoRows {
		c.JSON(400, gin.H{"error": "product not found"})
		return
	}

	// Delete cart
	q = "DELETE FROM cart WHERE userId=" + userId + " AND item='" + body.ProdId + "'"
	fmt.Print(q)

	_, err = db.Query(q)
	if err != nil {
		fmt.Print(err.Error())
		c.JSON(500, gin.H{"error": "database error"})
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
	if err != nil {
		c.JSON(400, gin.H{"error": "invalid request form"})
		return
	}

	// Check for user id
	q := "SELECT * FROM usr WHERE userId=" + userId
	row := db.QueryRow(q)
	if row.Scan() == sql.ErrNoRows {
		c.JSON(400, gin.H{"error": "user not found"})
		return
	}

	// Check for product id
	q = "SELECT * FROM product WHERE productId=" + body.ProdId
	err = db.Get(&product, q)
	if err != nil {
		c.JSON(400, gin.H{"error": "product not found"})
		return
	}

	// Insert new item
	q = "INSERT INTO cart(userId,item) VALUES (" + userId + "," + product.ProductID + ")"
	_, err = db.Query(q)
	if err != nil {
		fmt.Print(err.Error())
		c.JSON(400, gin.H{"error": "product not found"})
		return
	}

	c.JSON(201, gin.H{"status": "item added succesfully"})
}

type ReqBody struct {
	ProdId string `json:"productId"`
}
