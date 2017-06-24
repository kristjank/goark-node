package api

//TransactionType to store the type)
type TransactionType byte

//Tx types
const (
	SENDARK         = 0
	SECONDSIGNATURE = 1
	CREATEDELEGATE  = 2
	VOTE            = 3
	MULTISIGNATURE  = 4
)

//Transaction struct - represents structure of ARK.io blockchain transaction
//It is used to post transaction to mainnet and to receive results from arkapi
//Empty fields are emmited by default
type Transaction struct {
	ID                    string            `json:"id,omitempty"`
	Timestamp             int32             `json:"timestamp,omitempty"`
	RecipientID           string            `json:"recipientId,omitempty"`
	Amount                int64             `json:"amount,omitempty"`
	Asset                 map[string]string `json:"asset,omitempty"`
	Fee                   int64             `json:"fee,omitempty"`
	Type                  TransactionType   `json:"type"`
	VendorField           string            `json:"vendorField,omitempty"`
	Signature             string            `json:"signature,omitempty"`
	SignSignature         string            `json:"signSignature,omitempty"`
	SenderPublicKey       string            `json:"senderPublicKey,omitempty"`
	SecondSenderPublicKey string            `json:"secondSenderPublicKey,omitempty"`
	RequesterPublicKey    string            `json:"requesterPublicKey,omitempty"`
	Blockid               string            `json:"blockid,omitempty"`
	Height                int               `json:"height,omitempty"`
	SenderID              string            `json:"senderId,omitempty"`
	Confirmations         int               `json:"confirmations,omitempty"`
}

//PostTransactionResponse structure for call /peer/list
type PostTransactionResponse struct {
	Success        bool     `json:"success"`
	Message        string   `json:"message"`
	Error          string   `json:"error"`
	TransactionIDs []string `json:"transactionIds"`
}

//TransactionPayload to be sent further
type TransactionPayload struct {
	Transactions []*Transaction `json:"transactions"`
}

//TransactionQueryParams for returing filtered list of transactions
type TransactionQueryParams struct {
	ID          string          `url:"id,omitempty"`
	BlockID     string          `url:"blockId,omitempty"`
	SenderID    string          `url:"senderId,omitempty"`
	RecipientID string          `url:"recipientId,omitempty"`
	Limit       int             `url:"limit,omitempty"`
	Offset      int             `url:"offset,omitempty"`
	OrderBy     string          `url:"orderBy,omitempty"` //"Name of column to order. After column name must go 'desc' or 'asc' to choose order type, prefix for column name is t_. Example: orderBy=t_timestamp:desc (String)"
	Type        TransactionType `url:"type,omitempty"`
}

//TransactionResponse structure holds parsed jsong reply from ark-node
//when calling list methods the Transactions [] has results
//when calling get methods the transaction object (Single) has results
type TransactionResponse struct {
	Success           bool          `json:"success"`
	Transactions      []Transaction `json:"transactions"`
	SingleTransaction Transaction   `json:"transaction"`
	Count             string        `json:"count"`
	Error             string        `json:"error"`
}
