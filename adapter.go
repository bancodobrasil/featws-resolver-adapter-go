package adapter

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/bancodobrasil/featws-resolver-adapter-go/routes"
	"github.com/bancodobrasil/featws-resolver-adapter-go/services"
	ginMonitor "github.com/bancodobrasil/gin-monitor"
	telemetry "github.com/bancodobrasil/gin-telemetry"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func init() {
	ex, err := os.Executable()
	if err != nil {
		log.Fatal(err)
	}
	exePath := filepath.Dir(ex)
	viper.AddConfigPath(exePath)
	viper.SetConfigType("env")
	viper.SetConfigName(".env")

	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)
	viper.AutomaticEnv()
	viper.SetDefault("RESOLVER_LOG_JSON", false)
	viper.SetDefault("RESOLVER_LOG_LEVEL", "error")
	viper.SetDefault("RESOLVER_SERVICE_NAME", "resolver-adapter-go")
	if err := viper.ReadInConfig(); err == nil {
		log.Infof("Using config file: %s", viper.ConfigFileUsed())
	}
}

// Config ...
type Config struct {
	Port string
}

// Run start the resolver server with resolverFunc
func Run(resolverFunc services.ResolverFunc, config Config) error {

	InitLogger()

	monitor, err := ginMonitor.New("v0.0.1-rc8", ginMonitor.DefaultErrorMessageKey, ginMonitor.DefaultBuckets)
	if err != nil {
		panic(err)
	}

	services.SetupResolver(resolverFunc)

	router := gin.New()

	// Register gin-monitor middleware
	router.Use(monitor.Prometheus())
	// Register metrics endpoint
	router.GET("/metrics", gin.WrapH(promhttp.Handler()))
	// Register gin-telemetry middleware
	router.Use(telemetry.Middleware(viper.GetString("RESOLVER_SERVICE_NAME")))

	routes.SetupRoutes(router)

	return router.Run(":" + config.Port)
}
