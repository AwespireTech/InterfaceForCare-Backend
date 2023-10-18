package models

const (
	PROPOSAL_STATUS_DEAD = iota
	PROPOSAL_STATUS_PROPOSED
	PROPOSAL_STATUS_APPROVED
)
const (
	PROPOSAL_TYPE_TRANSFER = iota
	PROPOSAL_TYPE_AGREEMENT
	PROPOSAL_TYPE_PROMPT
)

type Proposal struct {
	ID              string   `json:"uid" bson:"_id"`
	TransactionType int      `json:"transactionType" bson:"transactionType"`
	Agreement       string   `json:"agreement" bson:"agreement"`
	Prompt          []string `json:"prompt" bson:"prompt"`
	TargetAddress   string   `json:"targetAddr" bson:"targetAddress"`
	TransferMutez   int      `json:"transferMutez" bson:"transferMutez"`
	ProposerAddress string   `json:"proposerAddr" bson:"proposerAddress"`
	Status          int      `json:"status" bson:"status"`
	CreatedTime     int64    `json:"createdTime" bson:"createdTime"`
	ExpiredTime     int64    `json:"expiredTime" bson:"expiredTime"`
	Approvals       []string `json:"approvals" bson:"approvals"`
	ApprovalsCount  int      `json:"approvalsCount" bson:"approvalsCount"`
	Generation      int      `json:"gen" bson:"generation"`
}
type ProposalUpdate struct {
}
