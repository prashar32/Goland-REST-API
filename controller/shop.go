package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/prashar32/rest-api-1/cache"
	"strconv"
	"time"

	//"github.com/prashar32/rest-api-1/cache"
	"github.com/prashar32/rest-api-1/config"
	"github.com/prashar32/rest-api-1/models"
	//"strconv"
	//"time"
)

var shopCache cache.RedisCache

// Cache Initialized
func CacheIntialize(host string, db int, exp time.Duration) {
	shopCache = *cache.RedisInit(host, db, exp)
}

type SHOP interface {
	GetShop()
	QueryShops()
	CreateShop()
}

// Get all shops
func GetShop(c *gin.Context) {
	shop := []*models.Shop{}
	result := config.DB.Find(&shop)
	if result.Error != nil {
		panic(result.Error)
	}

	c.JSON(200, shop)
}

// Get all shops corresponding to the particular category {like fashion}
func QueryShops(c *gin.Context) {
	var tofind string = "QueryShops" + c.Query("shopCategory")
	var shop []models.Shop = shopCache.Get(tofind)
	if shop != nil {
		fmt.Println("hello")
		c.JSON(200, shop)
		return
	}
	shopCategory, _ := strconv.Atoi(c.Query("shopCategory"))
	result := []models.Shop{}
	config.DB.Where("shopCategory = ?", shopCategory).Find(&result)
	shopCache.Set(tofind, result)
	fmt.Println("Query")
	c.JSON(200, result)
}

// Add a shop
func CreateShop(c *gin.Context) {
	var shop models.Shop
	c.BindJSON(&shop)
	config.DB.Create(&shop)
	c.JSON(200, shop)
}
