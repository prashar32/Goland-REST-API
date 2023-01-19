package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/prashar32/rest-api-1/config"
	"github.com/prashar32/rest-api-1/models"
)

// Get all users
func GetUserInfo(c *gin.Context) {
	var user models.User
	c.BindJSON(&user)
	config.DB.Create(&user)
	c.JSON(200, &user)
}

// Get all users corresponding to the particular shop visited
func FetchUserInfo(c *gin.Context) {
	shopVisited := c.Query("shopVisited")
	result := []models.User{}
	config.DB.Raw("SELECT * FROM users WHERE shopVisited = ?", shopVisited).Scan(&result)
	c.JSON(200, result)
}

// Add a user
func AddUser(c *gin.Context) {
	var user models.User
	c.BindJSON(&user)
	config.DB.Create(&user)
	c.JSON(200, user)
}
