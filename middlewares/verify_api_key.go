package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

// VerifyAPIKeyMiddleware ...
type VerifyAPIKeyMiddleware struct {
	key string
}

var verifyAPIKeyMiddleware *VerifyAPIKeyMiddleware

// VerifyAPIKey ...
func VerifyAPIKey() gin.HandlerFunc {
	return verifyAPIKeyMiddleware.Run()
}

// NewVerifyAPIKeyMiddleware ...
func NewVerifyAPIKeyMiddleware() {
	verifyAPIKeyMiddleware = &VerifyAPIKeyMiddleware{
		key: viper.GetString("RESOLVER_API_KEY"),
	}
}

// Run ...
func (m *VerifyAPIKeyMiddleware) Run() gin.HandlerFunc {
	return func(c *gin.Context) {
		if m.key == "" {
			c.Next()
			return
		}

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
