package base

import (
	"sync"

	"github.com/asdine/storm"
	"github.com/kristjank/ark-go/core"
)

//ArkNodeDB interface is setup in goark-node.go.
//interface is visible in the whole package api
var ArkNodeDB *storm.DB

//ArkAPIClient - using to talk with other nodes with the ARK-GO Client API
var ArkAPIClient *core.ArkClient

//IsBlockchainSynced - Global locker (semaphore) - to ignore writes and receive blocks if blocks arent at the height
var IsBlockchainSynced bool

//IsBlockSaving - Global locker (semaphore) - to ignore writes of secondly received blocks if the same...
//Important for eventhook api
var IsBlockSaving bool

//ChainSyncChannel - to signalend execution of chain sync
//make() called somewhere else (buffered)
var ChainSyncChannel chan int

//SyncMutex to be wrapped around sync threads
var SyncMutex = &sync.RWMutex{}

//SaveBlockMutex to be wrapped around sync threads
var SaveBlockMutex = &sync.RWMutex{}

func init() {
	ChainSyncChannel = make(chan int, 1)
	//set
	setBlockChainSyncStatus(true)
}
