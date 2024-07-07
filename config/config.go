package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/lai0xn/squid-tech/pkg/types"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/facebook"
	"golang.org/x/oauth2/google"
)

var JWT_SECRET string

func Load() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file")
	}

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

	fmt.Print("GOOGLE_CLIENT_ID: ", os.Getenv("GOOGLE_CLIENT_ID"), "\n")
	fmt.Print("GOOGLE_CLIENT_SECRET: ", os.Getenv("GOOGLE_CLIENT_SECRET"), "\n")
	fmt.Print("GOOGLE_REDIRECT_URL: ", os.Getenv("GOOGLE_REDIRECT_URL"), "\n")
	fmt.Print("FACEBOOK_CLIENT_ID: ", os.Getenv("FACEBOOK_CLIENT_ID"), "\n")
	fmt.Print("FACEBOOK_CLIENT_SECRET: ", os.Getenv("FACEBOOK_CLIENT_SECRET"), "\n")
	fmt.Print("FACEBOOK_REDIRECT_URL: ", os.Getenv("FACEBOOK_REDIRECT_URL"), "\n")
}
