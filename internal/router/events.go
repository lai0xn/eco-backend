package router

import (
	"net/http"

	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/lai0xn/squid-tech/pkg/types"
)

func eventRoutes(e echo.Group){
  events := e.Group("/events")
  events.Use(echojwt.WithConfig(echojwt.Config{
    NewClaimsFunc: func(c echo.Context) jwt.Claims {
      return new(types.Claims)
    },
  }))
  events.GET("/",func(c echo.Context) error {
    return c.String(http.StatusOK,"Authorized")
  })

}
