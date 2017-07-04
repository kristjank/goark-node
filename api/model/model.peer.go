package model

//PeerStatus response sending structure
type PeerStatus struct {
	Success        bool `json:"success"`
	Height         int  `json:"height"`
	ForgingAllowed bool `json:"forgingAllowed"`
	CurrentSlot    int  `json:"currentSlot"`
	Header         struct {
		ID                   string `json:"id"`
		Height               int    `json:"height"`
		Version              int    `json:"version"`
		TotalAmount          int64  `json:"totalAmount"`
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
