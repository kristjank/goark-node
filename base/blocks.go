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

	//log.Info("New block received - id: ", recv.Block.ID, " height:", recv.Block.Height, " transactions:", len(recv.Block.Transactions), " peer:", c.Request.RemoteAddr)
	lastBlock, _ := getLastBlock()

	if recv.Block.Height-lastBlock.Height == 1 {
		log.Info("Saving new block: ", recv.Block.ID, " height:", recv.Block.Height, " transactions:", len(recv.Block.Transactions), " peer:", c.ClientIP())
		err := ArkNodeDB.Save(&recv.Block)
		if err != nil {
			log.Error(err.Error())
		}
		c.JSON(200, gin.H{"success": true, "blockId": recv.Block.ID})
	} else {
		c.JSON(200, gin.H{"success": false, "error": "Chain not at the same height. Unable to receive"})
	}
}
