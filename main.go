package main

import (
	"github.com/gin-gonic/gin"
	"github.com/prashar32/rest-api-1/config"
	"github.com/prashar32/rest-api-1/controller"
	"github.com/prashar32/rest-api-1/routes"
)

func main() {
	router := gin.New()
	config.Connect()
	controller.CacheIntialize("localhost:6379", 0, 100)
	routes.ShopRoutes(router)
	routes.UserRoute(router)
	routes.DetailRoutes(router)
	router.Run(":8080")
}
