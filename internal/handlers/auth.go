package handlers

import (
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"github.com/lai0xn/squid-tech/config"
	"github.com/lai0xn/squid-tech/internal/services"
	"github.com/lai0xn/squid-tech/pkg/types"
	"github.com/lai0xn/squid-tech/pkg/utils"
)

type authHandler struct{
  srv *services.AuthService
}

func NewAuthHandler()*authHandler{
  return &authHandler{
    srv: services.NewAuthService(),
  }
}

// Login example
//
//	@Summary	Login endpoint
//	@Tags		auth
//	@Accept		json
//	@Produce	json
//	@Param		email		body		string	true	"email"
//	@Param		password	body		string	true	"password"
//	@Success	200			{object}	string
//	@Router		/auth/login [post]
func(s *authHandler) Login(c echo.Context) error {
  
	 var payload types.LoginPayload
  if err := c.Bind(&payload);err!= nil {
    return echo.NewHTTPError(http.StatusBadRequest,err.Error())
  }
  err := validate.Struct(payload)
  if err != nil {
    e := err.(validator.ValidationErrors)
    return c.JSON(http.StatusBadRequest,utils.NewValidationError(e))
  }
  user,err := s.srv.CheckUser(payload.Email,payload.Password)
  if err!=nil {
    return echo.NewHTTPError(http.StatusBadRequest,err)
  } 
  token := jwt.NewWithClaims(jwt.SigningMethodHS256,&types.Claims{
    user.Name,
    user.Email,
    jwt.RegisteredClaims{
      ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72)),
    }, 
  })
  token_string,err := token.SignedString([]byte(config.JWT_SECRET))
  if err != nil {
    return echo.NewHTTPError(http.StatusBadRequest,err.Error())
  }
	return c.JSON(http.StatusOK, types.Response{
    "token":token_string,
  })
}



// Registration example
//
//	@Summary	Registration endpoint
//	@Tags		auth
//	@Accept		json
//	@Produce	json
//	@Param		email		body		string	true	"example@gmail.com"
//	@Param		password	body		string	true	"password"
//	@Param		full_name	body		string	true	"aymen charfaoui"
//	@Success	200			{object}	string
//	@Router		/auth/register [post]
func(s *authHandler) Register(c echo.Context) error {
  var payload types.RegisterPayload
  if err := c.Bind(&payload);err!= nil {
    return echo.NewHTTPError(http.StatusBadRequest,err.Error())
  }
  err := validate.Struct(payload)
  if err != nil {
    e := err.(validator.ValidationErrors)
    return c.JSON(http.StatusBadRequest,utils.NewValidationError(e))
  }
  if err := s.srv.CreateUser(payload.Name,payload.Email,payload.Password);err!=nil {
    return echo.NewHTTPError(http.StatusBadRequest,err)
  } 
	return c.JSON(http.StatusOK, types.Response{
    "message":"user created successfully",
  })
}
