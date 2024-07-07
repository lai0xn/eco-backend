package router

import (
	"github.com/labstack/echo/v4"
	"github.com/lai0xn/squid-tech/internal/handlers"
)

func authRoutes(e *echo.Group){
  h := handlers.NewAuthHandler()
  auth := e.Group("/auth")
  auth.POST("/register",h.Register)
  auth.POST("/login",h.Login)


}
