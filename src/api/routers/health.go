package routers

import (
	"github.com/Mahdi-ak/golang-carshop-api/src/api/handlers"
	"github.com/gin-gonic/gin"
)

// Health registers the health-related API endpoints with the provided router group.
func Health(r *gin.RouterGroup) {
	handlers := handlers.NewHealthHandler()

	r.GET("/", handlers.Health)
	r.POST("/", handlers.HealthPost)
	r.POST("/:id", handlers.HealthPostById)

}
