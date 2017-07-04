package api

import (
	"log"

	"github.com/kristjank/goark-node/api/model"
	"gopkg.in/gin-gonic/gin.v1"
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
	//x, _ := ioutil.ReadAll(c.Request.Body)
	var recv model.BlockReceiveStruct
	c.BindJSON(&recv)

	DBClient.SaveBlock(recv.Block)
	//DBClient.
	log.Printf("%v", recv)

}

func SendPeerStatus(c *gin.Context) {
	var peerStat model.PeerStatus

	peerStat.Success = true
	//peerStat.Header.

	c.JSON(200, peerStat)
}
