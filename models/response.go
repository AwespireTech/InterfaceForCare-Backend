package models

type RiversResponse struct {
	Rivers []River `json:"river"`
	Count  int     `json:"count"`
}
