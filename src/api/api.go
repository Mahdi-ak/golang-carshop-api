package api

import (
	"fmt"

	"github.com/Mahdi-ak/golang-carshop-api/src/api/routers"
	"github.com/Mahdi-ak/golang-carshop-api/src/config"
	"github.com/gin-gonic/gin"
)

func InitServer() {
	// Initialize the server with configuration settings
	cfg := config.GetConfig()

	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())

	v1 := r.Group("/api/v1")
	{
		health := v1.Group("/health")
		routers.Health(health)

	}
	// listen and serve on 0.0.0.0:5005
	r.Run(fmt.Sprintf(":%d", cfg.Server.Port))
}
