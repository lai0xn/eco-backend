package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var JWT_SECRET string
var OAuth2Config *oauth2.Config

func Load() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file")
	}

	// OAuth configuration
	clientID := os.Getenv("GOOGLE_CLIENT_ID")
	clientSecret := os.Getenv("GOOGLE_CLIENT_SECRET")
	redirectURL := os.Getenv("REDIRECT_URL")

	OAuth2Config = &oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		RedirectURL:  redirectURL,
		Scopes:       []string{"profile", "email"},
		Endpoint:     google.Endpoint,
	}

	// JWT Secret
	JWT_SECRET = os.Getenv("JWT_SECRET")

	fmt.Printf("GOOGLE_CLIENT_ID: %s\n", clientID)
	fmt.Printf("GOOGLE_CLIENT_SECRET: %s\n", clientSecret)
	fmt.Printf("REDIRECT_URL: %s\n", redirectURL)
	fmt.Printf("JWT_SECRET: %s\n", JWT_SECRET)
}
