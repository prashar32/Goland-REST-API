package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/prashar32/rest-api-1/config"
	"github.com/prashar32/rest-api-1/models"
)

func GetUserInfo(c *gin.Context) {
	var user models.User
	c.BindJSON(&user)
	config.DB.Create(&user)
	c.JSON(200, &user)
}

func FetchUserInfo(c *gin.Context) {
	shopVisited := c.Query("shopVisited")
	result := []models.User{}
	config.DB.Raw("SELECT * FROM users WHERE shopVisited = ?", shopVisited).Scan(&result)
	c.JSON(200, result)
}

func AddUser(c *gin.Context) {
	var user models.User
	c.BindJSON(&user)
	config.DB.Create(&user)
	c.JSON(200, user)
}
