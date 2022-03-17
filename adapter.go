package adapter

import (
	"os"

	"github.com/bancodobrasil/featws-resolver-adapter-go/routes"
	"github.com/bancodobrasil/featws-resolver-adapter-go/services"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
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

	services.SetupResolver(resolverFunc)

	router := gin.New()

	routes.SetupRoutes(router)

	return router.Run(":" + config.Port)
}
