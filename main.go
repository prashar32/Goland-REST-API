package main

import (
	"github.com/gin-gonic/gin"
	"github.com/prashar32/rest-api-1/config"
	"github.com/prashar32/rest-api-1/controller"
	"github.com/prashar32/rest-api-1/routes"
)

func main() {
	router := gin.New()

	// Connect the database
	config.Connect()

	// Initialize at Cache
	controller.CacheIntialize("localhost:6379", 0, 100)

	// Routes for Shops
	routes.ShopRoutes(router)

	// Routes for Users
	routes.UserRoute(router)

	// Routes to handle concurrent details of Users with Shop
	routes.DetailRoutes(router)

	router.Run(":8080")
}
