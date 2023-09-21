package models

import "time"

// Agreement stroes agreement changing proposals, not the agreement itself
type Agreement struct {
	ID          string    `json:"uid" bson:"_id"`
	Name        string    `json:"name" bson:"name"`
	Agreement   string    `json:"agreement" bson:"agreement"`
	Status      int       `json:"status" bson:"status"`
	CreateTime  time.Time `json:"createTime" bson:"createTime"`
	ExpiredTime time.Time `json:"expiredTime" bson:"expiredTime"`
	Voted       []string  `json:"voted" bson:"voted"`
	RejectCount int       `json:"rejectCount" bson:"rejectCount"`
	SignCount   int       `json:"signCount" bson:"signCount"`
}
