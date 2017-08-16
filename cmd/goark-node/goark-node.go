package main

import (
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/asdine/storm"
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

	// You could set this to any `io.Writer` such as a file
	file, err := os.OpenFile(viper.GetString("logFileName"), os.O_CREATE|os.O_WRONLY, 0666)
	if err == nil {
		log.SetOutput(io.MultiWriter(file, os.Stdout))
	} else {
		log.Error("Failed to log to file, using default stderr")
	}

	//TODO set log level according to cfg/settings
	//log.SetLevel(log.InfoLevel)
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
		log.Panic("Read configuration error", err.Error())
	}
}

func initializeDB() {
	var err error
	api.ArkNodeDB, err = storm.Open(viper.GetString("db.filename"))

	if err != nil {
		log.Panic("Storm DB init error", err.Error())
	}

	log.Info("Storm DB Opened at:", api.ArkNodeDB.Path)

	//defer api.ArkNodeDB.Close()
}

func initializeRoutes() {
	log.Info("Initializing routes")

	// Group peer related routes together
	peerRoutes := router.Group("/peer")
	peerRoutes.Use(api.CheckNetworkHeaders())
	{
		peerRoutes.GET("/list", api.SendPeerList)
		peerRoutes.GET("/status", api.SendPeerStatus)

		//TODO fix parallel semaphore for concurency exection of BS syncing
		//peerRoutes.POST("/blocks", api.CheckIfChainLoading(), api.ReceiveBlocks)
		//peerRoutes.POST("/transactions", api.CheckIfChainLoading(), api.ReceiveTransactions)
		peerRoutes.POST("/blocks", api.ReceiveBlocks)
		peerRoutes.POST("/transactions", api.ReceiveTransactions)
	}

	transactionRoutes := router.Group("/api/transactions")
	transactionRoutes.Use(api.CheckNetworkHeaders())
	{
		transactionRoutes.GET("", api.SendTransactions)
	}

	apiRoutes := router.Group("/api/blocks")
	apiRoutes.Use(api.CheckNetworkHeaders())
	{
		apiRoutes.GET("/getHeight", api.SendHeight)
	}

	autoconfigureRoutes := router.Group("/api/loader")
	{
		autoconfigureRoutes.GET("/autoconfigure", api.SendAutoConfigureParams)
	}

}

func initBlockChain() {
	api.SyncBlockChain()
}

///////////////////////////////////////////////////////////////////////////////////////
func main() {
	//init arkapi client - to get other peers and sync with blockchain
	api.ArkAPIClient = core.NewArkClient(nil)
	api.ArkAPIClient = api.ArkAPIClient.SetActiveConfiguration(core.MAINNET)

	//api.ArkApiClient = core.TestMethodNewArkClient(nil)

	//reading commandline args
	networkMode := flag.Bool("devnet", false, "Is devnet mode")
	flag.Parse()

	loadConfig(*networkMode)
	initLogger()
	initializeDB()
	//starting blockchain sync in a thread...
	//TODO needs testing
	log.Info("---- GOARK Relay Node Starting ----")
	go initBlockChain()

	log.Info("Flag arguments", flag.Args())
	if *networkMode {
		log.Info("DEVNET mode active")
		api.ArkAPIClient = api.ArkAPIClient.SetActiveConfiguration(core.DEVNET)
	}

	//GIN SERVER INIT
	// Set the router as the default one provided by Gin
	f, err := os.OpenFile("logs/server.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()

	gin.SetMode(gin.DebugMode)
	gin.DefaultWriter = io.MultiWriter(f)
	router = gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	//router = gin.Default()
	// Initialize the routes
	initializeRoutes()

	// Start serving the application
	pNodeInfo := fmt.Sprintf("%s:%d", viper.GetString("address"), viper.GetInt("port"))
	log.Info("Starting server, listening on:", pNodeInfo)
	router.Run(pNodeInfo)
}
