package handlers

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

var (
	validate = validator.New(validator.WithRequiredStructEnabled())
)

type Handler interface {
	Get(c echo.Context) error
	Create(c echo.Context) error
	Update(c echo.Context) error
	Delete(c echo.Context) error
}
