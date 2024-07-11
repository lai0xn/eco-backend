package gql

import (

	"github.com/graphql-go/handler"
	"github.com/labstack/echo/v4"
	"github.com/lai0xn/squid-tech/internal/middlewares/gql"
)

func Execute(e *echo.Echo) {
	h := handler.New(&handler.Config{
		Schema:     &Schema,
		Pretty:     true,
		GraphiQL:   true,
		Playground: true,
	})
  
  e.GET("/graphql",func(c echo.Context) error {
    m := middlewares.HeaderMiddleware(h)
    m.ServeHTTP(c.Response(),c.Request())
    return nil

  })
   
  e.POST("/graphql",func(c echo.Context) error {
    m := middlewares.HeaderMiddleware(h)
    m.ServeHTTP(c.Response(),c.Request())
    return nil

  })

}
