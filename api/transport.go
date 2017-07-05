package api

import (
	"github.com/gin-gonic/gin"
	"github.com/kristjank/goark-node/api/model"
	log "github.com/sirupsen/logrus"
)

var DBClient IBoltClient

//GetTransactions Returns a list of peers to client call. Response is in JSON
func GetTransactions(c *gin.Context) {
	/*res, err := QueryTransactions()
	if err != nil {
		log.Println(err.Error())
	}

	var a core.TransactionResponse

	a.Success = true
	//logger.pr
	//a.&Transactions = res

	c.JSON(200, gin.H{
		"message": "pong",
	})

	c.JSON(200, a)*/
}

//ReceiveBlocks from blockchain
func ReceiveBlocks(c *gin.Context) {
	log.Debug("Received blocks")
	//x, _ := ioutil.ReadAll(c.Request.Body)
	var recv model.BlockReceiveStruct
	c.BindJSON(&recv)

	DBClient.SaveBlock(recv.Block)
	//DBClient.
	//log.Printf("%v", recv)
	c.JSON(200, gin.H{"message": "OK"})
}

func SendPeerStatus(c *gin.Context) {
	var peerStat model.PeerStatus

	lastBlock, _ := DBClient.LastBlock()

	peerStat.Success = true
	peerStat.Header = lastBlock
	peerStat.Height = lastBlock.Height

	c.JSON(200, peerStat)
}

//ReceiveBlocks from blockchain
func GetHeight(c *gin.Context) {
	c.JSON(200, gin.H{"success": true, "height": 0, "id": ""})
}
