package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/prashar32/rest-api-1/controller"
)

func DetailRoutes(router *gin.Engine) {
	detailapi := router.Group("/api/shop/user/")
	{
		detailapi.GET("/getInfo/:id", controller.GetDetails)
	}
}
