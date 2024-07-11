package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/lai0xn/squid-tech/internal/services"
	"github.com/lai0xn/squid-tech/pkg/types"
	"github.com/lai0xn/squid-tech/pkg/utils"
	"golang.org/x/oauth2"

	"github.com/labstack/echo/v4"
)

type oauthHandler struct {
	srv *services.AuthService
}

func NewOAuthHandler() *oauthHandler {
	return &oauthHandler{
		srv: services.NewAuthService(),
	}
}

func (h *oauthHandler) getConfig(provider string) (*oauth2.Config, error) {
	oauthProvider, exists := types.OAuth2Configs[provider]
	if !exists {
		return nil, echo.NewHTTPError(http.StatusBadRequest, "unknown provider")
	}
	return oauthProvider.Config, nil
}

func (h *oauthHandler) handleLogin(c echo.Context, provider string) error {
	oauthConfig, err := h.getConfig(provider)
	if err != nil {
		return err
	}

	url := oauthConfig.AuthCodeURL("state")
	return c.Redirect(http.StatusTemporaryRedirect, url)
}

func (h *oauthHandler) handleCallback(c echo.Context, provider string) error {
	oauthConfig, err := h.getConfig(provider)
	if err != nil {
		return err
	}

	code := c.QueryParam("code")
	if code == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "missing code")
	}

	// Exchange the code for a token
	token, err := oauthConfig.Exchange(context.Background(), code)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "failed to exchange token: "+err.Error())
	}

	client := oauthConfig.Client(context.Background(), token)
	userInfo, err := client.Get(getUserInfoURL(provider))
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

	// Check if the user exists
	existingUser, err := h.srv.GetUserByEmail(user.Email)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "failed to sqdfqsdfqsdf user existence")
	}

	// If user doesn't exist, create a new user
	if existingUser == nil {
		if _, err := h.srv.CreateUser(user.Name, user.Email, "", false); err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, "failed to create user: "+err.Error())
		}
	}

	// Generate JWT token
	tokenString, err := utils.GenerateJWT(user.ID, user.Email, user.Name)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "failed to generate token: "+err.Error())
	}

	return c.JSON(http.StatusOK, types.Response{
		"token": tokenString,
	})
}

// handleLogin handles the redirect to Google's OAuth2 login page.
//
//	@Summary	Initiates Google OAuth2 login
//	@Tags		auth
//	@Produce	json
//	@Router		/oauth/google/login [get]
func (h *oauthHandler) GoogleLogin(c echo.Context) error {
	return h.handleLogin(c, "google")
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
	return h.handleCallback(c, "google")
}

// facebookLogin handles the redirect to Facebook's OAuth2 login page.
//
//	@Summary	Initiates Facebook OAuth2 login
//	@Tags		auth
//	@Produce	json
//	@Router		/oauth/facebook/login [get]
func (h *oauthHandler) FacebookLogin(c echo.Context) error {
	return h.handleLogin(c, "facebook")
}

// FacebookCallback handles the OAuth2 callback from Facebook and processes the user info.
//
//	@Summary	Handles Facebook OAuth2 callback
//	@Tags		auth
//	@Accept		json
//	@Produce	json
//	@Param		code	query	string	true	"The OAuth2 authorization code"
//	@Router		/oauth/facebook/callback [get]
func (h *oauthHandler) FacebookCallback(c echo.Context) error {
	return h.handleCallback(c, "facebook")
}

func getUserInfoURL(provider string) string {
	switch provider {
	case "google":
		return "https://www.googleapis.com/oauth2/v2/userinfo"
	case "facebook":
		return "https://graph.facebook.com/v13.0/me?fields=id,name,email"
	default:
		return ""
	}
}
