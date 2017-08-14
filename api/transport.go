package api

import (
	"net/http"
	"strconv"

	"github.com/asdine/storm"
	"github.com/gin-gonic/gin"
	"github.com/kristjank/ark-go/core"
	"github.com/kristjank/goark-node/api/model"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

//ArkNodeDB interface is setup in goark-node.go.
//interface is visible in the whole package api
var ArkNodeDB *storm.DB

//ArkAPIClient - using to talk with other nodes with the ARK-GO Client API
var ArkAPIClient *core.ArkClient

//IsBlockchainSynced - Global locker (semaphore) - to ignore writes and receive blocks if blocks arent at the height
var IsBlockchainSynced *bool

//SendTransactions Returns a list of peers to client call. Response is in JSON
func SendTransactions(c *gin.Context) {
	tx2Send, err := getTransactions(0)

	if err == nil {
		var response model.TransactionGetResponse
		response.Success = true
		response.Transactions = tx2Send
		response.Count = strconv.Itoa(len(tx2Send))
		c.JSON(200, response)
	} else {
		c.JSON(500, gin.H{"success": false, "message": err.Error()})
	}

}

//SendPeerList Returns a list of peers to client call. Response is in JSON
func SendPeerList(c *gin.Context) {
	c.JSON(200, gin.H{"success": true, "peers": core.EnvironmentParams.Network.PeerList})
}

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

//SendPeerStatus respondes to other peers about node statuts
func SendPeerStatus(c *gin.Context) {
	var peerStat model.PeerStatus
	lastBlock, err := getLastBlock()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"success": false, "message": err.Error()})
	} else {
		peerStat.Success = true
		peerStat.Height = lastBlock.Height
		peerStat.Header = lastBlock
		c.JSON(200, peerStat)
	}
}

//SendHeight returns local blockchain height
func SendHeight(c *gin.Context) {
	lastBlock, err := getLastBlock()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"success": false, "message": err.Error()})
		return
	}

	c.JSON(200, gin.H{"success": true, "height": lastBlock.Height, "id": lastBlock.ID})
}

//SendAutoConfigureParams - send autoconfigure parameters
func SendAutoConfigureParams(c *gin.Context) {
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

	err := c.BindJSON(&recv)
	if err != nil {
		log.Error(err.Error())
	}

	//saving tx to db
	for _, element := range recv.Transactions {
		err := ArkNodeDB.Save(&element)
		if err != nil {
			log.Error(err.Error())
		}
		txIDs = append(txIDs, element.ID)
		log.Debug(txIDs)
		log.Debug(element)
	}

	//sending response
	c.JSON(200, gin.H{"success": true, "transactionIds": txIDs})
}
