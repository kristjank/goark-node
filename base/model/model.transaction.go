package model

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"encoding/hex"
	"log"

	"github.com/bitgoin/address/base58"
	"github.com/kristjank/ark-go/arkcoin"
)

//TransactionType to make it more readable - enum
type TransactionType byte

//Transaction struct - represents structure of ARK.io blockchain transaction
//It is used to post transaction to mainnet and to receive results from arkapi
//Empty fields are emmited by default
type Transaction struct {
	ID                    string                 `json:"id" storm:"id"`
	Timestamp             int32                  `json:"timestamp" storm:"index"`
	RecipientID           string                 `json:"recipientId"`
	Amount                int64                  `json:"amount" storm:"index"`
	Asset                 map[string]interface{} `json:"asset"`
	Fee                   int64                  `json:"fee"`
	Type                  TransactionType        `json:"type"`
	VendorField           string                 `json:"vendorField,omitempty"`
	Signature             string                 `json:"signature"`
	SignSignature         string                 `json:"signSignature,omitempty"`
	Signatures            []string               `json:"signatures,omitempty"`
	SenderPublicKey       string                 `json:"senderPublicKey"`
	SecondSenderPublicKey string                 `json:"secondSenderPublicKey,omitempty"`
	RequesterPublicKey    string                 `json:"requesterPublicKey,omitempty"`
	Blockid               string                 `json:"blockid" storm:"index"`
	Height                int                    `json:"height" storm:"index"`
	SenderID              string                 `json:"senderId"`
	Confirmations         int                    `json:"confirmations"`
	PkIndex               uint64                 `json:"-" storm:"index,increment"`
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

//PostTransactionResponse structure for call /peer/list
type PostTransactionResponse struct {
	Success        bool     `json:"success"`
	Message        string   `json:"message"`
	Error          string   `json:"error"`
	TransactionIDs []string `json:"transactionIds"`
}

func quickHexDecode(data string) []byte {
	res, err := hex.DecodeString(data)
	if err != nil {
		log.Fatal(err.Error())
	}
	return res
}

//ToBytes returns bytearray of the Transaction object to be signed and send to blockchain
func (tx *Transaction) toBytes(skipSignature, skipSecondSignature bool) []byte {
	txBuf := new(bytes.Buffer)
	binary.Write(txBuf, binary.LittleEndian, tx.Type)
	binary.Write(txBuf, binary.LittleEndian, uint32(tx.Timestamp))

	binary.Write(txBuf, binary.LittleEndian, quickHexDecode(tx.SenderPublicKey))

	if tx.RequesterPublicKey != "" {
		res, err := base58.Decode(tx.RequesterPublicKey)
		if err != nil {
			binary.Write(txBuf, binary.LittleEndian, res)
		}
	}

	if tx.RecipientID != "" {
		res, err := base58.Decode(tx.RecipientID)
		if err != nil {
			log.Fatal("Error converting Decoding b58 ", err.Error())
		}
		binary.Write(txBuf, binary.LittleEndian, res)
	} else {
		binary.Write(txBuf, binary.LittleEndian, make([]byte, 21))
	}

	if tx.VendorField != "" {
		vendorBytes := []byte(tx.VendorField)
		if len(vendorBytes) < 65 {
			binary.Write(txBuf, binary.LittleEndian, vendorBytes)

			bs := make([]byte, 64-len(vendorBytes))
			binary.Write(txBuf, binary.LittleEndian, bs)
		}
	} else {
		binary.Write(txBuf, binary.LittleEndian, make([]byte, 64))
	}

	binary.Write(txBuf, binary.LittleEndian, uint64(tx.Amount))
	binary.Write(txBuf, binary.LittleEndian, uint64(tx.Fee))

	switch tx.Type {
	case 1:
		if signature, ok := tx.Asset["signature"].(string); ok {
			binary.Write(txBuf, binary.LittleEndian, quickHexDecode(signature))
		}
	case 2:
		if usernameBytes, ok := tx.Asset["username"].([]byte); ok {
			//usernameBytes := []byte(username)
			binary.Write(txBuf, binary.LittleEndian, usernameBytes)
		}
	case 3:
		if voteBytes, ok := tx.Asset["votes"].([]byte); ok {
			//voteBytes := []byte(tx.Asset["votes"])
			binary.Write(txBuf, binary.LittleEndian, voteBytes)
		}
	}

	if !skipSignature && len(tx.Signature) > 0 {
		binary.Write(txBuf, binary.LittleEndian, quickHexDecode(tx.Signature))
	}

	if !skipSecondSignature && len(tx.SignSignature) > 0 {
		binary.Write(txBuf, binary.LittleEndian, quickHexDecode(tx.SignSignature))
	}

	return txBuf.Bytes()
}

//Verify function verifies if tx is validly signed
//if return == nill verification was succesfull
func (tx *Transaction) Verify() error {
	key, err := arkcoin.NewPublicKey(quickHexDecode(tx.SenderPublicKey), arkcoin.ActiveCoinConfig)
	if err != nil {
		log.Fatal(err.Error())
	}
	trHashBytes := sha256.New()
	trHashBytes.Write(tx.toBytes(true, true))
	return key.Verify(quickHexDecode(tx.Signature), trHashBytes.Sum(nil))
}

//SecondVerify function verifies if tx is validly signed
//if return == nill verification was succesfull
func (tx *Transaction) SecondVerify() error {
	key, err := arkcoin.NewPublicKey(quickHexDecode(tx.SecondSenderPublicKey), arkcoin.ActiveCoinConfig)
	if err != nil {
		log.Fatal(err.Error())
	}
	trHashBytes := sha256.New()
	trHashBytes.Write(tx.toBytes(false, true))
	return key.Verify(quickHexDecode(tx.SignSignature), trHashBytes.Sum(nil))
}
