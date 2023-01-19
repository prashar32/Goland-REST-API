package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/prashar32/rest-api-1/controller"
)

func UserRoute(router *gin.Engine) {
	userapi := router.Group("/api/user")
	{
		userapi.GET("/getInfo", controller.GetUserInfo)
		userapi.GET("/fetchUser", controller.FetchUserInfo)
		userapi.POST("/addUser", controller.AddUser)

	}
}
