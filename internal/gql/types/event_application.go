package types

import "github.com/graphql-go/graphql"


var ApplicationType = graphql.NewObject(graphql.ObjectConfig{
  Name:"Application",
  Fields: graphql.Fields{
    "ID" : &graphql.Field{
      Type: graphql.String,
    },
    "eventID":&graphql.Field{
      Type: graphql.String,
    },
    "content":&graphql.Field{
      Type: graphql.String,
    },
    "userId":&graphql.Field{
      Type: graphql.String,
    },
    "motivation":&graphql.Field{
      Type: graphql.String,
    },
    "accepted":&graphql.Field{
      Type: graphql.Boolean,
    },
    "extra":&graphql.Field{
      Type: graphql.String,
    },
  },
})

