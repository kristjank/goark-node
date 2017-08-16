package base

import (
	"github.com/asdine/storm"
	"github.com/kristjank/ark-go/core"
)

//ArkNodeDB interface is setup in goark-node.go.
//interface is visible in the whole package api
var ArkNodeDB *storm.DB

//ArkAPIClient - using to talk with other nodes with the ARK-GO Client API
var ArkAPIClient *core.ArkClient

//IsBlockchainSynced - Global locker (semaphore) - to ignore writes and receive blocks if blocks arent at the height
var IsBlockchainSynced *bool
