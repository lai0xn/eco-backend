package gql

import (
	"net/http"

	"github.com/graphql-go/handler"
)

func Execute() {
	h := handler.New(&handler.Config{
		Schema:     &Schema,
		Pretty:     true,
		GraphiQL:   true,
		Playground: true,
	})
	s := http.NewServeMux()
  s.Handle("/graphql",h)
	http.ListenAndServe(":5000", s)

}
