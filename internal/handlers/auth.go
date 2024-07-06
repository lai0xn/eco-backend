package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type authHandler struct{}

func NewAuthHandler()*authHandler{
  return &authHandler{}
}

// Login
//
//	@Summary	Login endpoint
//	@Tags		auth
//	@Accept		json
//	@Produce	json
//	@Param		email		body		string	true	"email"
//	@Param		password	body		string	true	"password"
//	@Success	200			{object}	string
//	@Router		/api/v1/auth/login [post]
func(s *authHandler) Login(c echo.Context) error {
	return c.String(http.StatusOK, "login endpoint")
}
