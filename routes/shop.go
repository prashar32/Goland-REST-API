package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/prashar32/rest-api-1/controller"
)

func ShopRoutes(router *gin.Engine) {
	shopapi := router.Group("/api/shop")
	{
		shopapi.GET("/getAllShops", controller.GetShop)
		shopapi.GET("/queryShops", controller.QueryShops)
		shopapi.POST("/addShop", controller.CreateShop)

	}
}
