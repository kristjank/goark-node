package main

import "github.com/kristjank/goark-node/api"

func initializeRoutes() {
	logger.Println("Initializing routes")

	// TODO implement checkHeaders function to check the correct header data
	// indicating whether the request was from an authenticated peer? or not
	// Handle the index route
	//router.GET("/")

	// Group peer related routes together
	peerRoutes := router.Group("/peer")
	{
		peerRoutes.GET("/list", api.CheckNetworkHeaders(), api.GetTransactions)
		peerRoutes.GET("/status", api.CheckNetworkHeaders(), api.SendPeerStatus)

		peerRoutes.POST("/blocks", api.CheckNetworkHeaders(), api.ReceiveBlocks)
	}

	transactionRoutes := router.Group("/api/transactions")
	{
		// Handle the GET requests at /peer/list
		transactionRoutes.GET("/", api.CheckNetworkHeaders(), api.GetTransactions)
	}

	apiRoutes := router.Group("/api/blocks")
	{
	apiRoutes.GET("/getHeight", api.CheckNetworkHeaders(), api.GetHeight)	
	}
}
