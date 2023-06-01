package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

// VerifyAPIKeyMiddleware contains a string key used for verifying API keys.
//
// Property:
//   - {string} key: is a string that represents the API key that will be used to verify requests in the middleware.
type VerifyAPIKeyMiddleware struct {
	key string
}

var verifyAPIKeyMiddleware *VerifyAPIKeyMiddleware

// VerifyAPIKey returns a Gin middleware function that verifies an API key.
func VerifyAPIKey() gin.HandlerFunc {
	return verifyAPIKeyMiddleware.Run()
}

// NewVerifyAPIKeyMiddleware initializes a new instance of VerifyAPIKeyMiddleware with a key obtained from a
// configuration file.
func NewVerifyAPIKeyMiddleware() {
	verifyAPIKeyMiddleware = &VerifyAPIKeyMiddleware{
		key: viper.GetString("RESOLVER_API_KEY"),
	}
}

// Run is a method of the `VerifyAPIKeyMiddleware` struct that returns a `gin.HandlerFunc`. The
// `gin.HandlerFunc` is a function that takes a `gin.Context` as its parameter and returns nothing.
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

// `extractKeyFromHeader is a method of the `VerifyAPIKeyMiddleware` struct that extracts the API key from the `X-API-Key` header of the
// incoming request. It takes a `gin.Context` as its parameter and returns a string. If the `X-API-Key`
// header is missing, it responds with an error message and a status code of 401 (Unauthorized).
func (m *VerifyAPIKeyMiddleware) extractKeyFromHeader(c *gin.Context) string {
	authorizationHeader := c.Request.Header.Get("X-API-Key")
	if authorizationHeader == "" {
		respondWithError(c, 401, "Missing X-API-Key Header")
	}
	return authorizationHeader
}
