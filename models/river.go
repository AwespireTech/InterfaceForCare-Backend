package models

import "time"

const (
	RIVER_STATUS_DEAD = iota
	RIVER_STATUS_ALIVE
)

// River is a struct that represent the river data

type River struct {
	ID            string     `json:"id" bson:"_id,omitempty"`
	Name          string     `json:"name" bson:"name"`
	Description   string     `json:"description" bson:"description"`
	Agreement     string     `json:"agreement" bson:"agreement"`
	Dataset       string     `json:"dataset" bson:"dataset"`
	Generation    int        `json:"gen" bson:"generation"`
	CreatedTime   time.Time  `json:"createdTime" bson:"createdTime"`
	ExpiredTime   time.Time  `json:"expiredTime" bson:"expiredTime"`
	Status        int        `json:"status" bson:"status"`
	Owners        []string   `json:"stewards" bson:"owners"`
	OwnersCount   int        `json:"stewardsCount" bson:"ownersCount"`
	TokenId       int        `json:"currentTokenId" bson:"tokenId"`
	TokenContract string     `json:"currentTokenContract" bson:"tokenContract"`
	EventData     []Event    `json:"events" bson:"-"`
	Events        []string   `json:"-" bson:"events"`
	WalletAddress string     `json:"walletAddr" bson:"walletAddress"`
	ProposalData  []Proposal `json:"proposals" bson:"-"`
	Proposals     []string   `json:"-" bson:"proposals"`
}
type RiversParams struct {
	Ascending bool   `json:"ascending"`
	SortBy    string `json:"sortBy"`
}

type RiverFilter struct {
	ID string `json:"id" bson:"_id,omitempty"`
}
