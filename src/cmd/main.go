package main

import (
	"log"

	"github.com/Mahdi-ak/golang-carshop-api/src/api"
	"github.com/Mahdi-ak/golang-carshop-api/src/config"
	"github.com/Mahdi-ak/golang-carshop-api/src/data/cache"
	"github.com/Mahdi-ak/golang-carshop-api/src/data/db"
)

func main() {

	// Initialize the redis
	cfg := config.GetConfig()
	err := cache.InitRedis(cfg)
	if err != nil {
		log.Fatal(err)
	}
	defer cache.CloseRedis()

	// Initialize the postgres
	err = db.InitDb(cfg)
	if err != nil {
		log.Fatal(err)
	}
	defer db.CloseDb()

	// Initialize the server
	api.InitServer(cfg)

}
