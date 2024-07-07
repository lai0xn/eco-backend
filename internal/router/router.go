package router

import (
	"net/http"

	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/lai0xn/squid-tech/config"
	"github.com/lai0xn/squid-tech/pkg/types"
)

var (
	jwtMiddelware = echojwt.WithConfig(echojwt.Config{
		SigningKey: []byte(config.JWT_SECRET),
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(types.Claims)
		},
	})
)

func SetRoutes(e *echo.Echo) {
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Server Working")
	})
	v1 := e.Group("/api/v1")
	AuthRoutes(v1)
	profileRoutes(v1)
	OAuthRoutes(v1)
}
