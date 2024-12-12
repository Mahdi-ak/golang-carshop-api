package routers

import (
	"github.com/Mahdi-ak/golang-carshop-api/src/api/handlers"
	"github.com/Mahdi-ak/golang-carshop-api/src/config"
	"github.com/gin-gonic/gin"
)

func User(router *gin.RouterGroup) {
	cfg := config.GetConfig()
	handlers := handlers.NewUserHandler(cfg)

	router.POST("/send-otp", handlers.SendOtp)
}
