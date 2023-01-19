package cache

import (
	"context"
	"encoding/json"
	"github.com/go-redis/redis/v8"
	"github.com/prashar32/rest-api-1/models"
	"time"
)

type RedisCache struct {
	host    string
	db      int
	expires time.Duration
}

func RedisInit(host string, db int, exp time.Duration) *RedisCache {
	return &RedisCache{
		host:    host,
		db:      db,
		expires: exp,
	}
}

func (cache *RedisCache) getClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     cache.host,
		Password: "eYVX7EwVmmxKPCDmwMtyKVge8oLd2t81",
		DB:       cache.db,
	})
}

func (cache *RedisCache) Set(key string, post []models.Shop) {
	client := cache.getClient()

	// serialize Post object to JSON
	json, err := json.Marshal(post)
	if err != nil {
		panic(err)
	}

	client.Set(context.Background(), key, json, cache.expires*time.Second)
}

func (cache *RedisCache) Get(key string) []models.Shop {
	client := cache.getClient()

	val, err := client.Get(context.Background(), key).Result()
	if err != nil {
		return nil
	}

	post := []models.Shop{}
	err = json.Unmarshal([]byte(val), &post)
	if err != nil {
		panic(err)
	}
	// fmt.Println("In Get")
	// fmt.Println(post)
	return post
}

//
//func (cache *RedisCache) Set(key string, post *models.Hotel) {
//	client := cache.getClient()
//
//	// serialize Post object to JSON
//	json, err := json.Marshal(post)
//	if err != nil {
//		panic(err)
//	}
//
//	client.Set(context.Background(), key, json, cache.expires*time.Second)
//}
//
//func (cache *RedisCache) Get(key string) *models.Hotel {
//	client := cache.getClient()
//
//	val, err := client.Get(context.Background(), key).Result()
//	if err != nil {
//		return nil
//	}
//
//	post := models.Hotel{}
//	err = json.Unmarshal([]byte(val), &post)
//	if err != nil {
//		panic(err)
//	}
//
//	return &post
//}
