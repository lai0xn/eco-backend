package router

import (
	"github.com/labstack/echo/v4"
	"github.com/lai0xn/squid-tech/internal/handlers"
)

func OAuthRoutes(e *echo.Group) {
	h := handlers.NewOAuthHandler()
	oauth := e.Group("/oauth")
	oauth.GET("/google", h.GoogleLogin)
	oauth.GET("/google/callback", h.GoogleCallback)
	oauth.GET("/facebook", h.FacebookLogin)
	oauth.GET("/facebook/callback", h.FacebookCallback)
}
