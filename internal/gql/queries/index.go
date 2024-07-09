package queries

import "github.com/graphql-go/graphql"


var RootQuery = graphql.NewObject(graphql.ObjectConfig{
  Name:"rootQuery",
  Fields: graphql.Fields{
      "Event":eventQuery,
      "Application":applicationsQuery,
      "SearchEvent":searchEventQuery,
      "OrgEvents":orgEventsQuery,
    },
  },
)
