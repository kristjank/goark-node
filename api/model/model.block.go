package model

//BlockReceiveStruct struktura prejetih blockov
type BlockReceiveStruct struct {
	Block Block `json:"block"`
}

//Block structure
type Block struct {
	ID                   string        `json:"id"`
	Version              int           `json:"version"`
	Timestamp            int           `json:"timestamp"`
	Height               int           `json:"height"`
	PreviousBlock        string        `json:"previousBlock"`
	NumberOfTransactions int           `json:"numberOfTransactions"`
	TotalAmount          int           `json:"totalAmount"`
	TotalFee             int           `json:"totalFee"`
	Reward               int           `json:"reward"`
	PayloadLength        int           `json:"payloadLength"`
	PayloadHash          string        `json:"payloadHash"`
	GeneratorPublicKey   string        `json:"generatorPublicKey"`
	GeneratorID          string        `json:"generatorId"`
	BlockSignature       string        `json:"blockSignature"`
	Confirmations        int           `json:"confirmations"`
	TotalForged          string        `json:"totalForged"`
	Transactions         []Transaction `json:"transactions,omitempty"`
}
