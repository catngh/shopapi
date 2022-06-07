package routes

import (
	"github.com/BerIincat/shopapi/controllers"
	"github.com/BerIincat/shopapi/database"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func Setup(d *database.DbControllers) *gin.Engine {
	handler := controllers.NewHandler(d)
	// Init the mux router
	router := gin.Default()
	//users
	router.GET("/login", handler.Login)
	router.POST("/register", handler.Register)
	//products
	router.GET("/products", handler.GetProducts)
	router.GET("/inventory/:userid", handler.GetUserInventory)
	//cart
	router.GET("/cart/:userid", handler.GetCart)
	router.POST("/cart/:userid", handler.AddCartItem)
	router.DELETE("/cart/:userid", handler.DelCartItem)
	//order
	router.POST("/order/:userid", handler.NewOrder)
	return router
}
