package base

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kristjank/ark-go/core"
	"github.com/kristjank/goark-node/base/model"
)

//SendPeerList Returns a list of peers to client call. Response is in JSON
func SendPeerList(c *gin.Context) {
	c.JSON(200, gin.H{"success": true, "peers": core.EnvironmentParams.Network.PeerList})
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
