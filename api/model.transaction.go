package api

//TransactionType to store the type)
type TransactionType byte

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
	Signatures            string            `json:"signatures,omitempty"`
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
	Success      bool           `json:"success,omitempty"`
	Transactions []*Transaction `json:"transactions,omitempty"`
	Count        string         `json:"count,omitempty"`
	Error        string         `json:"error,omitempty"`
}

//,signatures::json as signatures,rawasset::json as asset
var listQuery = "SELECT id,\"blockId\",type,timestamp,amount,fee,COALESCE(\"vendorField\", '') as \"vendorField\",\"senderId\",COALESCE(\"recipientId\", '') as \"recipientId\",COALESCE(encode(\"senderPublicKey\", 'hex'), '') as \"senderPublicKey\",COALESCE(encode(\"requesterPublicKey\", 'hex'), '') as \"requesterPublicKey\",COALESCE(encode(\"signature\", 'hex'), '') as \"signature\",COALESCE(encode(\"signSignature\", 'hex'), '') as \"signSignature\" FROM transactions"

//QueryTransactions returns TX from node database
func QueryTransactions() ([]*Transaction, error) {
	rows, err := db.Query(listQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	bks := make([]*Transaction, 0)
	for rows.Next() {
		bk := new(Transaction)

		err := rows.Scan(&bk.ID, &bk.Blockid, &bk.Type, &bk.Timestamp, &bk.Amount, &bk.Fee, &bk.VendorField, &bk.SenderID, &bk.RecipientID, &bk.SenderPublicKey, &bk.RequesterPublicKey, &bk.Signature, &bk.SignSignature)
		if err != nil {
			return nil, err
		}
		bks = append(bks, bk)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return bks, nil
}
