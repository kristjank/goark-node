package base

import (
	"github.com/kristjank/goark-node/base/model"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var hookReader *viper.Viper

func init() {
	hookReader = viper.New()
	hookReader.SetConfigType("json")
	hookReader.AddConfigPath("hooks")
	hookReader.SetConfigName("hook")

	var C model.Hook

	err := hookReader.Unmarshal(&C)

	if err != nil {
		log.Error("Unmarshal hook settings failes")
	}

	log.Println(hookReader.GetString("id"))
}

func matchHooks(block model.Block) {
	if len(block.Transactions) > 0 {

	}
}
