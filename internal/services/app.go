package services

import (
	"context"

	"github.com/lai0xn/squid-tech/pkg/utils"
	"github.com/lai0xn/squid-tech/prisma"
	"github.com/lai0xn/squid-tech/prisma/db"
)

type AppService struct{}

func NewAppService() *AppService {
	return &AppService{}
}


func (s *EventsService) GetApp(id string) (*db.EventApplicationModel, error) {
	ctx := context.Background()
	utils.Logger.LogInfo().Fields(map[string]interface{}{
		"query":  "get app",
		"params": id,
	}).Msg("DB Query")

	result, err := prisma.Client.EventApplication.FindUnique(
		db.EventApplication.ID.Equals(id),
	).With(db.EventApplication.Event.Fetch()).Exec(ctx)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s *EventsService) GetEventApps(id string) ([]db.EventApplicationModel, error) {
	ctx := context.Background()
	utils.Logger.LogInfo().Fields(map[string]interface{}{
		"query":  "get event apps",
		"params": id,
	}).Msg("DB Query")
	result, err := prisma.Client.EventApplication.FindMany(
		db.EventApplication.EventID.Equals(id),
	).Exec(ctx)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s *EventsService) CreateApp(eventId string,userId string,content string,extra string) (*db.EventApplicationModel, error) {
	ctx := context.Background()
	utils.Logger.LogInfo().Fields(map[string]interface{}{
		"query":  "create event app",
	}).Msg("DB Query")
	result, err := prisma.Client.EventApplication.CreateOne(
		db.EventApplication.Event.Link(db.Event.ID.Equals(eventId)),
    db.EventApplication.User.Link(db.User.ID.Equals(userId)),
    db.EventApplication.Motivation.Set(content),
    db.EventApplication.Accepted.Equals(false),
    db.EventApplication.Extra.Set(extra),  

	).Exec(ctx)
	if err != nil {
		return nil, err
	}
	return result, nil
}
