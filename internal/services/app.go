package services

import (
	"context"
	"log"

	"github.com/lai0xn/squid-tech/pkg/utils"
	"github.com/lai0xn/squid-tech/prisma"
	"github.com/lai0xn/squid-tech/prisma/db"
)

type AppService struct{}

func NewAppService() *AppService {
	return &AppService{}
}


func (s *AppService) GetApp(id string) (*db.EventApplicationModel, error) {
	ctx := context.Background()
	utils.Logger.LogInfo().Fields(map[string]interface{}{
		"query":  "get app",
		"params": id,
	}).Msg("DB Query")

	result, err := prisma.Client.EventApplication.FindUnique(
		db.EventApplication.ID.Equals(id),
	).With(
    db.EventApplication.Event.Fetch(),
    ).Exec(ctx)

	if err != nil {
		return nil, err
	}
	return result, nil
}


func (s *AppService) AcceptApp(id string) (*db.EventApplicationModel, error) {
	ctx := context.Background()
	utils.Logger.LogInfo().Fields(map[string]interface{}{
		"query":  "accept app",
		"params": id,
	}).Msg("DB Query")

	result, err := prisma.Client.EventApplication.FindUnique(
		db.EventApplication.ID.Equals(id),
	).Update(
    db.EventApplication.Accepted.Set(true),
    ).Exec(ctx)
  _, err = prisma.Client.Event.FindUnique(
		db.Event.ID.Equals(result.EventID),
	).Update(
    db.Event.Particapnts.Link(db.User.ID.Equals(result.UserID)),
    ).Exec(ctx)

	if err != nil {
		return nil, err
	}
	return result, nil
}


func (s *AppService) DeleteApp(id string) (*db.EventApplicationModel, error) {
	ctx := context.Background()
	utils.Logger.LogInfo().Fields(map[string]interface{}{
		"query":  "get app",
		"params": id,
	}).Msg("DB Query")

	result, err := prisma.Client.EventApplication.FindUnique(
		db.EventApplication.ID.Equals(id),
	).Delete().Exec(ctx)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s *AppService) GetEventApps(id string) ([]db.EventApplicationModel, error) {
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

func (s *AppService) CreateApp(eventId string,userId string,content string,extra string) (*db.EventApplicationModel, error) {
	ctx := context.Background()
	utils.Logger.LogInfo().Fields(map[string]interface{}{
		"query":  "create event app",
	}).Msg("DB Query")
	result, err := prisma.Client.EventApplication.CreateOne(
		db.EventApplication.Event.Link(db.Event.ID.Equals(eventId)),
    db.EventApplication.User.Link(db.User.ID.Equals(userId)),
    db.EventApplication.Motivation.Set(content),
    db.EventApplication.Accepted.Set(false),
    db.EventApplication.Extra.Set(extra),  

	).Exec(ctx)
	if err != nil {
    log.Println(err.Error())
		return nil, err
	}
	return result, nil
}
