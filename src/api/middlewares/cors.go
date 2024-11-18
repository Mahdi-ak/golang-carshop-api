package middlewares

import (
	"github.com/gin-gonic/gin"

	"github.com/Mahdi-ak/golang-carshop-api/src/config"
)

// Cors returns a middleware that handles CORS requests based on the provided config
func Cors(cfg *config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {

		// Set the allowed origins for CORS
		c.Writer.Header().Set("Access-Control-Allow-Origin", cfg.Cors.AllowOrigins)

		// Allow credentials (cookies, authorization headers, etc.)
		c.Header("Access-Control-Allow-Credentials", "true")

		// Specify the allowed headers for requests
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")

		// Define the allowed HTTP methods
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE,UPDATE")

		// Set the max age for preflight requests (in seconds)
		c.Header("Access-Control-Max-Age", "21600")

		// Set the default content type for responses
		c.Set("content-type", "application/json")

		// Handle preflight requests (OPTIONS method)
		if c.Request.Method == "OPTIONS" {
			// Abort and return status 204 (no content) for preflight requests
			c.AbortWithStatus(204)
			return
		}

		// Continue to the next middleware or handler
		c.Next()
	}
}
