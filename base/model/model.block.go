package model

//BlockReceiveStruct from POST request - handling received blocks from the active network
type BlockReceiveStruct struct {
	Block Block `json:"block"`
}

//BlockResponse structure to receive blocks from a random peer  - from GET request
type BlockResponse struct {
	Success bool           `json:"success"`
	Blocks  []BlockTmpSync `json:"blocks"`
	Count   int            `json:"count,omitempty"`
}

//BlockHeightResponse structure to receive blocks from a random peer - from GET request
type BlockHeightResponse struct {
	Success bool   `json:"success"`
	Height  int    `json:"height"`
	ID      string `json:"id"`
}

//Block structure to store block data
type Block struct {
	ID                   string        `json:"id" storm:"id"`
	Version              int           `json:"version"`
	Height               int           `json:"height" storm:"index"`
	Timestamp            int           `json:"timestamp"`
	PreviousBlock        string        `json:"previousBlock"`
	NumberOfTransactions int           `json:"numberOfTransactions"`
	TotalAmount          int           `json:"totalAmount"`
	TotalFee             int           `json:"totalFee"`
	Reward               int           `json:"reward"`
	PayloadLength        int           `json:"payloadLength"`
	PayloadHash          string        `json:"payloadHash"`
	GeneratorPublicKey   string        `json:"generatorPublicKey"`
	BlockSignature       string        `json:"blockSignature"`
	Transactions         []Transaction `json:"transactions"`
	PkIndex              uint64        `json:"-" storm:"index,increment"`
}

//BlockTmpSync structure to store block data
//used for large syncing /peer/blocks/?lastblockheight
//different types TMP FIX until main node update
type BlockTmpSync struct {
	ID                   string        `json:"id"`
	Version              int           `json:"version"`
	Height               int           `json:"height"`
	Timestamp            int           `json:"timestamp"`
	PreviousBlock        string        `json:"previousBlock"`
	NumberOfTransactions int           `json:"numberOfTransactions"`
	TotalAmount          string        `json:"totalAmount"`
	TotalFee             string        `json:"totalFee"`
	Reward               string        `json:"reward"`
	PayloadLength        int           `json:"payloadLength"`
	PayloadHash          string        `json:"payloadHash"`
	GeneratorPublicKey   string        `json:"generatorPublicKey"`
	BlockSignature       string        `json:"blockSignature"`
	Transactions         []Transaction `json:"transactions"`
	PkIndex              uint64        `json:"-"`
}
