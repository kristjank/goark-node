package main

import (
	"fmt"

	"github.com/kristjank/goark-node/base/model"
	"github.com/spf13/viper"
)

var cfg HooksConfig
var hookReader *viper.Viper

type HooksConfig struct {
	Hooks []struct {
		HookID      string
		CallbackURI string
		HookType    string
		TriggerRule struct {
			RecipientID string
			Amount      int
			Asset       string
			Type        int
			VendorField string
			Height      int
			SenderID    string
		}
	}
}

func initHooksConfig() {
	hookReader := viper.New()

	hookReader.SetConfigType("json")
	hookReader.AddConfigPath("hooks")
	hookReader.SetConfigName("hooks")

	hookReader.ReadInConfig()
	//fmt.Println(viper.GetInt("port"))

	err := hookReader.Unmarshal(&cfg)

	if err != nil {
		fmt.Println("Unmarshal hook settings failes")
	}

	for _, el := range cfg.Hooks {
		fmt.Println(el.TriggerRule.RecipientID)
	}

	fmt.Println(hookReader.AllSettings())
}

func matchHooks(block model.Block) {
	if len(block.Transactions) > 0 {

	}
}

func main() {
	initHooksConfig()
}
