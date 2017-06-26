package api

import "gopkg.in/gin-gonic/gin.v1"

//GetTransactions Returns a list of peers to client call. Response is in JSON
func GetTransactions(c *gin.Context) {

	res, _ := QueryTransactions()

	var a TransactionResponse

	a.Success = true
	a.Transactions = res

	/*c.JSON(200, gin.H{
		"message": "pong",
	})*/

	c.JSON(200, a)
}
