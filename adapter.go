package adapter

import (
	"os"

	"github.com/bancodobrasil/featws-resolver-adapter-go/routes"
	"github.com/bancodobrasil/featws-resolver-adapter-go/services"
	ginMonitor "github.com/bancodobrasil/gin-monitor"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	log "github.com/sirupsen/logrus"
	ginlogrus "github.com/toorop/gin-logrus"
)

func setupLog() {
	log.SetFormatter(&log.TextFormatter{FullTimestamp: true})

	log.SetOutput(os.Stdout)

	log.SetLevel(log.DebugLevel)
}

// Config ...
type Config struct {
	Port string
}

// Run start the resolver server with resolverFunc
func Run(resolverFunc services.ResolverFunc, config Config) error {

	setupLog()

	monitor, err := ginMonitor.New("v0.0.1-rc7", ginMonitor.DefaultErrorMessageKey, ginMonitor.DefaultBuckets)
	if err != nil {
		panic(err)
	}

	gin.DefaultWriter = log.StandardLogger().WriterLevel(log.DebugLevel)
	gin.DefaultErrorWriter = log.StandardLogger().WriterLevel(log.ErrorLevel)

	services.SetupResolver(resolverFunc)

	router := gin.New()
	// Register ginLogrus log format to gin
	router.Use(ginlogrus.Logger(log.StandardLogger()), gin.Recovery())

	// Register gin-monitor middleware
	router.Use(monitor.Prometheus())
	// Register metrics endpoint
	router.GET("/metrics", gin.WrapH(promhttp.Handler()))

	routes.SetupRoutes(router)

	return router.Run(":" + config.Port)
}
