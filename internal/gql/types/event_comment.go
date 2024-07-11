package types

import "github.com/graphql-go/graphql"

var EventCommentType = graphql.NewObject(graphql.ObjectConfig{
	Name: "EventComment",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.String,
		},
		"eventID": &graphql.Field{
			Type: graphql.String,
		},
		"content": &graphql.Field{
			Type: graphql.String,
		},
		"userId": &graphql.Field{
			Type: graphql.String,
		},
	},
})
