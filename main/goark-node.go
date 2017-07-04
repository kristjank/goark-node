package main

import (
	"fmt"
	"log"
	"os"

	"github.com/kristjank/ark-go/core"
	"github.com/kristjank/goark-node/api"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	"gopkg.in/gin-gonic/gin.v1"
)

var errorlog *os.File
var logger *log.Logger
var router *gin.Engine

var env *ArkEnv

//ArkEnv struct
type ArkEnv struct {
	logger *log.Logger
}

func init() {
	initLogger()
	loadConfig()
	initializeBoltClient()

	//api.InitDB(viper.Sub("db"))
}

func initLogger() {
	errorlog, err := os.OpenFile("log/goark-node.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Printf("error opening file: %v", err)
		os.Exit(1)
	}

	logger = log.New(errorlog, "ark-go: ", log.Lshortfile|log.LstdFlags)
	env = &ArkEnv{logger: logger}
}

func loadConfig() {
	viper.SetConfigName("config.mainnet") // name of config file (without extension)
	viper.SetConfigType("json")
	viper.AddConfigPath("cfg") // path to look for the config file in

	err := viper.ReadInConfig() // Find and read the config file

	if err != nil {
		logger.Panic("No configuration file")
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
	logger.Println("GOArk Relay Node Starting")
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
