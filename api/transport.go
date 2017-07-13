package api

import (
	"github.com/gin-gonic/gin"
	"github.com/kristjank/goark-node/api/model"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

//DBClient interface is setup in goark-node.go.
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

//SendPeerStatus respondes to other peers about node statuts
func SendPeerStatus(c *gin.Context) {
	var peerStat model.PeerStatus

	lastBlock, _ := DBClient.LastBlock()

	peerStat.Success = true
	peerStat.Header = lastBlock
	peerStat.Height = lastBlock.Height

	c.JSON(200, peerStat)
}

//GetHeight returns local blockchain height
func GetHeight(c *gin.Context) {
	lastBlock, _ := DBClient.LastBlock()
	c.JSON(200, gin.H{"success": true, "height": lastBlock.Height, "id": lastBlock.ID})
}

//GetAutoConfigureParams - send autoconfigure parameters
func GetAutoConfigureParams(c *gin.Context) {
	var resp model.AutoConfigureResponse

	resp.Success = true
	resp.Network.Explorer = viper.GetString("network.explorer")
	resp.Network.Nethash = viper.GetString("network.nethash")
	resp.Network.Symbol = viper.GetString("network.symbol")
	resp.Network.Token = viper.GetString("network.token")
	resp.Network.Version = viper.GetInt("network.version")

	c.JSON(200, resp)
}

//ReceiveTransactions Returns a list of peers to client call. Response is in JSON
func ReceiveTransactions(c *gin.Context) {
	log.Debug("Received Tx from network")
	var recv model.TransactionPayload
	var txIDs []string

	c.BindJSON(&recv)

	//saving tx to db
	for _, element := range recv.Transactions {
		id, err := DBClient.SaveTransaction(element)
		if err == nil {
			txIDs = append(txIDs, id)
			log.Debug(txIDs)
		}
		log.Debug(element)
	}

	//preparing response to client
	var txResponse model.TransactionPostResponse
	txResponse.Success = true
	txResponse.TransactionIDs = txIDs

	//sending response
	c.JSON(200, txResponse)
}
