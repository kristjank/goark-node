package model

//TransactionType to store the type)
type TransactionType byte

//Transaction struct - represents structure of ARK.io blockchain transaction
//It is used to post transaction to mainnet and to receive results from arkapi
//Empty fields are emmited by default
type Transaction struct {
	ID                    string                 `json:"id" storm:"id"`
	Timestamp             int32                  `json:"timestamp"`
	RecipientID           string                 `json:"recipientId"`
	Amount                int64                  `json:"amount"`
	Asset                 map[string]interface{} `json:"asset"`
	Fee                   int64                  `json:"fee"`
	Type                  TransactionType        `json:"type"`
	VendorField           string                 `json:"vendorField,omitempty"`
	Signature             string                 `json:"signature"`
	SignSignature         string                 `json:"signSignature,omitempty"`
	Signatures            string                 `json:"signatures,omitempty"`
	SenderPublicKey       string                 `json:"senderPublicKey"`
	SecondSenderPublicKey string                 `json:"secondSenderPublicKey,omitempty"`
	RequesterPublicKey    string                 `json:"requesterPublicKey,omitempty"`
	Blockid               string                 `json:"blockid"`
	Height                int                    `json:"height,omitempty"`
	SenderID              string                 `json:"senderId"`
	Confirmations         int                    `json:"confirmations"`
}

//TransactionGetResponse to send resposes
type TransactionGetResponse struct {
	Success      bool          `json:"success"`
	Transactions []Transaction `json:"transactions"`
	Count        string        `json:"count"`
}

//TransactionPayload - list of tx to send to network
//METHOD POST receive tx payload from network
//Parse transaction from Array- e.g. block or post transaction
type TransactionPayload struct {
	Transactions []Transaction `json:"transactions"`
}

//TransactionPostResponse structure for call /peer/transaction
//METHOD POST response structure
//To send response to client
type TransactionPostResponse struct {
	Success        bool     `json:"success"`
	Message        string   `json:"message,omitempty"`
	Error          string   `json:"error,omitempty"`
	TransactionIDs []string `json:"transactionIds"`
}
