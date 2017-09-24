package base

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/kristjank/ark-go/core"
	"github.com/kristjank/goark-node/base/model"
	log "github.com/sirupsen/logrus"
)

//ReceiveTransactions Returns a list of peers to client call. Response is in JSON
func ReceiveTransactions(c *gin.Context) {
	log.Info("Received Tx(s) from network, peer: ", c.ClientIP())
	var recv model.TransactionPayload
	var txIDs []string

	err := c.BindJSON(&recv)
	if err != nil {
		log.Error(err.Error())
	}
	multiBroadCastTx(recv)

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

func multiBroadCastTx(txPayload model.TransactionPayload) {
	numberOfPeers2MultiBroadCastTo := 10
	if numberOfPeers2MultiBroadCastTo > 15 {
		numberOfPeers2MultiBroadCastTo = 15
		log.Warn("Max broadcast number too high - set by user, reseting to value 15")
	}
	log.Info("Starting multibroadcast/multithreaded parallel payout to ", numberOfPeers2MultiBroadCastTo, " number of peers")
	peers := ArkAPIClient.GetRandomXPeers(numberOfPeers2MultiBroadCastTo)
	for i := 0; i < numberOfPeers2MultiBroadCastTo; i++ {
		//treaded function
		go func(txPayload2Send model.TransactionPayload, peer core.Peer) {
			//defer wg.Done()
			arkTmpClient := core.NewArkClientFromPeer(peer)
			res, _, _ := arkTmpClient.RelayNodeTransaction2Nodes(txPayload2Send)
			//TODO - maybe check for success of sending
			if res.Success {
				log.Debug("Transactions retransmited OK ", res.TransactionIDs)
			}
		}(txPayload, peers[i])
	}
}
