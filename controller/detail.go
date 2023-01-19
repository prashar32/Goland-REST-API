package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/prashar32/rest-api-1/config"
	"github.com/prashar32/rest-api-1/models"
	"sync"
)

var wg sync.WaitGroup

func GetDetails(c *gin.Context) {
	url := c.FullPath()
	fmt.Println(url)
	var shop models.Shop
	var user models.User
	wg.Add(2)
	go func(shop models.Shop) {
		config.DB.Where("shopId", c.Param("shopId")).Find(&shop)
		c.JSON(200, &shop)
		fmt.Println("I am at goroutine_1")
		wg.Done()
	}(shop)

	go func(user models.User) {
		config.DB.Where("shopVisited", c.Param("shopVisited")).Find(&user)
		c.JSON(200, &user)
		fmt.Println("I am at goroutine_2")
		wg.Done()
	}(user)
	wg.Wait()
	fmt.Println("GoRoutines done")
}
