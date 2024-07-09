package types

import (
	"time"

	"github.com/graphql-go/graphql"
	"github.com/lai0xn/squid-tech/prisma/db"
)


var EventType = graphql.NewObject(graphql.ObjectConfig{
  Name: "Event",
  Fields: graphql.Fields{
    "id":&graphql.Field{
      Type: graphql.String,
    },
    "title":&graphql.Field{
      Type: graphql.String,
    },
    "description":&graphql.Field{
      Type: graphql.String,
    },
    "date":&graphql.Field{
      Type: graphql.DateTime,
    },
    "public":&graphql.Field{
      Type: graphql.Boolean,
    },
    "organizationId":&graphql.Field{
      Type: graphql.Boolean,
    },
    "images":&graphql.Field{
      Type: graphql.NewList(graphql.String),
    },
    "participants":&graphql.Field{
      Type: graphql.NewList(UserType),
    },


  },
}) 



var EventCreationArgs  = graphql.FieldConfigArgument{
  "title":&graphql.ArgumentConfig{
    Type: graphql.String,
  },
  "description":&graphql.ArgumentConfig{
    Type: graphql.String,
  },
  "date":&graphql.ArgumentConfig{
    Type: graphql.DateTime,
  },
  "public":&graphql.ArgumentConfig{
    Type: graphql.Boolean,
  },
  "organizationId":&graphql.ArgumentConfig{
    Type: graphql.String,
  },
}




type Event struct {
  ID string `json:"id"`
  Title string `json:"title"`
  Description string `json:"description"`
  Public bool `json:"public"`
  Date time.Time `json:"time"`
  OrganizationId string `json:"organizationId"`
  Particapants []User `json:"participants"`
  Images []string `json:"images"`
}


func EventFromModel(m *db.EventModel)Event{
  p := m.Particapnts()
  var participants []User 
  for _,pr := range p {
      var phone string
      var address string
      phone,ok := pr.Phone()
      if !ok {
        phone = "" 
      }
      address,ok = pr.Adress()

      if !ok {
         address = ""
      }

      participant := User{
      ID: pr.ID,
      Name: pr.Name,
      Phone: phone,
      Address: address,
      Email: pr.Email,
      Image: pr.Image,
      Joined: pr.CreatedAt,
      Gender: pr.Gender,
      }
      participants = append(participants,participant)
  }
  return Event{
    ID:m.ID,
    Title: m.Title,
    Description: m.Description,
    Public: m.Public,
    Date: m.Date,
    OrganizationId: m.OrganizerID,
    Particapants: participants,
    Images: m.Images,
  }
}
