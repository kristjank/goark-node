package base

import (
	"github.com/gin-gonic/gin"
	"github.com/kristjank/ark-go/core"
	"github.com/kristjank/goark-node/base/model"
	log "github.com/sirupsen/logrus"
)

//ReceiveBlocks from blockchain
func ReceiveBlocks(c *gin.Context) {
	var recv model.BlockReceiveStruct
	err := c.BindJSON(&recv)
	if err != nil {
		log.Error(err.Error())
	}

	lastBlock, _ := getLastBlock()
	blockDiff := recv.Block.Height - lastBlock.Height
	if blockDiff == 1 && getBlockChainSyncStatus() {
		//Mutex to wait - for saving 2 complete
		//Then common blocks can continue
		if !GetSaveBlockMutex() {
			SetSaveBlockMutex(true)
			log.Info("Saving new block: ", recv.Block.ID, " height:", recv.Block.Height, " transactions:", recv.Block.NumberOfTransactions, " peer:", c.ClientIP())
			err := ArkNodeDB.Save(&recv.Block)
			if err != nil {
				log.Error(err.Error())
			}
			c.JSON(200, gin.H{"success": true, "blockId": recv.Block.ID})
			multiBroadCastBlock(recv)
			SetSaveBlockMutex(false)
		} else {
			log.Debug("Saving process 2 DB active")
		}
	} else if blockDiff > 1 && getBlockChainSyncStatus() {
		//look like there was a small network jump - resyncing blocks - here calling it not in thread...
		log.Info("Calling resync - lost ", blockDiff, " blocks!")
		go SyncBlockChain(recv.Block.Height)
		c.JSON(200, gin.H{"success": false, "message": "ECHAIN_NOT_SYNCED"})
	} else if blockDiff == 0 && getBlockChainSyncStatus() {
		if checkCommonBlock(recv.Block) {
			log.Info("Found common block: ", recv.Block.ID, " height:", recv.Block.Height, " with peer:", c.ClientIP())
			c.JSON(200, gin.H{"success": true, "blockId": recv.Block.ID})
		} else {
			log.Info("Received different block than saved one: ", recv.Block.ID, " height:", recv.Block.Height, " with peer:", c.ClientIP())
			c.JSON(200, gin.H{"success": false, "blockId": recv.Block.ID})
		}
	}
}

func multiBroadCastBlock(block model.BlockReceiveStruct) {
	numberOfPeers2MultiBroadCastTo := 10
	if numberOfPeers2MultiBroadCastTo > 15 {
		numberOfPeers2MultiBroadCastTo = 15
		log.Warn("Max broadcast number too high - set by user, reseting to value 15")
	}
	log.Info("Starting multibroadcast/multithreaded parallel payout to ", numberOfPeers2MultiBroadCastTo, " number of peers")
	peers := ArkAPIClient.GetRandomXPeers(numberOfPeers2MultiBroadCastTo)
	for i := 0; i < numberOfPeers2MultiBroadCastTo; i++ {
		//treaded function
		go func(blockPayload model.BlockReceiveStruct, peer core.Peer) {
			//defer wg.Done()
			arkTmpClient := core.NewArkClientFromPeer(peer)
			res, _, _ := arkTmpClient.PostBlock(blockPayload)
			//TODO - maybe check for success of sending
			if res.Success {
				log.Debug("Block ", blockPayload.Block.ID, " retransmited to ", peer)
			}
		}(block, peers[i])
	}
}

func checkCommonBlock(block model.Block) bool {
	dbBlock, err := getBlockByID(block.ID)
	if err != nil {
		log.Error("Error check common block", err.Error())
		return false
	}
	if compareBlocks(dbBlock, block) {
		return true
	}
	return false
}

func compareBlocks(a, b model.Block) bool {
	if &a == &b {
		return true
	}
	if a.Height != b.Height || a.Timestamp != b.Timestamp {
		return false
	}
	if len(a.Transactions) != len(b.Transactions) || len(a.Transactions) != len(b.Transactions) {
		return false
	}
	return true
}
