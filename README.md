# Shopping Mall API

## Introduction 
It is a simple implementation of basic of Golang server based functionality.

## Setup
First run `docker-compose.yml` to setup two container one for `Postgres` and another for `redis-Cache`

### Docker Compose Build Command
> docker compose up --build

Then run `go run main.go`

Postgres port: 5432 
Golang Project port: 8080

## API Calls

### Get All Shops
http://localhost:8080/api/shop/getAllShops

### Find Shops corresponding to the Shop Category
http://localhost:8080/api/shop/queryShops/"shopCategory"

### Add a Shop
http://localhost:8080/api/shop/addShop

### Get All Users
http://localhost:8080/api/user/getInfo

### Find Users corresponding to the paricular Shop Visited
http://localhost:8080/api/user/fetchUser/"shopVisited"

### Add a User
http://localhost:8080/api/shop/addUser

### This is implemented using goroutines and will give the data of users who viisted the shop with "id":id  and also the infomation of that shop
http://localhost:8080/api/shop/user/getInfo/"id"

## Used `Gorm` library of Golang, `gin-gonic` as a HTTP framework, `redis-cache` as a cache, `golang-migrate` as a Migration Library for database and `goroutines` as a function and `interface` as a type. 
