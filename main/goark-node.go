package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/kristjank/ark-go/core"
	"github.com/kristjank/goark-node/api"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var router *gin.Engine

//var log = logrus.New()

func init() {
	initLogger()
	loadConfig()
	initializeBoltClient()

	//api.InitDB(viper.Sub("db"))
}

func initLogger() {
	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.JSONFormatter{})

	//log.SetOutput(os.Stdout)

	// You could set this to any `io.Writer` such as a file
	file, err := os.OpenFile("log/goark-node.log", os.O_CREATE|os.O_WRONLY, 0666)
	if err == nil {
		log.SetOutput(file)
	} else {
		log.Info("Failed to log to file, using default stderr")
	}

	// Only log the warning severity or above.
	log.SetLevel(log.DebugLevel)
}

func loadConfig() {
	viper.SetConfigName("config.mainnet") // name of config file (without extension)
	viper.SetConfigType("json")
	viper.AddConfigPath("cfg") // path to look for the config file in

	err := viper.ReadInConfig() // Find and read the config file

	if err != nil {
		log.Fatal("No configuration file")
	}
}

func initializeBoltClient() {
	api.DBClient = &api.BoltClient{}
	api.DBClient.OpenBoltDb()
	api.DBClient.InitializeBucket()
	//transport.DBClient.Seed()
}

///////////////////////////
func main() {
	log.Info("---- GOArk Relay Node Starting ----")
	arkapi := core.NewArkClient(nil)
	arkapi = arkapi.SetActiveConfiguration(core.MAINNET)
	// Set the router as the default one provided by Gin
	router = gin.Default()

	// Initialize the routes
	initializeRoutes()

	// Start serving the application
	pNodeInfo := fmt.Sprintf("%s:%d", viper.GetString("address"), viper.GetInt("port"))
	router.Run(pNodeInfo)

}
