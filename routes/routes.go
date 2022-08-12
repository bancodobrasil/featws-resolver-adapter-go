package routes

import (
	"github.com/bancodobrasil/featws-resolver-adapter-go/routes/api"
	"github.com/bancodobrasil/featws-resolver-adapter-go/routes/health"
	telemetry "github.com/bancodobrasil/gin-telemetry"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

// SetupRoutes define all routes
func SetupRoutes(router *gin.Engine) {

	homeRouter(router.Group("/"))
	health.Router(router.Group("/health"))
	// setup swagger docs
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}

func APIRoutes(router *gin.Engine) {
	// inject middleware
	group := router.Group("/api")
	group.Use(telemetry.Middleware(viper.GetString("RESOLVER_SERVICE_NAME")))
	api.Router(group)
}
