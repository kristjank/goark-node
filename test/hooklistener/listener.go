package main

import (
	"os"

	log "github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
	"github.com/kristjank/ark-go/core"
)

var arkclient = core.NewArkClient(nil)
var router *gin.Engine

func initLogger() {
	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
}

func initializeRoutes() {
	log.Info("Initializing routes")

	// Group peer related routes together
	peerRoutes := router.Group("/")
	peerRoutes.Use()
	{
		peerRoutes.GET("/testcall", testAPICall)
	}
}

func testAPICall(c *gin.Context) {
	//postTransaction()
	c.JSON(200, gin.H{"success": true, "msg": "hi "})
}

func main() {
	initLogger()
	arkclient = arkclient.SetActiveConfiguration(core.DEVNET)
	router = gin.Default()
	port := "45678"

	// Initialize the routes
	initializeRoutes()

	router.Run(":" + port)
}
