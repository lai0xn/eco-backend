package queries

import (
	"github.com/graphql-go/graphql"
	"github.com/lai0xn/squid-tech/internal/gql/types"
)

var applicationQuery = &graphql.Field{
  Type: types.ApplicationType,
  Args: graphql.FieldConfigArgument{
    "id":&graphql.ArgumentConfig{
      Type: graphql.String,
    },
  },
  Resolve: resolvers.EventResolver.Event,
}


var applicationsQuery = &graphql.Field{
  Type: graphql.NewList(types.ApplicationType),
  Args: graphql.FieldConfigArgument{
    "id":&graphql.ArgumentConfig{
      Type: graphql.String,
    },
  },
  Resolve: resolvers.EventResolver.OrgEvents,
}
