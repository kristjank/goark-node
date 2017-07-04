package api

import (
	"log"

	"gopkg.in/gin-gonic/gin.v1"
)

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
	var recv BlockReceiveStruct
	c.BindJSON(&recv)
	log.Printf("%v", recv)

}

func SendPeerStatus(c *gin.Context) {
	var peerStat PeerStatus

	peerStat.Success = true
	//peerStat.Header.

	c.JSON(200, peerStat)
}
