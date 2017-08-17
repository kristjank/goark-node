package base

import (
	"fmt"

	log "github.com/sirupsen/logrus"
)

//SyncBlockChain syncs blockchain to the lastest block
//it is run when starting the node
func SyncBlockChain() {
	setBCStatus(false)

	localLastBlock, err := getLastBlock()
	if err != nil {
		log.Error("Error getting local last block", err.Error())
		return
	}

	localHeight := localLastBlock.Height
	peerSwitcher := 0
	blockChainHeight := ArkAPIClient.GetActivePeer().Height
	//In case peer is behind main nodes...
	if blockChainHeight <= localHeight {
		log.Info("Localheight ", localHeight, "is greater than connected peer height")
		blockChainHeight = switchPeer()
		localHeight -= 50
	}

	for localHeight < blockChainHeight {
		//Switch peer ever 10K blocks - while syncing with blockchain -
		if peerSwitcher < (localHeight / 10000) {
			log.Info("Blockchain not in sync. ", fmt.Sprintf("Synced level at: %3.3f%%", float64(localHeight)/float64(blockChainHeight)*100.0), " [from block:", localHeight, " to current blockchain height:", blockChainHeight)
			peerSwitcher = (localHeight / 10000)
			blockChainHeight = switchPeer()
		}

		respData, err, _ := ArkAPIClient.GetFullBlocksFromPeer(localHeight)
		if err.ErrorObj != nil {
			log.Error("Error getting blocks: ", err.Error(), ArkAPIClient.GetActivePeer())
			log.Info("Blockchain not in sync. ", fmt.Sprintf("Synced level at: %3.3f%%", float64(localHeight)/float64(blockChainHeight)*100.0), " [from block:", localHeight, " to current blockchain height:", blockChainHeight)
			blockChainHeight = switchPeer()
			continue
		}

		if respData.Success {
			localHeight = saveBlocks2Database(respData.Blocks)
		}

		//final alignment - if syncing from 0, a lot of new block can be produced...
		//to align with the lastest produced blocks
		if blockChainHeight-localHeight < 100 {
			blockChainHeight = switchPeer()
		}
	}
	setBCStatus(true)
}

func switchPeer() int {
	ArkAPIClient = ArkAPIClient.SwitchPeer()
	peerHeight := ArkAPIClient.GetActivePeer().Height
	log.Info("Switched active peer for blockchain sync: ", ArkAPIClient.GetActivePeer(), " peer height: ", peerHeight)
	return peerHeight
}

func setBCStatus(status bool) {
	SyncMutex.Lock()
	IsBlockchainSynced = status
	SyncMutex.Unlock()
}

func getBCStatus() bool {
	SyncMutex.RLock()
	defer SyncMutex.RUnlock()
	return IsBlockchainSynced
}
