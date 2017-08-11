package api

import (
	"fmt"

	log "github.com/sirupsen/logrus"
)

//SyncBlockChain syncs blockchain to the lastest block
func SyncBlockChain() {
	localLastBlock, err := getLastBlock()
	if err != nil {
		log.Error("Error getting local last block", err.Error())
		return
	}

	localHeight := localLastBlock.Height
	peerSwitcher := 0
	blockChainHeight := ArkAPIClient.GetActivePeer().Height
	for localHeight < blockChainHeight {
		log.Info("Blockchain not in sync. ", fmt.Sprintf("Synced level at: %3.3f%%", float64(localHeight)/float64(blockChainHeight)*100.0), " [from block:", localHeight, " to current blockchain height:", blockChainHeight)

		//Switch peer ever 10K blocks - while syncing with blockchain -
		if peerSwitcher < (localHeight / 10000) {
			peerSwitcher = (localHeight / 10000)
			switchPeer()
			blockChainHeight = ArkAPIClient.GetActivePeer().Height
		}

		respData, err, _ := ArkAPIClient.GetFullBlocksFromPeer(localHeight)
		if err.ErrorObj != nil {
			log.Error("Error communicating with peer: ", err.Error(), ArkAPIClient.GetActivePeer())
			switchPeer()
			continue
		}

		if respData.Success {
			localHeight = saveBlocks2Database(respData.Blocks)
		}
	}
}

//Helpers
func switchPeer() {
	ArkAPIClient = ArkAPIClient.SwitchPeer()
	log.Info("Switched active peer for blockchain sync: ", ArkAPIClient.GetActivePeer(), " peer height: ", ArkAPIClient.GetActivePeer().Height)
}
