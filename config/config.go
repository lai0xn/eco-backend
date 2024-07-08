package config

import (
	"os"

	"github.com/lai0xn/squid-tech/pkg/types"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/facebook"
	"golang.org/x/oauth2/google"
)

var JWT_SECRET string

func Load() {

	// OAuth configuration
	types.OAuth2Configs = map[string]*types.OAuthProvider{
		"google": {
			Config: &oauth2.Config{
				ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
				ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
				RedirectURL:  os.Getenv("GOOGLE_REDIRECT_URL"),
				Scopes:       []string{"profile", "email"},
				Endpoint:     google.Endpoint,
			},
		},
		"facebook": {
			Config: &oauth2.Config{
				ClientID:     os.Getenv("FACEBOOK_CLIENT_ID"),
				ClientSecret: os.Getenv("FACEBOOK_CLIENT_SECRET"),
				RedirectURL:  os.Getenv("FACEBOOK_REDIRECT_URL"),
				Scopes:       []string{"public_profile", "email"},
				Endpoint:     facebook.Endpoint,
			},
		},
	}

	// JWT Secret
	JWT_SECRET = os.Getenv("JWT_SECRET")
}
