package main

import (
	"github.com/Mahdi-ak/golang-carshop-api/src/api"
	"github.com/Mahdi-ak/golang-carshop-api/src/config"
	"github.com/Mahdi-ak/golang-carshop-api/src/data/cache"
)

func main() {

	// Initialize the redis
	cfg := config.GetConfig()
	cache.InitRedis(cfg)
	defer cache.CloseRedis()

	// Initialize the server
	api.InitServer(cfg)

}
