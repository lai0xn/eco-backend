package router

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func eventRoutes(e echo.Group) {
	events := e.Group("/events")
	events.Use(jwtMiddelware)
	events.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Authorized")
	})

}
