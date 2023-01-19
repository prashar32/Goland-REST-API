package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/prashar32/rest-api-1/controller"
)

func UserRoute(router *gin.Engine) {
	userapi := router.Group("/api/user")
	{
		// http://localhost:8080/api/user/getInfo (This will give the all the users in the database)
		userapi.GET("/getInfo", controller.GetUserInfo)

		// http://localhost:8080/api/user?shopVisited (This will give the users corresponding to the particular shop)
		userapi.GET("/fetchUser", controller.FetchUserInfo)

		// http://localhost:8080/api/user/addUser (This will add a user)
		userapi.POST("/addUser", controller.AddUser)

	}
}
