package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/prashar32/rest-api-1/controller"
)

func ShopRoutes(router *gin.Engine) {
	shopapi := router.Group("/api/shop")
	{
		// http://localhost:8080/api/shop/getAllShops (This will give the all the shops in the mall)
		shopapi.GET("/getAllShops", controller.GetShop)

		// http://localhost:8080/api/shop/queryShops?shopCategory (This will give all the shops corresponding to the particular category)
		shopapi.GET("/queryShops", controller.QueryShops)

		// http://localhost:8080/api/shop/addShop (This will add a shop)
		shopapi.POST("/addShop", controller.CreateShop)
	}
}
