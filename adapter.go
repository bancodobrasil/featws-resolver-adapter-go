package adapter

import (
	"os"
	"path/filepath"
	"strings"

	// Specificate the docs package
	_ "github.com/bancodobrasil/featws-resolver-adapter-go/docs"
	"github.com/bancodobrasil/featws-resolver-adapter-go/middlewares"
	"github.com/bancodobrasil/featws-resolver-adapter-go/routes"
	"github.com/bancodobrasil/featws-resolver-adapter-go/services"
	ginMonitor "github.com/bancodobrasil/gin-monitor"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	ginlogrus "github.com/toorop/gin-logrus"
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
	viper.SetDefault("EXTERNAL-HOST", "localhost:7000")
	viper.SetDefault("RESOLVER_LOG_JSON", false)
	viper.SetDefault("RESOLVER_LOG_LEVEL", "error")
	viper.SetDefault("RESOLVER_SERVICE_NAME", "resolver-adapter-go")
	viper.SetDefault("RESOLVER_API_KEY", "")
	if err := viper.ReadInConfig(); err == nil {
		log.Infof("Using config file: %s", viper.ConfigFileUsed())
	}
}

// Config ...
type Config struct {
	Port string
}

// @title FeatWS Resolver Adapter

// @version 1.0
// @Description
// @Description	O Resolver Adapter Project é uma biblioteca que fornece um conjunto de resolvers para serem utilizados por outros projetos. Resolvers são um conceito fundamental no desenvolvimento de software que são responsáveis por determinar o valor de um campo em um esquema GraphQL.
// @Description
// @Description	No contexto do GraphQL, um resolver é uma função que resolve um campo GraphQL, buscando os dados de uma fonte de dados, como um banco de dados ou uma API. O Resolver Adapter Project fornece um conjunto de resolvers pré-construídos que podem ser usados por outros projetos para lidar com cenários comuns de busca de dados.
// @Description
// @Description	Por exemplo, se você tem um esquema GraphQL que inclui um campo para buscar os dados de um usuário. O Resolver Adapter Project pode ser usado para buscar os dados do banco de dados interno do BB sobre um cliente especifico do banco, podendo retornar diversos parâmetros do cliente, como os seguintes:
// @Description
// @Description	- account (conta)
// @Description	- accountType (tipo de conta)
// @Description	- age (idade)
// @Description	- agenciaDet (tipo de conta)
// @Description	- branch (agencia)
// @Description	- branchState (Estado da agencia)
// @Description	- customerBase (agencia)
// @Description	- dataNascimento (data de nascimento)
// @Description	- employeeDependency (Dependecia do empregado do banco - só trará um retorno se a pessoa for funci do banco)
// @Description	- employeeKey (se é empregado do banco)
// @Description	- enterpriseKey (chave empresarial)
// @Description	- gender (sexo)
// @Description	- holder (titularidade da conta)
// @Description	- holderState (estado do titularidade da conta)
// @Description	- mci
// @Description	- mcipj
// @Description	- wallet (carteira)
// @Description
// @Description No geral, o Resolver Adapter Project é uma biblioteca útil que pode simplificar o desenvolvimento de APIs GraphQL fornecendo resolvers pré-construídos que podem ser facilmente integrados em outros projetos.
// @Description
// @Description Antes de realizar as requisições no Swagger, é necessário autorizar o acesso clicando no botão **Authorize**, ao lado, e inserindo a senha correspondente. Após inserir o campo **value** e clicar no botão **Authorize**, o Swagger estará disponível para ser utilizado.

// @termsOfService http://swagger.io/terms/

// @contact.name API Support

// @contact.url http://www.swagger.io/support

// @contact.email support@swagger.io

// @license.name Apache 2.0

// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:7000

// @BasePath /api/v1

// @securityDefinitions.apikey Authentication Api Key
// @in header
// @name X-API-Key

// @x-extension-openapi {"example": "value on a json format"}

// Run ...
func Run(resolverFunc services.ResolverFunc, config Config) error {

	InitLogger()

	monitor, err := ginMonitor.New("v0.0.1-rc8", ginMonitor.DefaultErrorMessageKey, ginMonitor.DefaultBuckets)
	if err != nil {
		panic(err)
	}

	gin.DefaultWriter = log.StandardLogger().WriterLevel(log.DebugLevel)
	gin.DefaultErrorWriter = log.StandardLogger().WriterLevel(log.ErrorLevel)

	services.SetupResolver(resolverFunc)

	router := gin.New()

	middlewares.InitializeMiddlewares()

	// Register ginLogrus log format to gin
	router.Use(ginlogrus.Logger(log.StandardLogger()), gin.Recovery())

	// Register gin-monitor middleware
	router.Use(monitor.Prometheus())
	// Register metrics endpoint
	router.GET("/metrics", gin.WrapH(promhttp.Handler()))
	routes.SetupRoutes(router)
	routes.APIRoutes(router)

	return router.Run(":" + config.Port)
}
