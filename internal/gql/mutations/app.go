package mutations

import (
	"github.com/graphql-go/graphql"
	"github.com/lai0xn/squid-tech/internal/gql/resolvers"
	"github.com/lai0xn/squid-tech/internal/gql/types"
)



var createApp = &graphql.Field{
  Type: types.ApplicationType,
  Args: graphql.FieldConfigArgument{
    "eventId":&graphql.ArgumentConfig{
      Type: graphql.String,
    },
    "motivation":&graphql.ArgumentConfig{
      Type: graphql.String,
    },
    "extra":&graphql.ArgumentConfig{
      Type: graphql.String,
    },
  },
  Resolve: resolvers.AppResolver.CreateApp,
}


var deleteApp = &graphql.Field{
  Type: types.ApplicationType,
  Args: graphql.FieldConfigArgument{
    "eventId":&graphql.ArgumentConfig{
      Type: graphql.String,
    },
    "motivation":&graphql.ArgumentConfig{
      Type: graphql.String,
    },
    "extra":&graphql.ArgumentConfig{
      Type: graphql.String,
    },
  },
  Resolve: resolvers.AppResolver.DeleteApp,
}


