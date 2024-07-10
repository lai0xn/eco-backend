package handlers

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/lai0xn/squid-tech/internal/services"
	"github.com/lai0xn/squid-tech/pkg/mail"
	"github.com/lai0xn/squid-tech/pkg/types"
	"github.com/lai0xn/squid-tech/pkg/utils"
)

type authHandler struct {
	srv *services.AuthService
  verifier *mail.EmailVerifier
}

func NewAuthHandler() *authHandler {
	return &authHandler{
		srv: services.NewAuthService(),
    verifier: mail.NewVerifier(),
	}
}

// Login example
//
//	@Summary	Login endpoint
//	@Tags		auth
//	@Accept		json
//	@Produce	json
//	@Param		body		body		types.LoginPayload	true	"Login details"
//	@Router		/auth/login [post]
func (h *authHandler) Login(c echo.Context) error {
	var payload types.LoginPayload
	if err := c.Bind(&payload); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	err := validate.Struct(payload)
	if err != nil {
		e := err.(validator.ValidationErrors)
		return c.JSON(http.StatusBadRequest, utils.NewValidationError(e))
	}
	user, err := h.srv.CheckUser(payload.Email, payload.Password)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	// Generate JWT token using the utility function
	tokenString, err := utils.GenerateJWT(user.ID, user.Email, user.Name)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, types.Response{
		"token": tokenString,
	})
}

// Registration example
//
//	@Summary	Registration endpoint
//	@Tags		auth
//	@Accept		json
//	@Produce	json
//	@Param		body		body		types.RegisterPayload	true	"Registration details"
//	@Router		/auth/register [post]
func (h *authHandler) Register(c echo.Context) error {
	var payload types.RegisterPayload
	if err := c.Bind(&payload); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	err := validate.Struct(payload)
	if err != nil {
		e := err.(validator.ValidationErrors)
		return c.JSON(http.StatusBadRequest, utils.NewValidationError(e))
	}
	//TODO: fix gender
	user,err := h.srv.CreateUser(payload.Name, payload.Email, payload.Password, payload.Gender)
  if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
  err = h.verifier.SendVerfication(user.ID,[]string{user.Email})
  if err != nil {
    return echo.NewHTTPError(http.StatusBadRequest,err)
  }
	return c.JSON(http.StatusOK,types.Response{
    "message":"verification email sent",
    "userID":user.ID,
  })
}




// Email verification example
//
//	@Summary	Verification endpoint
//	@Tags		auth
//	@Accept		json
//	@Produce	json
//	@Param		id		query		string	true	"userid"
//	@Param		otp		query		string	true	"otp"
//	@Router		/auth/verify [post]
func (h *authHandler) VerifyUser(c echo.Context) error {
  id := c.QueryParam("id")
  otp := c.QueryParam("otp")
  if err := h.verifier.Verify(id,otp);err!= nil {
    return echo.NewHTTPError(http.StatusBadRequest,err)
  }
  if err := h.srv.ActivateUser(id);err!=nil{
     return echo.NewHTTPError(http.StatusBadRequest,err)
  }
	return c.JSON(http.StatusOK,types.Response{
    "message":"user activated",
  })
}
