package config

import (
	"os"

	"github.com/markbates/goth"
	"github.com/markbates/goth/providers/google"
)

func InitOAuth() {
	// Required even with Redis
	secret := os.Getenv("SESSION_SECRET")
	if secret == "" {
		panic("SESSION_SECRET must be set")
	}

	goth.UseProviders(
		google.New(
			os.Getenv("GOOGLE_CLIENT_ID"),
			os.Getenv("GOOGLE_CLIENT_SECRET"),
			os.Getenv("GOOGLE_CALLBACK_URL"),
		),
	)
}
