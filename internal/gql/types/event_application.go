package types

import "github.com/graphql-go/graphql"

var ApplicationType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Application",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.String,
		},
		"eventID": &graphql.Field{
			Type: graphql.String,
		},
		"event": &graphql.Field{
			Type: EventType,
		},
		"userId": &graphql.Field{
			Type: graphql.String,
		},
		"user": &graphql.Field{
			Type: UserType,
		},
		"motivation": &graphql.Field{
			Type: graphql.String,
		},
		"accepted": &graphql.Field{
			Type: graphql.Boolean,
		},
		"extra": &graphql.Field{
			Type: graphql.String,
		},
	},
})
