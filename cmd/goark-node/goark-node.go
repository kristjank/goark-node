package main

import (
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/asdine/storm"
	"github.com/gin-gonic/gin"
	"github.com/kristjank/ark-go/core"
	"github.com/kristjank/goark-node/base"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var router *gin.Engine
var GOARKNodeVersion string

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

	//redirecting stdOut of client ArkApiClient to logger
	//base.ArkAPIClient.Std = file
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
	base.ArkNodeDB, err = storm.Open(viper.GetString("db.filename"))

	if err != nil {
		log.Panic("Storm DB init error", err.Error())
	}

	log.Info("Storm DB Opened at:", base.ArkNodeDB.Path)

	//defer api.ArkNodeDB.Close()
}

func initializeRoutes() {
	log.Info("Initializing routes")

	// Group peer related routes together
	peerRoutes := router.Group("/peer")
	peerRoutes.Use(base.CheckNetworkHeaders())
	{
		peerRoutes.GET("/list", base.SendPeerList)
		peerRoutes.GET("/status", base.SendPeerStatus)

		peerRoutes.POST("/blocks", base.CheckIfChainLoading(), base.ReceiveBlocks)
		peerRoutes.POST("/transactions", base.CheckIfChainLoading(), base.ReceiveTransactions)
		peerRoutes.GET("/blocks/common", base.CheckIfChainLoading(), base.CheckCommonBlocks)
	}

	transactionRoutes := router.Group("/api/transactions")
	transactionRoutes.Use(base.CheckNetworkHeaders())
	{
		transactionRoutes.GET("", base.SendTransactions)
	}

	apiRoutes := router.Group("/api/blocks")
	apiRoutes.Use(base.CheckNetworkHeaders())
	{
		apiRoutes.GET("/getHeight", base.SendHeight)
	}

	autoconfigureRoutes := router.Group("/api/loader")
	{
		autoconfigureRoutes.GET("/autoconfigure", base.SendAutoConfigureParams)
	}

}

///////////////////////////////////////////////////////////////////////////////////////
func main() {
	//sending GOARK-NODE Version that we are working with payments
	//setting the version
	GOARKNodeVersion = "v0.2.0"
	//init arkapi client - to get other peers and sync with blockchain
	base.ArkAPIClient = core.NewArkClient(nil)

	//reading commandline args
	networkMode := flag.Bool("devnet", false, "Is devnet mode")
	flag.Parse()

	loadConfig(*networkMode)
	initLogger()
	initializeDB()
	//starting blockchain sync in a thread...
	//TODO needs testing
	//go base.SyncBlockChain(base.ArkAPIClient.GetActivePeer().Height)

	log.Info("---- GOARK Relay Node Starting ----")

	log.Info("Flag arguments", flag.Args())
	if *networkMode {
		log.Info("DEVNET mode active")
		base.ArkAPIClient = base.ArkAPIClient.SetActiveConfiguration(core.DEVNET)
	}

	//GIN SERVER INIT
	// Set the router as the default one provided by Gin
	f, err := os.OpenFile("logs/server.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()

	gin.SetMode(gin.DebugMode)
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
	//gin.DefaultWriter = io.MultiWriter(f)

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
