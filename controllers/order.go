package controllers

import (
	"database/sql"
	"fmt"
	"strconv"

	"github.com/BerIincat/shopapi/database"
	"github.com/BerIincat/shopapi/models"
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
		cartId = cart[i].CartID
		i++
	}
	for _, ele := range ProdIdList {
		q := "SELECT * FROM product WHERE productId=" + ele
		db.Get(&product, q)
		products = append(products, product)
	}
	// products hold list of product object

	// get sub total by summing all product prices
	for _, ele := range products {
		price, _ := strconv.ParseFloat(ele.Price, 64)
		total += price
	}
	fmt.Println(total)
	subTotal := strconv.FormatFloat(total, 'f', -1, 64)
	fmt.Println(subTotal)
	q = "INSERT INTO `order`(cartId, subTotal) VALUES (" + cartId + "," + subTotal + ")"
	res, err := db.Exec(q)
	if err != nil {
		fmt.Print(err.Error())
		c.JSON(400, gin.H{"error": "order error"})
		return
	}

	temp, _ := res.LastInsertId()
	response.OrdId = string(strconv.FormatInt(temp, 10))
	response.Total = subTotal
	response.Items = products
	c.JSON(201, response)
}

type ResponseBody struct {
	OrdId string
	Items []models.Product
	Total string
}
