package routers

import (
	"github.com/Mahdi-ak/golang-carshop-api/src/api/handlers"
	"github.com/Mahdi-ak/golang-carshop-api/src/api/middlewares"
	"github.com/Mahdi-ak/golang-carshop-api/src/config"
	"github.com/gin-gonic/gin"
)

func User(router *gin.RouterGroup) {
	cfg := config.GetConfig()
	handlers := handlers.NewUserHandler(cfg)

	router.POST("/send-otp", middlewares.OtpLimiter(cfg), handlers.SendOtp)
	router.POST("/login-by-username", middlewares.OtpLimiter(cfg), handlers.LoginByUsername)
	router.POST("/login-by-mobile", middlewares.OtpLimiter(cfg), handlers.RegisterLoginByMobileNumber)
	router.POST("/register-by-username", middlewares.OtpLimiter(cfg), handlers.RegisterByUsername)

}
