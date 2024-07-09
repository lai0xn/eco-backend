package types

import (
	"time"

	"github.com/graphql-go/graphql"
)

var UserType = graphql.NewObject(graphql.ObjectConfig{
  Name: "User",
  Fields: graphql.Fields{
    "id":&graphql.Field{
      Type: graphql.String,
    },
    "email":&graphql.Field{
      Type: graphql.String,
    },
    "name":&graphql.Field{
      Type: graphql.String,
    },
    "bio":&graphql.Field{
      Type: graphql.String,
    },
    "phone":&graphql.Field{
      Type: graphql.String,
    },
    "address":&graphql.Field{
      Type: graphql.String,
    },
    "joined":&graphql.Field{
      Type: graphql.DateTime,
    },
    "image":&graphql.Field{
      Type: graphql.String,
    },
    "gender":&graphql.Field{
      Type: graphql.Boolean,
    },


  },
})

type User struct {
  ID string `json:"id"`
  Name string `json:"name"`
  Phone string `json:"phone"`
  Email string `json:"email"`
  Address string `json:"address"`
  Image string `json:"image"`
  Joined time.Time `json:"joinedAt"`
  Gender bool `json:"gender"`
}
