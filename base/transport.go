package base

import (
	"github.com/gin-gonic/gin"
	"github.com/kristjank/goark-node/base/model"
	"github.com/spf13/viper"
)

//SendAutoConfigureParams - send autoconfigure parameters
func SendAutoConfigureParams(c *gin.Context) {
	var resp model.AutoConfigureResponse

	resp.Success = true
	resp.Network.Explorer = viper.GetString("network.explorer")
	resp.Network.Nethash = viper.GetString("network.nethash")
	resp.Network.Symbol = viper.GetString("network.symbol")
	resp.Network.Token = viper.GetString("network.token")
	resp.Network.Version = viper.GetInt("network.version")

	c.JSON(200, resp)
}
