package services

import (
	"context"
	"fmt"

	"github.com/lai0xn/squid-tech/pkg/types"
	"github.com/lai0xn/squid-tech/pkg/utils"
	"github.com/lai0xn/squid-tech/prisma"
	"github.com/lai0xn/squid-tech/prisma/db"
)

type EventsService struct{}

func NewEventsService() *EventsService {
	return &EventsService{}
}

func (s *EventsService) GetEvent(id string) (*db.EventModel, error) {
	ctx := context.Background()
	utils.Logger.LogInfo().Fields(map[string]interface{}{
		"query":  "get event",
		"params": id,
	}).Msg("DB Query")

	result, err := prisma.Client.Event.FindUnique(
		db.Event.ID.Equals(id),
	).With(db.Event.Organizer.Fetch(),db.Event.Particapnts.Fetch()).Exec(ctx)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s *EventsService) GetOrgEvents(id string) ([]db.EventModel, error) {
	ctx := context.Background()
	utils.Logger.LogInfo().Fields(map[string]interface{}{
		"query":  "get org events",
		"params": id,
	}).Msg("DB Query")
	result, err := prisma.Client.Event.FindMany(
		db.Event.OrganizerID.Equals(id),
	).Exec(ctx)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s *EventsService) SearchEvent(name string) ([]db.EventModel, error) {
	ctx := context.Background()
	utils.Logger.LogInfo().Fields(map[string]interface{}{
		"query":  "search event",
		"params": name,
	}).Msg("DB Query")
	result, err := prisma.Client.Event.FindMany(
		db.Event.Or(
        db.Event.Title.Contains(name),
        db.Event.Description.Contains(name),
    ),
	).Exec(ctx)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s *EventsService) UpdateEvent(id string, payload types.EventPayload) (*db.EventModel, error) {
	ctx := context.Background()
	utils.Logger.LogInfo().Fields(map[string]interface{}{
		"query":  "update org",
		"id":     id,
		"params": payload,
	}).Msg("DB Query")
	results, err := prisma.Client.Event.FindUnique(
		db.Event.ID.Equals(id),
	).Update(
		db.Event.Title.Set(payload.Title),
		db.Event.Description.Set(payload.Description),
		db.Event.Public.Set(payload.Public),
    db.Event.Date.Set(payload.Date),
	).Exec(ctx)
	if err != nil {
		return nil, err
	}
	return results, nil
}

func (s *EventsService) CreateEvent(id string, payload types.EventPayload) (*db.EventModel, error) {
	utils.Logger.LogInfo().Fields(map[string]interface{}{
		"query":   "create org",
		"ownerId": id,
		"params":  payload,
	}).Msg("DB Query")
	ctx := context.Background()
	results, err := prisma.Client.Event.CreateOne(
		db.Event.Title.Set(payload.Title),
		db.Event.Description.Set(payload.Description),
    db.Event.Organizer.Link(db.Organization.ID.Equals(id)),
    db.Event.Date.Set(payload.Date),
    db.Event.Public.Set(payload.Public),
	).Exec(ctx)
	if err != nil {
		return nil, err
	}
	return results, nil
}

func (s *EventsService) AddImage(id string, path string) ([]string, error) {
	fmt.Println(id)
	ctx := context.Background()
	result, err := prisma.Client.Event.FindUnique(
		db.Event.ID.Equals(id),
	).Update(
		db.Event.Images.Push([]string{path}),
	).Exec(ctx)
	if err != nil {
		return nil, err
	}
	return result.Images, nil
}

func (s *EventsService)JoinEvent(eventId string,userId string)(*db.EventModel,error) {
  ctx := context.Background()
  result,err:= prisma.Client.Event.FindUnique(
      db.Event.ID.Equals(eventId),
  ).Update(
      db.Event.Particapnts.Link(db.User.ID.Equals(userId)),
  ).Exec(ctx)
  if err != nil {
     return nil,err
  }
  return result,nil
}

func (s *EventsService) DeleteEvent(id string) (string, error) {
	utils.Logger.LogInfo().Fields(map[string]interface{}{
		"query":  "delete org",
		"params": id,
	}).Msg("DB Query")
	ctx := context.Background()
	deleted, err := prisma.Client.Event.FindUnique(
		db.Event.ID.Equals(id),
	).Delete().Exec(ctx)
	if err != nil {
		return "", nil
	}
	fmt.Println(deleted.ID)
	return deleted.ID, nil
}

