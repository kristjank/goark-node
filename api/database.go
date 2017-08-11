package api

import (
	"github.com/asdine/storm"
	"github.com/kristjank/goark-node/api/model"
	log "github.com/sirupsen/logrus"
)

//DB FUNCTIONS
func saveBlocks2Database(blocks []model.Block) int {
	tx, err := ArkNodeDB.Begin(true)
	if err != nil {
		return -1
	}
	defer tx.Rollback()

	for _, block := range blocks {
		//saving block
		err = tx.Save(&block)
		if err != nil {
			log.Error(err.Error())
			return -1
		}

		//saving transactions from block - also to transaction bucket...
		if len(block.Transactions) > 0 {
			for _, trans := range block.Transactions {
				//adding blockid before saving to transaction bucket
				trans.Blockid = block.ID
				err = tx.Save(&trans)
				if err != nil {
					log.Error(err.Error())
					return -1
				}
			}
		}
	}
	err = tx.Commit()

	if err != nil {
		log.Error("Error comming Storm transaction", err.Error())
	}

	localLastBlock, _ := getLastBlock()
	return localLastBlock.Height
}

func getLastBlock() (model.Block, error) {
	var results []model.Block
	err := ArkNodeDB.AllByIndex("Height", &results, storm.Limit(1), storm.Reverse())

	if err != nil {
		log.Error("GetLastblock ", err.Error())
		//empty database - i.e. first run
		if err.Error() == "not found" {
			return model.Block{Height: 0}, nil
		}
	}

	return results[0], err
}

func getBlockByID(blockID string) (model.Block, error) {
	var block model.Block
	err := ArkNodeDB.One("ID", blockID, &block)

	if err != nil {
		log.Error(err.Error())
	}
	return block, err
}

func getTransactionByID(transID string) (model.Transaction, error) {
	var trans model.Transaction
	err := ArkNodeDB.One("ID", transID, &trans)

	if err != nil {
		log.Error(err.Error())
	}
	return trans, err
}
