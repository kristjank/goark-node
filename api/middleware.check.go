package api

import (
	"fmt"
	"net/http"

	"gopkg.in/gin-gonic/gin.v1"
)

//CheckNetworkHeaders check on middleware level
func CheckNetworkHeaders() gin.HandlerFunc {
	return func(c *gin.Context) {
		a := c.Request.Header.Get("version")
		fmt.Println(a)
		if a != "1.0.0" {
			c.AbortWithStatus(http.StatusBadRequest)
		}
	}
}
