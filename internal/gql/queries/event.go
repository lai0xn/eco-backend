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


var eventsQuery = &graphql.Field{
  Type: graphql.NewList(types.EventType),
  Args: graphql.FieldConfigArgument{
    "page":&graphql.ArgumentConfig{
      Type: graphql.Int,
    },
  },
  Resolve: resolvers.EventResolver.GetEvents,
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

var orgEventsQuery = &graphql.Field{
  Type: graphql.NewList(types.EventType),
  Args: graphql.FieldConfigArgument{
    "id":&graphql.ArgumentConfig{
      Type: graphql.String,
    },
  },
  Resolve: resolvers.EventResolver.OrgEvents,
}

var eventComment = &graphql.Field{
  Type: graphql.NewList(types.EventCommentType),
  Args: graphql.FieldConfigArgument{
    "id":&graphql.ArgumentConfig{
      Type: graphql.String,
    },
  },
  Resolve: resolvers.EventResolver.GetComment,
}





