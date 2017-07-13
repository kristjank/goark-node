package main

import (
	"github.com/kristjank/goark-node/api"
	log "github.com/sirupsen/logrus"
)

func initializeRoutes() {
	log.Info("Initializing routes")

	// TODO implement checkHeaders function to check the correct header data
	// indicating whether the request was from an authenticated peer? or not
	// Handle the index route
	//router.GET("/")

	// Group peer related routes together
	peerRoutes := router.Group("/peer")
	{
		peerRoutes.GET("/list", api.GetTransactions)
		peerRoutes.GET("/status", api.SendPeerStatus)

		peerRoutes.POST("/blocks", api.ReceiveBlocks)
		peerRoutes.POST("/transactions", api.ReceiveTransactions)
	}

	transactionRoutes := router.Group("/api/transactions")
	{
		// Handle the GET requests at /peer/list
		transactionRoutes.GET("", api.GetTransactions)
	}

	apiRoutes := router.Group("/api/blocks")
	{
		apiRoutes.GET("/getHeight", api.GetHeight)
	}

	autoconfigureRoutes := router.Group("/api/loader")
	{
		autoconfigureRoutes.GET("/autoconfigure", api.GetAutoConfigureParams)
	}
}
