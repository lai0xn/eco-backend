package config

import (
	"os"

	"github.com/joho/godotenv"
)

var JWT_SECRET string

func Load() {
	godotenv.Load()
	os.Getenv("JWT_SECRET")
}
