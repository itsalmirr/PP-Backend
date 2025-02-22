package config

import (
	"github.com/markbates/goth"
	"github.com/markbates/goth/providers/github"
	"github.com/markbates/goth/providers/google"
)

func InitOAuth(cfg *Config) {
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
}
