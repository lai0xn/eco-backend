package types

import "github.com/graphql-go/graphql"

var AchievmentType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Organization",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.String,
		},
		"name": &graphql.Field{
			Type: graphql.String,
		},
		"details": &graphql.Field{
			Type: graphql.String,
		},
		"eventID": &graphql.Field{
			Type: graphql.String,
		},
	},
})
