package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

type VerifyAPIKeyMiddleware struct {
	key string
}

var verifyAPIKeyMiddleware *VerifyAPIKeyMiddleware

// Middleware function to verify the JWT token
func VerifyAPIKey() gin.HandlerFunc {
	return verifyAPIKeyMiddleware.Run()
}

func NewVerifyAPIKeyMiddleware() {
	verifyAPIKeyMiddleware = &VerifyAPIKeyMiddleware{
		key: viper.GetString("RESOLVER_API_KEY"),
	}
}

func (m *VerifyAPIKeyMiddleware) Run() gin.HandlerFunc {
	return func(c *gin.Context) {
		key := m.extractKeyFromHeader(c)

		if key != m.key {
			respondWithError(c, 401, "Unauthorized")
		}

		c.Next()
	}
}

func (m *VerifyAPIKeyMiddleware) extractKeyFromHeader(c *gin.Context) string {
	authorizationHeader := c.Request.Header.Get("X-API-Key")
	if authorizationHeader == "" {
		respondWithError(c, 401, "Missing X-API-Key Header")
	}
	return authorizationHeader
}
