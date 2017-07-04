package model

//BlockReceiveStruct struktura prejetih blockov
type BlockReceiveStruct struct {
	Block Block `json:"block"`
}

//Block structure
type Block struct {
	ID                   string        `json:"id"`
	Height               int           `json:"height"`
	Version              int           `json:"version"`
	TotalAmount          int           `json:"totalAmount"`
	TotalFee             int           `json:"totalFee"`
	Reward               int           `json:"reward"`
	PayloadHash          string        `json:"payloadHash"`
	Timestamp            int           `json:"timestamp"`
	NumberOfTransactions int           `json:"numberOfTransactions"`
	PayloadLength        int           `json:"payloadLength"`
	PreviousBlock        string        `json:"previousBlock"`
	GeneratorPublicKey   string        `json:"generatorPublicKey"`
	BlockSignature       string        `json:"blockSignature"`
	Transactions         []Transaction `json:"transactions,omitempty"`
}
