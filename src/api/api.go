package api

import (
	"fmt"

	"github.com/Mahdi-ak/golang-carshop-api/src/api/middlewares"
	"github.com/Mahdi-ak/golang-carshop-api/src/api/routers"
	"github.com/Mahdi-ak/golang-carshop-api/src/api/validations"
	"github.com/Mahdi-ak/golang-carshop-api/src/config"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func InitServer(cfg *config.Config) {

	// Initialize the server with configuration settings
	r := gin.New()
	RegisterValidations()
	r.Use(gin.Logger(), gin.Recovery())
	r.Use(middlewares.Cors(cfg))

	RegisterRoutes(r)
	// listen and serve on 0.0.0.0:5005
	r.Run(fmt.Sprintf(":%d", cfg.Server.Port))
}
func RegisterValidations() {
	val, ok := binding.Validator.Engine().(*validator.Validate)
	if ok {
		val.RegisterValidation("mobile", validations.IranianMobileNumberValidator, true)
		// val.RegisterValidation("password", validations.PasswordValidator, true)
	}
}
func RegisterRoutes(r *gin.Engine) {
	api := r.Group("/api")

	v1 := api.Group("/v1")
	{
		health := v1.Group("/health")
		test_router := v1.Group("/test")

		routers.Health(health)
		routers.TestRouter(test_router)

	}

	v2 := api.Group("/v2")
	{
		health := v2.Group("/health")
		routers.Health(health)
	}

}
