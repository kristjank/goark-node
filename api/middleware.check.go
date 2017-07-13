package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

//CheckNetworkHeaders check on middleware level
//NOT USED ATM
func CheckNetworkHeaders() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.Header.Get("nethash") != viper.GetString("network.nethash") {
			c.AbortWithStatus(http.StatusBadRequest)
		}
		if c.Request.Header.Get("version") != viper.GetString("network.version") {
			c.AbortWithStatus(http.StatusBadRequest)
		}
		c.Next()
	}
}
