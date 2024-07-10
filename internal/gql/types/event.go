package types

import (
	"encoding/json"
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
      Type: graphql.NewList(graphql.String),
    },
    "organization":&graphql.Field{
      Type: OrgType,
    },
    "images":&graphql.Field{
      Type: graphql.NewList(graphql.String),
    },
    "participants":&graphql.Field{
      Type: graphql.NewList(UserType),
    },


  },
})

var OrgType = graphql.NewObject(graphql.ObjectConfig{
  Name: "Organization",
  Fields: graphql.Fields{
    "id":&graphql.Field{
      Type: graphql.String,
    },
    "name":&graphql.Field{
      Type: graphql.String,
    },
    "description":&graphql.Field{
      Type: graphql.String,
    },
    "ownerId":&graphql.Field{
      Type: graphql.String,
    },
    "image":&graphql.Field{
      Type: graphql.String,
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
  "location":&graphql.ArgumentConfig{
    Type: graphql.String,
  },
}



type Acheivment struct {
  ID string `json:"id"`
  Name string `json:"title"`
  Description string `json:"description"`
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
  Organization Organization `json:"organization"`
  Acheivments []Acheivment
}

type Organization struct {
  ID string `json:"id"`
  Name string `json:"title"`
  Description string `json:"description"`
  Image string `json:"image"`
  OwnerID string `json:"OwnerID"`
}

func EventToStruct(e *db.EventModel)Event {
  var acheivments []Acheivment
  var participants []User
  var org Organization
  eventJson,_ := json.Marshal(e)
  acheivmentsJson,_ := json.Marshal(e.Achievments())
  orgJson,_ := json.Marshal(e.Organizer())
  participantsJson,_ := json.Marshal(e.Particapnts())
  json.Unmarshal(orgJson,&org)
  json.Unmarshal(participantsJson,&participants)
  json.Unmarshal(acheivmentsJson,&acheivments)
  var event = Event {
    Organization:org,
    Acheivments: acheivments,
    Particapants: participants,
  }
  json.Unmarshal(eventJson,&event)

  return event
  

}




