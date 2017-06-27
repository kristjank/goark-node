package api

import (
	"log"

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
	a.Transactions = res

	/*c.JSON(200, gin.H{
		"message": "pong",
	})*/

	c.JSON(200, a)
}
