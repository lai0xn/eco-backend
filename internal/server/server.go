package server

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/lai0xn/squid-tech/config"
	"github.com/lai0xn/squid-tech/internal/middlewares"
	"github.com/lai0xn/squid-tech/internal/router"
	"github.com/lai0xn/squid-tech/pkg/utils"
	echoSwagger "github.com/swaggo/echo-swagger"
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
	router.SetRoutes(e)

	// CORS configuration
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"}, // TODO: change this
		AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.DELETE, echo.OPTIONS},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
	}))

	// Load Config
	config.Load()

	// Logger
	utils.NewLogger()
	e.Use(middlewares.LoggingMiddleware)
}

func (s *Server) Run() {
	e := echo.New()
	s.Setup(e)
	utils.Logger.LogInfo().Msg(e.Start(s.PORT).Error())

}
