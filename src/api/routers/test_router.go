package routers

import (
	"github.com/Mahdi-ak/golang-carshop-api/src/api/handlers"
	"github.com/Mahdi-ak/golang-carshop-api/src/api/middlewares"
	"github.com/gin-gonic/gin"
)

func TestRouter(r *gin.RouterGroup) {
	h := handlers.NewTestHandler()

	r.GET("/", middlewares.TestMiddleware(), h.Test)
	r.GET("/users", h.Users)
	r.GET("/user/:id", h.UserById)
	r.GET("/user/get-user-by-username/:username", h.UserByUsername)
	r.GET("/user/:id/accounts", h.Accounts)

	r.POST("/add-user", h.AddUser)

	r.POST("/binder/header1", h.HeaderBinder1)
	r.POST("/binder/header2", h.HeaderBinder2)

}
