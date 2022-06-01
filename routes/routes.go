package routes

import (
	"github.com/BerIincat/shopapi/controllers"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func Setup() *gin.Engine {
	// Init the mux router
	router := gin.Default()
	//users
	router.GET("/login", controllers.Login)
	router.POST("/register", controllers.Register)
	//products
	router.GET("/products", controllers.GetProducts)
	router.GET("/inventory/:userid", controllers.GetUserInventory)
	//cart
	router.GET("/cart/{userId}", controllers.GetCart)
	router.POST("/cart/{userId}", controllers.AddCartItem)
	router.DELETE("/cart/{userId}", controllers.DelCartItem)
	//order
	router.POST("/order", controllers.NewOrder)
	return router
}
