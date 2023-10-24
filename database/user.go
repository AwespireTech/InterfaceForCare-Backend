package database

import (
	"context"

	"github.com/AwespireTech/InterfaceForCare-Backend/models"
	"go.mongodb.org/mongo-driver/bson"
)

func GetStewardshipTokensByUser(user string) ([]models.StewardHistory, error) {
	db := GetClient().Database("InterfaceForCare").Collection("stewardshipHistory")
	cur, err := db.Find(context.Background(), bson.D{
		{
			Key: "user", Value: user,
		},
	})
	if err != nil {
		return nil, err
	}
	var stewardshipHistory []models.StewardHistory
	err = cur.All(context.Background(), &stewardshipHistory)
	if err != nil {
		return nil, err
	}
	return stewardshipHistory, nil

}
func GetEventTokensByUser(user string) ([]models.EventHistory, error) {
	db := GetClient().Database("InterfaceForCare").Collection("eventHistory")
	cur, err := db.Find(context.Background(), bson.D{
		{
			Key: "user", Value: user,
		},
	})
	if err != nil {
		return nil, err
	}
	var eventHistory []models.EventHistory
	err = cur.All(context.Background(), &eventHistory)
	if err != nil {
		return nil, err
	}
	return eventHistory, nil

}
func GetHostedEvents(user string) ([]models.Event, error) {
	db := GetClient().Database("InterfaceForCare").Collection("event")
	cur, err := db.Find(context.Background(), bson.D{
		{
			Key: "host", Value: user,
		},
	})
	if err != nil {
		return nil, err
	}
	var events []models.Event
	err = cur.All(context.Background(), &events)
	if err != nil {
		return nil, err
	}
	return events, nil
}
