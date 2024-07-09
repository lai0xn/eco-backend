package router

import (

	"github.com/labstack/echo/v4"
	"github.com/lai0xn/squid-tech/internal/handlers"
)

func eventRoutes(e *echo.Group) {
	events := e.Group("/events")
  h:= handlers.NewEventHandler()
	events.Use(jwtMiddelware)
	events.POST("/event/:id/upload",h.AddImage)

}
