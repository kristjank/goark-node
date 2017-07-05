package api

import (
	"net/http"

	"gopkg.in/gin-gonic/gin.v1"
)

//CheckNetworkHeaders check on middleware level
func CheckNetworkHeaders() gin.HandlerFunc {
	return func(c *gin.Context) {
		a := c.Request.Header.Get("nethash")
		//fmt.Println(a)
		if a != "6e84d08bd299ed97c212c886c98a57e36545c8f5d645ca7eeae63a8bd62d8988" {
			c.AbortWithStatus(http.StatusBadRequest)
		}
		c.Next()
	}
}
