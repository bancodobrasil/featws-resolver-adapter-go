package middlewares

import "github.com/gin-gonic/gin"

// Middleware defines a interface and initializes a middleware for verifying API keys.
//
// Property:
//   - Run: is a method that belongs to the Middleware interface. It defines the behavior of the middleware when it is executed. The implementation of this method will vary depending on the specific middleware being used.
type Middleware interface {
	Run()
}

// InitializeMiddlewares a middleware for verifying API keys.
func InitializeMiddlewares() {
	NewVerifyAPIKeyMiddleware()
}

// Helper function to abort the request with an error status code and message
func respondWithError(c *gin.Context, code int, message interface{}) {
	c.AbortWithStatusJSON(code, gin.H{"error": message})
}
