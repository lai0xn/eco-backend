package router

import (
	"context"
	"fmt"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/lai0xn/squid-tech/config"
	"github.com/lai0xn/squid-tech/internal/sse"
	"github.com/lai0xn/squid-tech/pkg/redis"
	"github.com/lai0xn/squid-tech/pkg/types"
)

var (
	jwtMiddelware echo.MiddlewareFunc
)

func init() {
	//Initialize the middlware
	config.Load()
	fmt.Println(config.JWT_SECRET)
	jwtMiddelware = echojwt.WithConfig(echojwt.Config{
		SigningKey: []byte(config.JWT_SECRET),
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(types.Claims)
		},
	})

}

func SetRoutes(e *echo.Echo) {
  sse := sse.NewNotifier()
	e.GET("/", func(c echo.Context) error {
    redis.GetClient().Publish(context.Background(),"notifs","hello world")
		return c.String(http.StatusOK, "Server Working")
	})
  e.GET("/notifications",sse.NotificationHandler)
	v1 := e.Group("/api/v1")
	AuthRoutes(v1)
	profileRoutes(v1)
	orgsRoutes(v1)
	OAuthRoutes(v1)
	eventRoutes(v1)
	postRoutes(v1)
}
