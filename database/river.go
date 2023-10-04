package database

import (
	"context"
	"time"

	"github.com/AwespireTech/InterfaceForCare-Backend/models"
)

func GetAllRivers(assending bool) ([]models.River, error) {
	return nil, nil
}

func GetRiverById(id int) (models.River, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	collection := client.Database("InterfaceForCare").Collection("river")
	var river models.River
	err := collection.FindOne(ctx, models.River{ID: id}).Decode(&river)
	river = FillEventAndProposalData(river)
	return river, err
}
func GetEventById(id string) (models.Event, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	collection := client.Database("InterfaceForCare").Collection("event")
	var event models.Event
	err := collection.FindOne(ctx, models.Event{ID: id}).Decode(&event)
	return event, err
}
func GetProposalById(id string) (models.Proposal, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	collection := client.Database("InterfaceForCare").Collection("proposal")
	var proposal models.Proposal
	err := collection.FindOne(ctx, models.Proposal{ID: id}).Decode(&proposal)
	return proposal, err
}
func FillEventAndProposalData(river models.River) models.River {
	var events []models.Event
	for _, eventID := range river.Events {
		event, err := GetEventById(eventID)
		if err != nil {
			continue
		}
		events = append(events, event)
	}
	river.EventData = events
	var proposals []models.Proposal
	for _, proposalID := range river.Proposals {
		proposal, err := GetProposalById(proposalID)
		if err != nil {
			continue
		}
		proposals = append(proposals, proposal)
	}
	river.ProposalData = proposals
	return river

}
func SaveEventAndProposalData(river models.River) (models.River, error) {
	var eventIDs []string
	for _, event := range river.EventData {
		eventIDs = append(eventIDs, event.ID)
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
		defer cancel()
		collection := client.Database("InterfaceForCare").Collection("event")
		_, err := collection.InsertOne(ctx, event)
		if err != nil {
			return river, err
		}
	}
	river.Events = eventIDs
	var proposalIDs []string
	for _, proposal := range river.ProposalData {
		proposalIDs = append(proposalIDs, proposal.ID)
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
		defer cancel()
		collection := client.Database("InterfaceForCare").Collection("proposal")
		_, err := collection.InsertOne(ctx, proposal)
		if err != nil {
			return river, err
		}
	}
	river.Proposals = proposalIDs
	return river, nil
}
