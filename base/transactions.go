package base

import (
	"strconv"

	"github.com/gin-gonic/gin"
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
