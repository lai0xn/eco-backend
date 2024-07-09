package queries

import "github.com/graphql-go/graphql"


var RootQuery = graphql.NewObject(graphql.ObjectConfig{
  Name:"rootQuery",
  Fields: graphql.Fields{
      "Event":eventQuery,
      "Events":eventsQuery,
      "EventComment":eventComment,
      "Application":applicationsQuery,
      "SearchEvent":searchEventQuery,
      "OrgEvents":orgEventsQuery,
    },
  },
)
