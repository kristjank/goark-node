package api

import (
	"log"
	"net/http"

	"gopkg.in/gin-gonic/gin.v1"
)

//GetTransactions Returns a list of peers to client call. Response is in JSON
func GetTransactions(c *gin.Context) {
	res, err := QueryTransactions()
	if err != nil {
		log.Println(err.Error())
	}

	var a TransactionResponse

	a.Success = true
	//logger.pr
	a.Transactions = res

	/*c.JSON(200, gin.H{
		"message": "pong",
	})*/

	c.JSON(200, a)
}

type Block struct {
	ID                   string        `json:"id"`
	Height               int           `json:"height"`
	Version              int           `json:"version"`
	TotalAmount          int           `json:"totalAmount"`
	TotalFee             int           `json:"totalFee"`
	Reward               int           `json:"reward"`
	PayloadHash          string        `json:"payloadHash"`
	Timestamp            int           `json:"timestamp"`
	NumberOfTransactions int           `json:"numberOfTransactions"`
	PayloadLength        int           `json:"payloadLength"`
	PreviousBlock        string        `json:"previousBlock"`
	GeneratorPublicKey   string        `json:"generatorPublicKey"`
	BlockSignature       string        `json:"blockSignature"`
	Transactions         []interface{} `json:"transactions"`
}

//BlockReceiveStruct struktura prejetih blockov
type BlockReceiveStruct struct {
	Block Block `json:"block"`
}

//ReceiveBlocks
func ReceiveBlocks(c *gin.Context) {
	//x, _ := ioutil.ReadAll(c.Request.Body)
	var recv BlockReceiveStruct
	c.BindJSON(&recv)
	log.Printf("%v", recv)

	c.JSON(http.StatusOK, c)
}
