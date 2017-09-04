package base

import (
	"github.com/gin-gonic/gin"
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
	//TODO add checking if same block is received - we take the last one
	if blockDiff <= 1 && getBlockChainSyncStatus() {
		log.Info("Saving new block: ", recv.Block.ID, " height:", recv.Block.Height, " transactions:", recv.Block.NumberOfTransactions, " peer:", c.ClientIP())
		err := ArkNodeDB.Save(&recv.Block)
		if err != nil {
			log.Error(err.Error())
		}
		c.JSON(200, gin.H{"success": true, "blockId": recv.Block.ID})
	} else if blockDiff > 1 && getBlockChainSyncStatus() {
		//look like there was a small network jump - resyncing blocks - here calling it not in thread...
		log.Info("Calling resync - lost ", blockDiff, " blocks!")
		go SyncBlockChain(recv.Block.Height)
		c.JSON(200, gin.H{"success": false, "message": "ECHAIN_NOT_SYNCED"})
	}
}
