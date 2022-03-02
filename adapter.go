package adapter

import (
	"os"

	"github.com/bancodobrasil/featws-resolver-adapter-go/config"
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

var Config = config.Config{}

func Run(resolverFunc services.ResolverFunc) {

	setupLog()

	err := config.LoadConfig(&Config)
	if err != nil {
		log.Fatalf("Não foi possível carregar as configurações: %s\n", err)
	}

	services.SetupResolver(resolverFunc)

	router := gin.New()

	routes.SetupRoutes(router)

	port := Config.Port

	router.Run(":" + port)
}
