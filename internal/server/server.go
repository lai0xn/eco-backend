package server

import (
	"github.com/labstack/echo/v4"
	"github.com/lai0xn/squid-tech/internal/router"
	"github.com/swaggo/echo-swagger"
)

type Server struct {
	PORT string
}

func NewServer(port string) *Server {
	return &Server{
		PORT: port,
	}
}

func (s *Server) Setup(e *echo.Echo) {
	e.GET("/swagger/*", echoSwagger.WrapHandler)
  e.Static("/public","public")
	router.SetRoutes(e)

}

func (s *Server) Run() {
	e := echo.New()
	s.Setup(e)
	e.Start(s.PORT)
}
