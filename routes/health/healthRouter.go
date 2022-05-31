package health

import (
	"github.com/bancodobrasil/featws-resolver-adapter-go/controllers"
	"github.com/gin-gonic/gin"
)

// Router ...
func Router(router *gin.RouterGroup) {
	router.GET("/live", controllers.HealthLiveHandler())
	router.GET("/ready", controllers.HealthLiveHandler())
}
