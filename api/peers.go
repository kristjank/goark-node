package api

import "gopkg.in/gin-gonic/gin.v1"

func GetPeers(c *gin.Context) {

	c.JSON(200, gin.H{
		"message": "pong",
	})
}
