package model

//PeerStatus response sending structure. It show the current status of the peer
type PeerStatus struct {
	Success        bool `json:"success"`
	Height         int  `json:"height"`
	ForgingAllowed bool `json:"forgingAllowed"`
	CurrentSlot    int  `json:"currentSlot"`
	Header         struct {
		ID                   string `json:"id"`
		Height               int    `json:"height"`
		Version              int    `json:"version"`
		TotalAmount          int    `json:"totalAmount"`
		TotalFee             int    `json:"totalFee"`
		Reward               int    `json:"reward"`
		PayloadHash          string `json:"payloadHash"`
		PayloadLength        int    `json:"payloadLength"`
		Timestamp            int    `json:"timestamp"`
		NumberOfTransactions int    `json:"numberOfTransactions"`
		PreviousBlock        string `json:"previousBlock"`
		GeneratorPublicKey   string `json:"generatorPublicKey"`
		BlockSignature       string `json:"blockSignature"`
	} `json:"header"`
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
