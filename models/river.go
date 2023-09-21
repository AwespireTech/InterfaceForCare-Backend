package models

import "time"

// River is a struct that represent the river data

type River struct {
	ID            int        `json:"id" bson:"_id"`
	Name          string     `json:"name" bson:"name"`
	Agreement     string     `json:"agreement" bson:"agreement"`
	Generation    int        `json:"gen" bson:"generation"`
	CreatedTime   time.Time  `json:"createdTime" bson:"createdTime"`
	ExpiredTime   time.Time  `json:"expiredTime" bson:"expiredTime"`
	Status        int        `json:"status" bson:"status"`
	Owners        []string   `json:"owners" bson:"owners"`
	OwnersCount   int        `json:"ownersCount" bson:"ownersCount"`
	Tokens        []int      `json:"tokens" bson:"tokens"`
	EventData     []Event    `json:"events" bson:"-"`
	Events        []string   `json:"-" bson:"events"`
	WalletAddress string     `json:"walletAddr" bson:"walletAddress"`
	ProposalData  []Proposal `json:"proposals" bson:"-"`
	Proposals     []string   `json:"-" bson:"proposals"`
}
