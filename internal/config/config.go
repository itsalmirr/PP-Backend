package config

import (
	"fmt"
	"os"
	"strings"
)

type Config struct {
	DBHost              string
	DBUser              string
	DBPassword          string
	DBName              string
	DBPort              string
	GoogleClientID      string
	GoogleSecret        string
	GoogleCallbackURL   string
	GitHubClientID      string
	GitHubSecret        string
	GitHubCallbackURL   string
	RedisURL            string
	SessionKey          string
	CloudinaryCloudName string
	CloudinaryAPIKey    string
	CloudinaryAPISecret string
	FrontendURL         string
	CookieSecure        bool
}

func LoadConfig() (*Config, error) {
	var missing []string
	get := func(key string) string {
		value, exists := os.LookupEnv(key)
		if !exists {
			missing = append(missing, key)
		}
		return value
	}

	cfg := &Config{
		DBHost:              get("DB_HOST"),
		DBUser:              get("DB_USER"),
		DBPassword:          get("DB_PASSWORD"),
		DBName:              get("DB_NAME"),
		DBPort:              get("DB_PORT"),
		GoogleClientID:      get("GOOGLE_CLIENT_ID"),
		GoogleSecret:        get("GOOGLE_CLIENT_SECRET"),
		GoogleCallbackURL:   get("GOOGLE_CALLBACK_URL"),
		GitHubClientID:      get("GITHUB_CLIENT_ID"),
		GitHubSecret:        get("GITHUB_CLIENT_SECRET"),
		GitHubCallbackURL:   get("GITHUB_CALLBACK_URL"),
		RedisURL:            get("REDIS_URL"),
		SessionKey:          get("SESSION_KEY"),
		CloudinaryCloudName: get("CLOUDINARY_CLOUD_NAME"),
		CloudinaryAPIKey:    get("CLOUDINARY_API_KEY"),
		CloudinaryAPISecret: get("CLOUDINARY_API_SECRET"),
	}

	if len(missing) > 0 {
		return nil, fmt.Errorf("missing required environment variables: %s", strings.Join(missing, ", "))
	}

	// Optional config with defaults
	cfg.FrontendURL = os.Getenv("FRONTEND_URL")
	if cfg.FrontendURL == "" {
		cfg.FrontendURL = "http://localhost:3000"
	}

	cfg.CookieSecure = os.Getenv("COOKIE_SECURE") == "true"

	return cfg, nil
}
