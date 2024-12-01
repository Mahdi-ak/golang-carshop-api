package main

import (
	"github.com/Mahdi-ak/golang-carshop-api/src/api"
	"github.com/Mahdi-ak/golang-carshop-api/src/config"
	"github.com/Mahdi-ak/golang-carshop-api/src/data/cache"
	"github.com/Mahdi-ak/golang-carshop-api/src/data/db"
	"github.com/Mahdi-ak/golang-carshop-api/src/data/db/migrations"
	"github.com/Mahdi-ak/golang-carshop-api/src/pkg/logging"
)

func main() {
	cfg := config.GetConfig()
	logger := logging.NewLogger(cfg)

	// Initialize the redis
	err := cache.InitRedis(cfg)
	if err != nil {
		logger.Fatal(logging.Redis, logging.Startup, err.Error(), nil)
	}
	defer cache.CloseRedis()

	// Initialize the postgres
	err = db.InitDb(cfg)
	if err != nil {
		logger.Fatal(logging.Postgres, logging.Startup, err.Error(), nil)
	}
	defer db.CloseDb()
	migrations.Up_1()

	// Initialize the server
	api.InitServer(cfg)

}
