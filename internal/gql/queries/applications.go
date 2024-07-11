package queries

import (
	"github.com/graphql-go/graphql"
	"github.com/lai0xn/squid-tech/internal/gql/resolvers"
	"github.com/lai0xn/squid-tech/internal/gql/types"
)

var applicationsQuery = &graphql.Field{
	Type: types.ApplicationType,
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
	},
	Resolve: resolvers.AppResolver.App,
}
