package models

type Proposal struct {
	ID              string   `json:"uid" bson:"_id"`
	TransactionType int      `json:"transactionType" bson:"transactionType"`
	Amount          int      `json:"amount" bson:"amount"`
	TargetAddress   string   `json:"targetAddr" bson:"targetAddress"`
	ProposerAddress string   `json:"proposerAddr" bson:"proposerAddress"`
	Status          int      `json:"status" bson:"status"`
	CreatedTime     int64    `json:"createdTime" bson:"createdTime"`
	ExpiredTime     int64    `json:"expiredTime" bson:"expiredTime"`
	Voted           []string `json:"voted" bson:"voted"`
	RejectCount     int      `json:"rejectCount" bson:"rejectCount"`
	SignCount       int      `json:"signCount" bson:"signCount"`
}
