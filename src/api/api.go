package api

import (
	"fmt"

	"github.com/Mahdi-ak/golang-carshop-api/src/api/middlewares"
	"github.com/Mahdi-ak/golang-carshop-api/src/api/routers"
	"github.com/Mahdi-ak/golang-carshop-api/src/api/validations"
	"github.com/Mahdi-ak/golang-carshop-api/src/config"
	"github.com/Mahdi-ak/golang-carshop-api/src/docs"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitServer(cfg *config.Config) {

	// Initialize the server with configuration settings
	r := gin.New()
	RegisterValidations()
	r.Use(gin.Logger(), gin.Recovery())
	r.Use(middlewares.Cors(cfg), middlewares.LimitByRequest())

	RegisterRoutes(r)
	RegisterSwagger(r, cfg)
	// listen and serve on 0.0.0.0:5005
	r.Run(fmt.Sprintf(":%s", cfg.Server.Port))
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

func RegisterSwagger(r *gin.Engine, cfg *config.Config) {

	docs.SwaggerInfo.Title = "golang web api "
	docs.SwaggerInfo.Description = "golang web api "
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.BasePath = "/api"
	docs.SwaggerInfo.Host = fmt.Sprintf("localhost:%s", cfg.Server.Port)
	docs.SwaggerInfo.Schemes = []string{"http"}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

}
