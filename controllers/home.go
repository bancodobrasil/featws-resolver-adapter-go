package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// func HomeHandler() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		c.String(http.StatusOK, "FeatWS Resolver Works!!!")
// 	}

// }

// HomeHandler function returns a string indicating that the FeatWS Resolver is working.
func HomeHandler(c *gin.Context) {
	c.String(http.StatusOK, "FeatWS Resolver Works!!!")
}
