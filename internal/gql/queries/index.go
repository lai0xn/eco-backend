package queries

import "github.com/graphql-go/graphql"


var RootQuery = graphql.NewObject(graphql.ObjectConfig{
  Name:"rootQuery",
  Fields: graphql.Fields{
      "Event":eventQuery,
      "SearchEvent":searchEventQuery,
      "OrgEvents":orgEventsQuery,
    },
  },
)
