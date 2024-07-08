package server

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/lai0xn/squid-tech/config"
	"github.com/lai0xn/squid-tech/internal/gql"
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
	// Load Config
	config.Load()

	e.GET("/swagger/*", echoSwagger.WrapHandler)
	e.Static("/public", "public")
	router.SetRoutes(e)

	// CORS configuration
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"}, // TODO: change this
		AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.DELETE, echo.OPTIONS},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
	}))

	// Logger
	utils.NewLogger()
	e.Use(middlewares.LoggingMiddleware)
}

func (s *Server) Run() {
	e := echo.New()
	s.Setup(e)
  utils.Logger.LogInfo().Msg("graphql server running on port 5000") 
  go gql.Execute()
	utils.Logger.LogInfo().Msg(e.Start(s.PORT).Error())

}
