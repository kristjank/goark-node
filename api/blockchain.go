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
	for localHeight < ArkAPIClient.GetActivePeer().Height {
		log.Info("Blockchain not in sync. ", fmt.Sprintf("Synced level at: %3.3f%%", float64(localHeight)/float64(ArkAPIClient.GetActivePeer().Height)*100.0), " [from block:", localHeight, " to current blockchain height:", ArkAPIClient.GetActivePeer().Height)

		//Switch peer ever 10K blocks - while syncing with blockchain -
		if peerSwitcher < (localHeight / 10000) {
			peerSwitcher = (localHeight / 10000)
			switchPeer()
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
	log.Info("Switched active peer for blockchain sync: ", ArkAPIClient.GetActivePeer())
}
