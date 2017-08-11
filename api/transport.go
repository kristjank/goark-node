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
		return errors.New("NetHash mismatch - network version mismatch")
	}
	/*if header.Get("version") != viper.GetString("version") {
		return errors.New("Version mismatch")
	}*/
	return nil
}

//GetTransactions Returns a list of peers to client call. Response is in JSON
func GetTransactions(c *gin.Context) {
	err := sanityCheck(c.Request.Header)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": err.Error()})
	} else {

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
}

//ReceiveBlocks from blockchain
func ReceiveBlocks(c *gin.Context) {
	err := sanityCheck(c.Request.Header)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": err.Error()})
	} else {
		var recv model.BlockReceiveStruct
		c.BindJSON(&recv)

		log.Info("New block received - id: ", recv.Block.ID, " height:", recv.Block.Height, " transactions:", len(recv.Block.Transactions), " peer:", c.Request.RemoteAddr)
		/*err := ArkNodeDB.Save(&recv.Block)
		if err != nil {
			log.Error(err.Error())
		}*/

		c.JSON(200, gin.H{"message": "OK"})
	}
}

//SendPeerStatus respondes to other peers about node statuts
func SendPeerStatus(c *gin.Context) {
	err := sanityCheck(c.Request.Header)
	if err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": err.Error()})
	} else {

		var peerStat model.PeerStatus

		lastBlock, err := getLastBlock()
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"success": false, "message": err.Error()})
		} else {
			peerStat.Success = true
			peerStat.Height = lastBlock.Height

			//Filling Header manually - there is a difference in types (fields TotalAmount, TotalFee and Reward)
			peerStat.Header.BlockSignature = lastBlock.BlockSignature
			peerStat.Header.GeneratorPublicKey = lastBlock.GeneratorPublicKey
			peerStat.Header.Height = lastBlock.Height
			peerStat.Header.ID = lastBlock.ID
			peerStat.Header.NumberOfTransactions = lastBlock.NumberOfTransactions
			peerStat.Header.PayloadHash = lastBlock.PayloadHash
			peerStat.Header.PayloadLength = lastBlock.PayloadLength
			peerStat.Header.Reward, _ = strconv.Atoi(lastBlock.Reward)
			peerStat.Header.Timestamp = lastBlock.Timestamp
			peerStat.Header.TotalAmount, _ = strconv.Atoi(lastBlock.TotalAmount)
			peerStat.Header.TotalFee, _ = strconv.Atoi(lastBlock.TotalFee)
			peerStat.Header.Version = lastBlock.Version
			c.JSON(200, peerStat)
		}
	}
}

//GetHeight returns local blockchain height
func GetHeight(c *gin.Context) {
	err := sanityCheck(c.Request.Header)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": err.Error()})
	} else {

		lastBlock, err := getLastBlock()
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"success": false, "message": err.Error()})
		} else {
			c.JSON(200, gin.H{"success": true, "height": lastBlock.Height, "id": lastBlock.ID})
		}
	}
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
	err := sanityCheck(c.Request.Header)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": err.Error()})
	} else {

		log.Debug("Received Tx from network")
		var recv model.TransactionPayload
		var txIDs []string

		c.BindJSON(&recv)

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
}
