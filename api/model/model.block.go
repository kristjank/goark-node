package model

//BlockReceiveStruct from POST request - handling received blocks from the active network
type BlockReceiveStruct struct {
	Block Block `json:"block"`
}

//BlockResponse structure to receive blocks from a random peer  - from GET request
type BlockResponse struct {
	Success bool    `json:"success"`
	Blocks  []Block `json:"blocks"`
	Count   int     `json:"count"`
}

//BlockHeightResponse structure to receive blocks from a random peer - from GET request
type BlockHeightResponse struct {
	Success bool   `json:"success"`
	Height  int    `json:"height"`
	ID      string `json:"id"`
}

//Block structure
type Block struct {
	ID                   string        `json:"id" storm:"id"`
	Version              int           `json:"version"`
	Timestamp            int           `json:"timestamp"`
	Height               int           `json:"height" storm:"index,unique"`
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
