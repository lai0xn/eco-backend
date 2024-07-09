package gql

import (
	"net/http"

	"github.com/graphql-go/handler"
	"github.com/lai0xn/squid-tech/internal/middlewares/gql"
)

func Execute() {
	h := handler.New(&handler.Config{
		Schema:     &Schema,
		Pretty:     true,
		GraphiQL:   true,
		Playground: true,
	})
	s := http.NewServeMux()
  
  s.Handle("/graphql",middlewares.HeaderMiddleware(h))
	http.ListenAndServe(":5000", s)

}
