package routes

import (
	"github.com/bancodobrasil/featws-resolver-adapter-go/docs"
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

	externalHost := viper.GetString("EXTERNAL-HOST")

	docs.SwaggerInfo.Host = externalHost

	homeRouter(router.Group("/"))
	health.Router(router.Group("/health"))
	// setup swagger docs
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}

// APIRoutes sets up API routes with middleware for a Gin router in Go.
func APIRoutes(router *gin.Engine) {
	// inject middleware
	group := router.Group("/api")
	group.Use(telemetry.Middleware(viper.GetString("RESOLVER_SERVICE_NAME")))
	api.Router(group)
}
