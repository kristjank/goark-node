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
		HookID      string `json:"hook-id,omitempty"`
		CallbackURI string `json:"callbackURI,omitempty"`
		Conditions  struct {
			RecipientID string `json:"RecipientID,omitempty"`
			Amount      int    `json:"Amount,omitempty"`
			Asset       string `json:"Asset,omitempty"`
			Type        int    `json:"Type,omitempty"`
			VendorField string `json:"VendorField,omitempty"`
			Height      int    `json:"Height,omitempty"`
			SenderID    string `json:"SenderID,omitempty"`
		} `json:"conditions,omitempty"`
	} `json:"hooks,omitempty"`
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
		fmt.Println(el.Conditions.RecipientID)
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
