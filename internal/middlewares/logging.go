package middlewares

import (
	"github.com/labstack/echo/v4"
	"github.com/lai0xn/squid-tech/pkg/utils"
)

func LoggingMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// log the request
		utils.Logger.LogInfo().Fields(map[string]interface{}{
			"method": c.Request().Method,
			"uri":    c.Request().URL.Path,
			"query":  c.Request().URL.RawQuery,
		}).Msg("Request")

		// call the next middleware/handler
		err := next(c)
		if err != nil {
			utils.Logger.LogError().Fields(map[string]interface{}{
				"error": err.Error(),
			}).Msg("Response")
			return err
		}

		return nil
	}
}
