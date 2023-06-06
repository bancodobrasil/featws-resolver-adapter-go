package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// HealthLiveHandler returns a Gin handler function that responds with a string indicating that the
// application is live.
func HealthLiveHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.String(http.StatusOK, "Application is live!!!")
	}

}
