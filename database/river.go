package database

import (
	"context"
	"log"
	"time"

	"github.com/AwespireTech/RiverCare-Backend/config"
	"github.com/AwespireTech/RiverCare-Backend/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ascending(isAscending bool) int {
	if isAscending {
		return 1
	}
	return -1
}
func GetAllRivers(param models.RiversParams) ([]models.River, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	collection := client.Database(config.DATABASE_NAME).Collection("river")
	var rivers []models.River
	findOptions := options.Find().SetSort(bson.D{{Key: param.SortBy, Value: ascending(param.Ascending)}})
	cursor, err := collection.Find(ctx, bson.D{}, findOptions)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var river models.River
		err := cursor.Decode(&river)
		if err != nil {
			return nil, err
		}
		rivers = append(rivers, river)
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}
	for i, river := range rivers {
		rivers[i] = fillEventAndProposalData(river)
	}
	return rivers, nil
}

func GetRiverById(id string) (models.River, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	collection := client.Database(config.DATABASE_NAME).Collection("river")
	var river models.River
	err := collection.FindOne(ctx, bson.M{
		"_id": id,
	}).Decode(&river)
	river = fillEventAndProposalData(river)
	return river, err
}
func GetEventById(id string) (models.Event, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	collection := client.Database(config.DATABASE_NAME).Collection("event")
	var event models.Event
	err := collection.FindOne(ctx, bson.M{
		"_id": id,
	}).Decode(&event)
	return event, err
}
func GetProposalById(id string) (models.Proposal, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	collection := client.Database(config.DATABASE_NAME).Collection("proposal")
	var proposal models.Proposal
	err := collection.FindOne(ctx, bson.M{
		"_id": id,
	}).Decode(&proposal)
	return proposal, err
}
func fillEventAndProposalData(river models.River) models.River {
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
			log.Println("Proposal:" + proposalID + " error:" + err.Error())
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
		collection := client.Database(config.DATABASE_NAME).Collection("event")
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
		collection := client.Database(config.DATABASE_NAME).Collection("proposal")
		_, err := collection.InsertOne(ctx, proposal)
		if err != nil {
			return river, err
		}
	}
	river.Proposals = proposalIDs
	return river, nil
}
