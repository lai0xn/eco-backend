package router

import (

	"github.com/labstack/echo/v4"
	"github.com/lai0xn/squid-tech/internal/handlers"
)

func eventRoutes(e *echo.Group) {
	events := e.Group("/events")
  h:= handlers.NewEventHandler()
	events.Use(jwtMiddelware)
  events.GET("",h.Get)
  events.GET("/event/event/get/:id",h.Get)
  events.GET("/event/event/search",h.Get)
  events.POST("/create",h.Create)
	events.POST("/event/:id/upload",h.AddImage)

}
