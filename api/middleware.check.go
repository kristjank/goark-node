package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

//CheckNetworkHeaders check on middleware level
//Function setup in goark-node.go -InitRoutes method
func CheckNetworkHeaders() gin.HandlerFunc {
	return func(c *gin.Context) {
		//cc := c.Copy()
		if c.Request.Header.Get("nethash") != viper.GetString("network.nethash") {
			c.AbortWithStatusJSON(http.StatusOK, gin.H{"success": false, "message": "ENETHASH - Headers NOT OK - NetHash mismatch - network version mismatch"})
		} else if c.Request.Header.Get("port") != viper.GetString("port") {
			c.AbortWithStatusJSON(http.StatusOK, gin.H{"success": false, "message": "EPORT - Headers NOT OK - Port mismatch"})
		} else {
			c.Next()
		}
	}
}

//CheckIfChainLoading check on middleware level
//Function setup in goark-node.go -InitRoutes method
func CheckIfChainLoading() gin.HandlerFunc {
	return func(c *gin.Context) {
		if !*IsBlockchainSynced {
			c.AbortWithStatusJSON(http.StatusOK, gin.H{"success": false, "message": "ECHAIN_LOADING - Blockchain is LOADING/Syncing"})
		} else {
			c.Next()
		}
	}
}
