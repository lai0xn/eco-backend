package mutations

import (
	"github.com/graphql-go/graphql"
	"github.com/lai0xn/squid-tech/internal/gql/resolvers"
	"github.com/lai0xn/squid-tech/internal/gql/types"
)
var createEvent = &graphql.Field{
  Type: types.EventType,
  Args: types.EventCreationArgs,
  Resolve: resolvers.EventResolver.CreateEvent,
}

var deleteEvent = &graphql.Field{
  Type: graphql.String,
  Args: graphql.FieldConfigArgument{
    "id":&graphql.ArgumentConfig{
      Type: graphql.String,
    },
  },
  Resolve: resolvers.EventResolver.DeleteEvent,
}
