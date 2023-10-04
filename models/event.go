package models

import "time"

type Event struct {
	ID                string    `json:"uid" bson:"_id"`
	Name              string    `json:"name" bson:"name"`
	ImageUri          string    `json:"imageUri,omitempty" bson:"imageUri"`
	TokenId           int       `json:"tokenId" bson:"tokenId"`
	TokenContract     string    `json:"tokenContract" bson:"tokenContract"`
	Description       string    `json:"description" bson:"description"`
	Editions          int       `json:"editions" bson:"editions"`
	Amount            int       `json:"amount" bson:"amount"`
	Status            int       `json:"status" bson:"status"`
	Host              string    `json:"host" bson:"host"`
	CreatedTime       time.Time `json:"createdTime" bson:"createdTime"`
	Participants      []string  `json:"participants" bson:"participants"`
	ParticipantsCount int       `json:"participantsCount" bson:"participantsCount"`
	Approvals         []string  `json:"approvals" bson:"approvals"`
	ApprovalsCount    int       `json:"approvalsCount" bson:"approvalsCount"`
}
