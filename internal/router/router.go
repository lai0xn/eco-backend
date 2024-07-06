package router

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func SetRoutes(e *echo.Echo) {
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Server Working")
	})
  v1 := e.Group("/api/v1")
  AuthRoutes(v1)

}
