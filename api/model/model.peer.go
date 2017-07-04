package model

//PeerStatus response sending structure
type PeerStatus struct {
	Success        bool  `json:"success"`
	Height         int   `json:"height"`
	ForgingAllowed bool  `json:"forgingAllowed"`
	CurrentSlot    int   `json:"currentSlot"`
	Header         Block `json:"header"`
}
