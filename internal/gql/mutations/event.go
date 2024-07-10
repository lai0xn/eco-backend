package mutations

import (
	"github.com/graphql-go/graphql"
	"github.com/lai0xn/squid-tech/internal/gql/resolvers"
	"github.com/lai0xn/squid-tech/internal/gql/types"
)

var createEvent = &graphql.Field{
	Type:    graphql.String,
	Args:    types.EventCreationArgs,
	Resolve: resolvers.EventResolver.CreateEvent,
}

var deleteEvent = &graphql.Field{
	Type: graphql.String,
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
	},
	Resolve: resolvers.EventResolver.DeleteEvent,
}

var joinEvent = &graphql.Field{
	Type: graphql.String,
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
	},
	Resolve: resolvers.EventResolver.JoinEvent,
}

var commentEvent = &graphql.Field{
	Type: types.EventCommentType,
	Args: graphql.FieldConfigArgument{
		"eventId": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
		"content": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
	},
	Resolve: resolvers.EventResolver.CreateComment,
}

var deleteComment = &graphql.Field{
	Type: graphql.String,
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
	},
	Resolve: resolvers.EventResolver.DeleteComment,
}
