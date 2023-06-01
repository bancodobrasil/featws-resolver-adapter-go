package routes

import (
	"github.com/bancodobrasil/featws-resolver-adapter-go/controllers"
	"github.com/gin-gonic/gin"
)

// homeRouter sets up a GET route for the home page using the Gin web framework.
func homeRouter(router *gin.RouterGroup) {
	router.GET("/", controllers.HomeHandler)
}
