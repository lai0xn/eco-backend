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
  events.GET("/event/acheivment/get/:id",h.GetAcheivment)
  events.POST("/event/acheivment/create",h.CreateAcheivment)
  events.GET("/event/event/search",h.Search)
  events.POST("/create",h.Create)
	events.POST("/event/:id/upload",h.AddImage)
  events.DELETE("/event/acheivment/:id/delete",h.DeleteAcheivment)


}
