package router

import (
	"github.com/labstack/echo/v4"
	"github.com/lai0xn/squid-tech/internal/handlers"
)

func profileRoutes(e *echo.Group) {
  p := e.Group("/profiles")
  h := handlers.NewProfileHandler()
  p.Use(jwtMiddelware)
  p.GET("/get/:id",h.Get)
  p.GET("/profile",h.CurrentUser)
  p.GET("/search",h.Search)
  p.PATCH("/profile/pfp",h.ChangePfp)
  p.PATCH("/profile/update",h.Update)
  p.DELETE("/profile/delete",h.Delete)



}
