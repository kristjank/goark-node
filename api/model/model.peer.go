package model

//PeerStatus response sending structure
type PeerStatus struct {
	Success        bool  `json:"success"`
	Height         int   `json:"height"`
	ForgingAllowed bool  `json:"forgingAllowed"`
	CurrentSlot    int   `json:"currentSlot"`
	Header         Block `json:"header"`
}

//TransactionPayload - list of tx to send to network
//METHOD POST receive tx payload from network
type TransactionPayload struct {
	Transactions []Transaction `json:"transactions"`
}

//TransactionPostResponse structure for call /peer/transaction
//METHOD POST response structure
type TransactionPostResponse struct {
	Success        bool     `json:"success"`
	Message        string   `json:"message,omitempty"`
	Error          string   `json:"error,omitempty"`
	TransactionIDs []string `json:"transactionIds"`
}
