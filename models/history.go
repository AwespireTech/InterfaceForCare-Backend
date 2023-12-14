package models

type StewardHistory struct {
	RiverId       string `json:"riverId" bson:"riverId"`
	User          string `json:"user" bson:"user"`
	TokenContract string `json:"tokenContract" bson:"tokenContract"`
	TokenId       int    `json:"tokenId" bson:"tokenId"`
	Generation    int    `json:"generation" bson:"generation"`
}
type EventHistory struct {
	RiverId       string `json:"riverId" bson:"riverId"`
	EventId       string `json:"eventId" bson:"eventId"`
	Generation    int    `json:"generation" bson:"generation"`
	User          string `json:"user" bson:"user"`
	TokenContract string `json:"tokenContract" bson:"tokenContract"`
	TokenId       int    `json:"tokenId" bson:"tokenId"`
}
type HistoryFilter struct {
	User string `json:"user" bson:"user,omitempty"`
	Host string `json:"host" bson:"host,omitempty"`
}
