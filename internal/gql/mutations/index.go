package mutations

import "github.com/graphql-go/graphql"

var RootMutation = graphql.NewObject(graphql.ObjectConfig{
	Name: "mutation",
	Fields: graphql.Fields{
		"EventCreate":        createEvent,
		"EventDelete":        deleteEvent,
		"EventJoin":          joinEvent,
		"EventCommemt":       commentEvent,
		"EventCommentDelete": deleteComment,
		"AppCreate":          createApp,
		"AppDelete":          deleteApp,
		"AppAccept":          acceptApp,
	},
})
