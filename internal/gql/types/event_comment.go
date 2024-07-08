package types

import "github.com/graphql-go/graphql"


var EventCommentType = graphql.NewObject(graphql.ObjectConfig{
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
  },
})
