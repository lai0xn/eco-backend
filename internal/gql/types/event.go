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
}

type Organization struct {
  ID string `json:"id"`
  Name string `json:"title"`
  Description string `json:"description"`
  Image string `json:"image"`
  OwnerID string `json:"OwnerID"`
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
  org := Organization{
    ID: m.OrganizerID,
    Name: m.Organizer().Name,
    Description: m.Organizer().Description,
    Image: m.Organizer().Image,
    OwnerID: m.Organizer().OwnerID,
    
  }
  return Event{
    ID:m.ID,
    Title: m.Title,
    Description: m.Description,
    Public: m.Public,
    Date: m.Date,
    Organization:org,
    OrganizationId: m.OrganizerID,
    Particapants: participants,
    Images: m.Images,
  }
}
