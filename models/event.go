package models

type Event struct {
	ID                string   `json:"uid" bson:"_id"`
	Name              string   `json:"name" bson:"name"`
	ImageUri          string   `json:"imageUri" bson:"imageUri"`
	Description       string   `json:"description" bson:"description"`
	Editions          int      `json:"editions" bson:"editions"`
	Status            int      `json:"status" bson:"status"`
	Participants      []string `json:"participants" bson:"participants"`
	ParticipantsCount int      `json:"participantsCount" bson:"participantsCount"`
	Approvals         []string `json:"approvals" bson:"approvals"`
	ApprovalsCount    int      `json:"approvalsCount" bson:"approvalsCount"`
}
