package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/prashar32/rest-api-1/controller"
)

func DetailRoutes(router *gin.Engine) {
	detailapi := router.Group("/api/shop/user/")
	{
		// http://localhost:8080/api/shop/user/getInfo?id (This will give all the users and shops corresponding to particular shop visited)
		detailapi.GET("/getInfo/:id", controller.GetDetails)
	}
}
