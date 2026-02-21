package config

import (
	"encoding/hex"
	"fmt"

	"github.com/gorilla/sessions"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/github"
	"github.com/markbates/goth/providers/google"
)

func InitOAuth(cfg *Config) error {
	goth.UseProviders(
		google.New(
			cfg.GoogleClientID,
			cfg.GoogleSecret,
			cfg.GoogleCallbackURL,
			"email",
			"profile",
			"https://www.googleapis.com/auth/userinfo.profile",
		),
		github.New(
			cfg.GitHubClientID,
			cfg.GitHubSecret,
			cfg.GitHubCallbackURL,
			"user:email",
		),
	)

	key, err := hex.DecodeString(cfg.SessionKey)
	if err != nil {
		return fmt.Errorf("decoding session secret: %w", err)
	}
	store := sessions.NewCookieStore(key)
	gothic.Store = store
	return nil
}
