package model

//PeerStatus response sending structure. It show the current status of the peer
type PeerStatus struct {
	Success        bool  `json:"success"`
	Height         int   `json:"height"`
	ForgingAllowed bool  `json:"forgingAllowed"`
	CurrentSlot    int   `json:"currentSlot"`
	Header         Block `json:"header"`
}

//AutoConfigureResponse to clients - so they can setup communcation accordingly
type AutoConfigureResponse struct {
	Success bool `json:"success"`
	Network struct {
		Nethash  string `json:"nethash,omitempty"`
		Token    string `json:"token,omitempty"`
		Symbol   string `json:"symbol,omitempty"`
		Explorer string `json:"explorer,omitempty"`
		Version  int    `json:"version,omitempty"`
	} `json:"network"`
}
