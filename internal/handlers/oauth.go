package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/lai0xn/squid-tech/internal/services"
	"github.com/lai0xn/squid-tech/pkg/types"
	"github.com/lai0xn/squid-tech/pkg/utils"

	"github.com/labstack/echo/v4"
	"github.com/lai0xn/squid-tech/config"
)

type oauthHandler struct {
	srv *services.AuthService
}

func NewOAuthHandler() *oauthHandler {
	return &oauthHandler{
		srv: services.NewAuthService(),
	}
}

// GoogleLogin handles the redirect to Google's OAuth2 login page.
//
//	@Summary	Initiates Google OAuth2 login
//	@Tags		auth
//	@Produce	json
//	@Router		/oauth/google/login [get]
func (h *oauthHandler) GoogleLogin(c echo.Context) error {
	if config.OAuth2Config == nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "OAuth2Config is not initialized")
	}
	url := config.OAuth2Config.AuthCodeURL("state")
	return c.Redirect(http.StatusTemporaryRedirect, url)
}

// GoogleCallback handles the OAuth2 callback from Google and processes the user info.
//
//	@Summary	Handles Google OAuth2 callback
//	@Tags		auth
//	@Accept		json
//	@Produce	json
//	@Param		code	query	string	true	"The OAuth2 authorization code"
//	@Router		/oauth/google/callback [get]
func (h *oauthHandler) GoogleCallback(c echo.Context) error {
	code := c.QueryParam("code")
	if code == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "missing code")
	}

	// Exchange the code for a token
	token, err := config.OAuth2Config.Exchange(context.Background(), code)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "failed to exchange token: "+err.Error())
	}

	client := config.OAuth2Config.Client(context.Background(), token)
	userInfo, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo")
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "failed to get user info: "+err.Error())
	}
	defer userInfo.Body.Close()

	var user struct {
		ID    string `json:"id"`
		Name  string `json:"name"`
		Email string `json:"email"`
	}
	if err := json.NewDecoder(userInfo.Body).Decode(&user); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "failed to decode user info: "+err.Error())
	}

	existingUser, err := h.srv.GetUserByEmail(user.Email)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "failed to check user existence: "+err.Error())
	}

	if existingUser == nil {
		if err := h.srv.CreateUser(user.Name, user.Email, ""); err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, "failed to create user: "+err.Error())
		}
	}

	tokenString, err := utils.GenerateJWT(user.Email, user.Name)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "failed to generate token: "+err.Error())
	}

	return c.JSON(http.StatusOK, types.Response{
		"token": tokenString,
	})
}
