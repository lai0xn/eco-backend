package resolvers

import (
	"errors"
	"fmt"
	"time"

	"github.com/graphql-go/graphql"
	t "github.com/lai0xn/squid-tech/internal/gql/types"
	"github.com/lai0xn/squid-tech/internal/middlewares/gql"
	"github.com/lai0xn/squid-tech/internal/services"
	"github.com/lai0xn/squid-tech/pkg/types"
)

func NewEventResolver() *eventResolver{
  return &eventResolver{
    srv:services.NewEventsService(),
    orgSrv:services.NewOrgService(),
  }
}

//Event Resolver
type eventResolver struct {
  srv *services.EventsService
  orgSrv *services.OrgService
}

func(r *eventResolver) hasPerm(p graphql.ResolveParams)error{
   orgId := p.Args["organizationId"].(string)
   u := p.Context.Value("user")
   if u == nil {
    return errors.New("Not Authorized")
   }
   user := u.(*types.Claims)
   org,err := r.orgSrv.GetOrg(orgId)
   if err != nil {
    return err
   }
   if org.OwnerID != user.ID {
    return errors.New("Not Authorized")
   }
   return nil
  
}

func(r *eventResolver) isOwner(p graphql.ResolveParams)error{
   eventId := p.Args["id"].(string)
   u := p.Context.Value("user")
   if u == nil {
    return errors.New("Not Authorized")
   }
   user := u.(*types.Claims)
   event,err := r.srv.GetEvent(eventId)
   if err != nil {
    return err
   }
   org := event.Organizer() 
   if  org.OwnerID!= user.ID {
    return errors.New("Not Authorized")
   }
   return nil
  
}
//Get Event By Id

func (r *eventResolver)Event(p graphql.ResolveParams) (interface{},error){
  id,ok := p.Args["id"].(string)
  if !ok {
    return nil ,errors.New("No Args Provided")
  }
  e,err := r.srv.GetEvent(id)
  if err != nil {
    return nil,err
  }
  event := t.EventFromModel(e)
  return event,nil
}

func (r *eventResolver)SearchEvent(p graphql.ResolveParams) (interface{},error){
  title,ok := p.Args["title"].(string)
  if !ok {
    return nil ,errors.New("No Args Provided")
  }
  e,err := r.srv.SearchEvent(title)
  if err != nil {
    return nil,err
  }
  var events []t.Event
  for _,event := range e {
    n := t.EventFromModel(&event)
    events = append(events, n)
  }
  fmt.Println(e)
  return events,nil
}

func (r *eventResolver)OrgEvents(p graphql.ResolveParams) (interface{},error){
  id,ok := p.Args["id"].(string)
  if !ok {
    return nil ,errors.New("No Args Provided")
  }
  e,err := r.srv.GetOrgEvents(id)
  if err != nil {
    return nil,err
  }
  var events []t.Event
  for _,event := range e {
    n := t.EventFromModel(&event)
    events = append(events, n)
  }
  fmt.Println(e)
  return events,nil
}


func (r *eventResolver)DeleteEvent(p graphql.ResolveParams) (interface{},error){
  err := r.isOwner(p)
  if err != nil {
    return nil,err
  }
  id,ok := p.Args["id"].(string)

  if !ok {
    return nil ,errors.New("No Args Provided")
  }
  e,err := r.srv.DeleteEvent(id)
  if err != nil {
    return nil,err
  }

  return e,nil
}

func (r *eventResolver)JoinEvent(p graphql.ResolveParams) (interface{},error){
 
  userId,err := middlewares.IsAuthenticated(p) 
  if err != nil {
    return nil,err
  }
  eventId,ok := p.Args["id"].(string)
  if !ok {
    return "" ,errors.New("No Args Provided")
  }
  e,err := r.srv.JoinEvent(eventId,userId)
  if err != nil {
    return nil,err
  }
  event := t.EventFromModel(e)
  return event,nil
}

func (r *eventResolver)CreateEvent(p graphql.ResolveParams) (interface{},error){
  if err := r.hasPerm(p);err!= nil {
    return nil,errors.New("Access Denied")
  }
  title,ok := p.Args["title"].(string)
  if !ok {
    return nil ,errors.New("No Args Provided")
  }
  description,ok := p.Args["description"].(string)
  if !ok {
    return nil ,errors.New("No Args Provided")
  }
  date,ok := p.Args["date"].(time.Time)
  if !ok {
    return nil ,errors.New("No Args Provided")
  }
  public,ok := p.Args["public"].(bool)
  if !ok {
    return nil ,errors.New("No Args Provided")
  }
  orgID,ok := p.Args["organizationId"].(string)
  if !ok {
    return nil ,errors.New("No Args Provided")
  }
  e,err := r.srv.CreateEvent(orgID,types.EventPayload{
    Title:title,
    Description: description,
    Public: public,
    Date: date,
  })
  if err != nil {
    return nil,err
  }
  event := t.EventFromModel(e)
  return event,nil
}
