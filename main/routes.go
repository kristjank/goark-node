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
		// Handle the GET requests at /peer/list
		peerRoutes.GET("/list/", api.CheckNetworkHeaders(), api.GetTransactions)
		// Handle the GET requests at /peer/list
		peerRoutes.POST("/blocks/", api.CheckNetworkHeaders(), api.ReceiveBlocks)
	}

	transactionRoutes := router.Group("/api/transactions/")
	{
		// Handle the GET requests at /peer/list
		transactionRoutes.GET("/", api.CheckNetworkHeaders(), api.GetTransactions)
	}
}
