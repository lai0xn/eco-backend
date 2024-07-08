package queries

import "github.com/graphql-go/graphql"


var RootQuery = graphql.NewObject(graphql.ObjectConfig{
  Name:"rootQuery",
  Fields: graphql.Fields{
    "helloWorld":&graphql.Field{
      Name: "hello World",
      Type: graphql.String,
      Resolve: func(p graphql.ResolveParams) (interface{}, error) {
        return "Graphql server is working",nil
      },
    },
  },
})
