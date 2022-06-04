package controllers

import (
	"strconv"

	"github.com/BerIincat/shopapi/database"
	"github.com/BerIincat/shopapi/models"
	"github.com/BerIincat/shopapi/utils"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func NewOrder(c *gin.Context) {
	userId := c.Param("userid")
	db := database.DB
	product := models.Product{}
	response := ResponseBody{}
	cartId := ""
	var total float64
	var products []models.Product
	var prodIdList []string
	//products := []models.Product{}
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

	// i := 0
	// for rows.Next() {
	// 	_ = rows.StructScan(&cart[i])
	// 	ProdIdList = append(ProdIdList, cart[i].Item)
	// 	cartId = cart[i].CartID
	// 	i++
	// }
	prodIdList, cartId = getQueryResult(rows, cart)

	// Use item ids in cart to query for full product info => Inefficient
	for _, ele := range prodIdList {
		db.Get(&product, "SELECT * FROM product WHERE productId=?", ele)
		products = append(products, product)
	}

	// get sub total by summing all product prices
	for _, ele := range products {
		price, _ := strconv.ParseFloat(ele.Price, 64)
		total += price
	}

	subTotal := strconv.FormatFloat(total, 'f', -1, 64)
	res, err := db.Exec("INSERT INTO `order`(cartId, subTotal) VALUES (?,?)", cartId, subTotal)
	if utils.PrintErrIfAny(err, 400, gin.H{"error": "order error"}, c) {
		return
	}

	temp, _ := res.LastInsertId()
	response.init(string(strconv.FormatInt(temp, 10)), subTotal, products)
	c.JSON(201, response)
}

type ResponseBody struct {
	OrdId string
	Items []models.Product
	Total string
}

func (r *ResponseBody) init(id string, total string, items []models.Product) {
	r.OrdId = id
	r.Total = total
	r.Items = items
}

func getQueryResult(rows *sqlx.Rows, cart [20]models.Cart) ([]string, string) {
	var productList []string
	var cartId string
	i := 0
	for rows.Next() {
		_ = rows.StructScan(&cart[i])
		productList = append(productList, cart[i].Item)
		cartId = cart[i].CartID
		i++
	}
	return productList, cartId
}
