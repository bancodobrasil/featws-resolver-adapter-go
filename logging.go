package adapter

import (
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func InitLogger() {
	log.SetFormatter(&log.TextFormatter{FullTimestamp: true})
	if viper.GetBool("RESOLVER_LOG_JSON") {
		log.SetFormatter(&log.JSONFormatter{})
	} else {
		log.SetFormatter(&log.TextFormatter{TimestampFormat: "2006-01-02 15:04:05", FullTimestamp: true})
	}

	log.SetOutput(os.Stdout)
	level := viper.GetString("RESOLVER_LOG_LEVEL")
	if level == "" {
		level = "error"
	}
	l, err := log.ParseLevel(level)
	if err != nil {
		log.Fatal(err)
	}
	log.SetLevel(l)
}
