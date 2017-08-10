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
		log.Error("GetLast block", err.Error())

		//empty database
		if err.Error() == "not found" {
			a := model.Block{Height: 0}

			return a, nil
		}
	}

	return results[0], err
}

//SyncBlockChain syncs blockchain to the lastest block
func SyncBlockChain() {
	resp, _, _ := ArkAPIClient.GetPeerHeight()

	if resp.Success {
		localLastBlock, err := getLastBlock()
		if err != nil {
			log.Error("Error getting local last block", err.Error())
			return
		}

		localHeight := localLastBlock.Height
		for localHeight < resp.Height {
			log.Info("Blockchain not in sync. ", localHeight/resp.Height*100, "% syncing [from block:", localHeight, " to current blockchain height:", resp.Height)
			respData, _, _ := ArkAPIClient.GetFullBlocksFromPeer(localHeight)

			if respData.Success {
				localHeight = saveBlocks2Database(respData.Blocks)
			}
		}
	}
}

func saveBlocks2Database(blocks []model.Block) int {
	tx, err := ArkNodeDB.Begin(true)
	if err != nil {
		return -1
	}
	defer tx.Rollback()

	for _, block := range blocks {
		err = tx.Save(&block)
		if err != nil {
			log.Error(err.Error())
			return -1
		}
	}

	err = tx.Commit()

	localLastBlock, _ := getLastBlock()
	return localLastBlock.Height
}
