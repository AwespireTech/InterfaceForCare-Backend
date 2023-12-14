package database

import (
	"context"

	"github.com/AwespireTech/RiverCare-Backend/config"
	"github.com/AwespireTech/RiverCare-Backend/models"
)

func GetStewardshipTokensByUser(user string) ([]models.StewardHistory, error) {
	db := GetClient().Database(config.DATABASE_NAME).Collection("stewardshipHistory")
	filter := models.HistoryFilter{
		User: user,
	}
	cur, err := db.Find(context.Background(), filter)
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
	db := GetClient().Database(config.DATABASE_NAME).Collection("eventHistory")
	filter := models.HistoryFilter{
		User: user,
	}
	cur, err := db.Find(context.Background(), filter)
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
	db := GetClient().Database(config.DATABASE_NAME).Collection("event")
	filter := models.HistoryFilter{
		Host: user,
	}
	cur, err := db.Find(context.Background(), filter)
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
