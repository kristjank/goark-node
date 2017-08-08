package api

import (
	"github.com/asdine/storm"
	"github.com/kristjank/goark-node/api/model"
	log "github.com/sirupsen/logrus"
)

func getLastBlock() (model.Block, error) {
	var results []model.Block
	err := ArkNodeDB.AllByIndex("Height", &results, storm.Limit(1), storm.Reverse())

	if err != nil {
		log.Error(err.Error())
	}

	return results[0], err
}
