package resolvers

import (
	"errors"

	"github.com/graphql-go/graphql"
	"github.com/lai0xn/squid-tech/internal/services"
	"github.com/lai0xn/squid-tech/pkg/types"
)

func NewAppResolver() *appResolver{
  return &appResolver{
    srv:services.NewAppService(),
    osrv: services.NewOrgService(),
  }
}

type appResolver struct{
  srv *services.AppService
  osrv *services.OrgService
}

func(r *appResolver) hasPerm(p graphql.ResolveParams)error{
   id := p.Args["id"].(string)
   u := p.Context.Value("user")
   if u == nil {
    return errors.New("Not Authorized")
   }
   user := u.(*types.Claims)
   app,err := r.srv.GetApp(id) 
   if err != nil {
    return err
   }
   if app.UserID != user.ID  && app.Event().OrganizerID != user.ID{
    return errors.New("Not Authorized")
   }
   return nil 
}
func(r *appResolver) isOrganizer(p graphql.ResolveParams)error{
   id := p.Args["id"].(string)
   u := p.Context.Value("user")
   if u == nil {
    return errors.New("Not Authorized")
   }
   user := u.(*types.Claims)
   app,err := r.srv.GetApp(id)
   if err != nil {
    return err
   }
   org,err := r.osrv.GetOrg(app.Event().OrganizerID)
   if err != nil {
    return err
  }
   if org.OwnerID != user.ID{
    return errors.New("Not Authorized")
   }
   return nil 
}
func(r *appResolver) isAuthenticated(p graphql.ResolveParams)(string,error){
   u := p.Context.Value("user")
   if u == nil {
    return "",errors.New("Not Authorized")
   }
   user := u.(*types.Claims)
   return user.ID,nil
}


func (r *appResolver)App(p graphql.ResolveParams) (interface{},error){
  err := r.hasPerm(p)
  if err != nil {
    return nil,err
  }
  id,ok := p.Args["id"].(string)
  if !ok {
    return nil ,errors.New("No Args Provided")
  }
  a,err := r.srv.GetApp(id)
  if err != nil {
    return nil,err
  }
  ex,ok := a.Extra()
  if !ok{
    return nil,errors.New("something wrong")
  }
  event := a.Event()
  app := map[string]interface{}{
    "id":a.ID,
    "eventID":a.EventID,
    "motivation":a.Motivation,
    "userId":a.UserID,
    "accepted":a.Accepted,
    "event":map[string]interface{}{
      "title":event.Title,
      "id":event.ID,
      "description":event.Description,
      "date":event.Date,
      "public":event.Public,
      "organizationID":event.OrganizerID,
    },
    "extra":ex,
  }
  return app,nil
}
func (r *appResolver)AcceptApp(p graphql.ResolveParams) (interface{},error){
  if err := r.isOrganizer(p);err!= nil {
    return nil,err
  }
  u := p.Context.Value("user")
  if u == nil {
    return "",errors.New("Unothorized")
  }
  appId,ok := p.Args["id"].(string)
  if !ok {
    return "" ,errors.New("No Args Provided")
  }
  a,err := r.srv.AcceptApp(appId)
  if err != nil {
    return nil,err
  }
  var ex string
  ex,ok = a.Extra()
  if !ok {
    ex = ""
  }
  app := map[string]interface{}{
    "id":a.ID,
    "eventID":a.EventID,
    "motivation":a.Motivation,
    "userId":a.UserID,
    "accepted":a.Accepted,
    "extra":ex,
  }
  return app,nil
}
func (r *appResolver)CreateApp(p graphql.ResolveParams) (interface{},error){
  uId,err := r.isAuthenticated(p)
  if err != nil {
    return nil,err
  }
  id,ok := p.Args["eventId"].(string)
  if !ok {
    return nil ,errors.New("No Args Provided")
  }
  content,ok := p.Args["motivation"].(string)
  if !ok {
    return nil ,errors.New("No Args Provided")
  }
  extra,ok := p.Args["extra"].(string)
  if !ok {
    extra = ""
  }
  
  a,err := r.srv.CreateApp(id,uId,content,extra)
  if err != nil {
    return nil,err
  }
  ex,ok := a.Extra()
  if !ok{
    return nil,errors.New("something wrong")
  }
  app := map[string]interface{}{
    "id":a.ID,
    "eventID":a.EventID,
    "motivation":a.Motivation,
    "userId":a.UserID,
    "accepted":a.Accepted,
    "extra":ex,
  }
  return app,nil
}

func (r *appResolver)DeleteApp(p graphql.ResolveParams) (interface{},error){
  err := r.hasPerm(p)
  if err != nil {
    return nil,err
  }
  id,ok := p.Args["id"].(string)
  if !ok {
    return nil ,errors.New("No Args Provided")
  }
  
  a,err := r.srv.DeleteApp(id)
  if err != nil {
    return nil,err
  }
  ex,ok := a.Extra()
  if !ok{
    return nil,errors.New("something wrong")
  }
  app := map[string]interface{}{
    "id":a.ID,
    "eventID":a.EventID,
    "motivation":a.Motivation,
    "userId":a.UserID,
    "accepted":a.Accepted,
    "extra":ex,
  }
  return app,nil
}
