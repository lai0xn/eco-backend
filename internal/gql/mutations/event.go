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
