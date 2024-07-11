package router

import (
	"github.com/labstack/echo/v4"
	"github.com/lai0xn/squid-tech/internal/handlers"
)

func orgsRoutes(e *echo.Group) {
	h := handlers.NewOrgHandler()
	g := e.Group("/organizations")
	g.Use(jwtMiddelware)
	g.GET("/org/get/:id", h.Get)
	g.GET("/me", h.MyOrgs)
	g.GET("/org/search", h.Search)
	g.POST("/create", h.Create)
	g.POST("/org/follow/:id", h.FollowHandler)
	g.PATCH("/org/:id/pfp", h.ChangePfp)
	g.PATCH("/org/:id/bg", h.ChangeBg)
	g.PATCH("/org/update/:id", h.Update)
	g.DELETE("/org/delete/:id", h.Delete)
}
