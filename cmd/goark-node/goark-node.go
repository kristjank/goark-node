package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/kristjank/ark-go/core"
	"github.com/kristjank/goark-node/api"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var router *gin.Engine

func initLogger() {
	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.JSONFormatter{})

	//log.SetOutput(os.Stdout)

	// You could set this to any `io.Writer` such as a file
	file, err := os.OpenFile(viper.GetString("logFileName"), os.O_CREATE|os.O_WRONLY, 0666)
	if err == nil {
		log.SetOutput(file)
	} else {
		log.Info("Failed to log to file, using default stderr")
	}

	//TODO set log level according to cfg/settings
	log.SetLevel(log.DebugLevel)
}

func loadConfig(isDEVNET bool) {
	viper.SetConfigType("json")
	viper.AddConfigPath("cfg") // path to look for the config file in

	if isDEVNET {
		viper.SetConfigName("config.devnet") // name of config file (without extension)
	} else {
		viper.SetConfigName("config.mainnet") // name of config file (without extension)
	}

	err := viper.ReadInConfig() // Find and read the config file

	if err != nil {
		fmt.Println(err.Error())
		panic("No configuration file")
	}
}

func initializeBoltClient() {
	api.DBClient = &api.BoltClient{}
	api.DBClient.OpenBoltDb()
	api.DBClient.InitializeBucket()
}

///////////////////////////////////////////////////////////////////////////////////////
func main() {
	log.Info("---- GOARK Relay Node Starting ----")
	arkapi := core.NewArkClient(nil)
	arkapi = arkapi.SetActiveConfiguration(core.MAINNET)

	//reading commandline args
	//SILENT MODE CHECKING AND AUTOMATION RUNNING
	networkMode := flag.Bool("devnet", false, "Is devnet mode")
	flag.Parse()

	loadConfig(*networkMode)
	initLogger()
	initializeBoltClient()

	log.Println(flag.Args())
	if *networkMode {
		log.Info("DEVNET mode active")
		arkapi = arkapi.SetActiveConfiguration(core.DEVNET)
	}

	// Set the router as the default one provided by Gin
	router = gin.Default()
	// Initialize the routes
	initializeRoutes()
	// Start serving the application
	pNodeInfo := fmt.Sprintf("%s:%d", viper.GetString("address"), viper.GetInt("port"))
	router.Run(pNodeInfo)

}
