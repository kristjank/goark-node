package api

import (
	"fmt"

	"gopkg.in/gin-gonic/gin.v1"
)

//GetTransactions Returns a list of peers to client call. Response is in JSON
func GetTransactions(c *gin.Context) {

	fmt.Println("Successfully connected!")

	c.JSON(200, gin.H{
		"message": "pong",
	})
}
