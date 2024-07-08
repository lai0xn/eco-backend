package queries

import (
	"github.com/graphql-go/graphql"
	"github.com/lai0xn/squid-tech/internal/gql/resolvers"
	"github.com/lai0xn/squid-tech/internal/gql/types"
)


var eventQuery = &graphql.Field{
  Type: types.EventType,
  Args: graphql.FieldConfigArgument{
    "id":&graphql.ArgumentConfig{
      Type: graphql.String,
    },
  },
  Resolve: resolvers.EventResolver.Event,
}

var searchEventQuery = &graphql.Field{
  Type: graphql.NewList(types.EventType),
  Args: graphql.FieldConfigArgument{
    "title":&graphql.ArgumentConfig{
      Type: graphql.String,
    },
  },
  Resolve: resolvers.EventResolver.SearchEvent,
}


