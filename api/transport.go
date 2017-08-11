package api

import (
	"errors"
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

//sanityCheck - checking if call came from correct network
func sanityCheck(header http.Header) error {
	if header.Get("nethash") != viper.GetString("network.nethash") {
		return errors.New("Headers NOT OK - NetHash mismatch - network version mismatch")
	}
	if header.Get("port") != viper.GetString("port") {
		return errors.New("Headers NOT OK - Port mismatch")
	}
	return nil
}

//SendTransactions Returns a list of peers to client call. Response is in JSON
func SendTransactions(c *gin.Context) {
	err := sanityCheck(c.Request.Header)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"success": false, "message": err.Error()})
		return
	}

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
	err := sanityCheck(c.Request.Header)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"success": false, "message": err.Error()})
	} else {
		c.JSON(200, gin.H{"success": true, "peers": core.EnvironmentParams.Network.PeerList})
	}
}

//ReceiveBlocks from blockchain
func ReceiveBlocks(c *gin.Context) {
	err := sanityCheck(c.Request.Header)
	if err != nil {
		//	log.Error(err.Error())
		c.JSON(http.StatusOK, gin.H{"success": false, "message": err.Error()})
		return
	}

	var recv model.BlockReceiveStruct
	err = c.BindJSON(&recv)

	if err != nil {
		log.Error(err.Error())
	}

	log.Info("New block received - id: ", recv.Block.ID, " height:", recv.Block.Height, " transactions:", len(recv.Block.Transactions), " peer:", c.Request.RemoteAddr)
	lastBlock, _ := getLastBlock()

	if recv.Block.Height-lastBlock.Height == 1 {
		log.Info("Saving block: ", recv.Block.ID, " height:", recv.Block.Height, " transactions:", len(recv.Block.Transactions), " peer:", c.Request.RemoteAddr)
		err := ArkNodeDB.Save(&recv.Block)
		if err != nil {
			log.Error(err.Error())
		}
		c.JSON(200, gin.H{"success": true, "blockId": recv.Block.ID})
	} else {
		c.JSON(200, gin.H{"success": false, "error": "Chain not at the same height. Unable to receive"})
	}

	//}
}

//SendPeerStatus respondes to other peers about node statuts
func SendPeerStatus(c *gin.Context) {
	err := sanityCheck(c.Request.Header)
	if err != nil {
		//log.Error(err.Error())
		c.JSON(http.StatusOK, gin.H{"success": false, "message": err.Error()})
		return
	}

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
	err := sanityCheck(c.Request.Header)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"success": false, "message": err.Error()})
		return
	}

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
	err := sanityCheck(c.Request.Header)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"success": false, "message": err.Error()})
	}

	log.Debug("Received Tx from network")
	var recv model.TransactionPayload
	var txIDs []string

	err = c.BindJSON(&recv)
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

	//preparing response to client
	var txResponse model.TransactionPostResponse
	txResponse.Success = true
	txResponse.TransactionIDs = txIDs

	//sending response
	c.JSON(200, txResponse)

}
