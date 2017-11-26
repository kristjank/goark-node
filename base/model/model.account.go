package model

import "time"

//Account structure
type Account struct {
	Address         string    `json:"id" storm:"id"`
	PublicKey       string    `json:"publicKey" storm:"index"`
	SecondPublicKey string    `json:"secondPublicKey" storm:"index"`
	Vote            string    `json:"vote" storm:"index"`
	Username        string    `json:"username" storm:"index"`
	Balance         int64     `json:"balance"`
	Votebalance     int64     `json:"voteBalance"`
	CreatedAt       time.Time `storm:"index"`
	UpdatedAt       time.Time `storm:"index"`
}
