package config

import (
	"os"
	"strings"
	"testing"
)

func setRequiredEnvVars(t *testing.T) {
	t.Helper()
	vars := map[string]string{
		"DB_HOST":                "localhost",
		"DB_USER":               "testuser",
		"DB_PASSWORD":           "testpass",
		"DB_NAME":               "testdb",
		"DB_PORT":               "5432",
		"GOOGLE_CLIENT_ID":      "gid",
		"GOOGLE_CLIENT_SECRET":  "gsecret",
		"GOOGLE_CALLBACK_URL":   "http://localhost/callback",
		"GITHUB_CLIENT_ID":      "ghid",
		"GITHUB_CLIENT_SECRET":  "ghsecret",
		"GITHUB_CALLBACK_URL":   "http://localhost/gh/callback",
		"REDIS_URL":             "localhost:6379",
		"SESSION_KEY":           strings.Repeat("ab", 32),
		"CLOUDINARY_CLOUD_NAME": "cloud",
		"CLOUDINARY_API_KEY":    "key",
		"CLOUDINARY_API_SECRET": "secret",
	}
	for k, v := range vars {
		t.Setenv(k, v)
	}
}

func TestLoadConfig_Success(t *testing.T) {
	setRequiredEnvVars(t)

	cfg, err := LoadConfig()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if cfg.DBHost != "localhost" {
		t.Errorf("expected DBHost=localhost, got %s", cfg.DBHost)
	}
	if cfg.DBPort != "5432" {
		t.Errorf("expected DBPort=5432, got %s", cfg.DBPort)
	}
}

func TestLoadConfig_MissingVars(t *testing.T) {
	// Clear all env vars that LoadConfig needs
	os.Clearenv()

	_, err := LoadConfig()
	if err == nil {
		t.Fatal("expected error for missing env vars")
	}
	if !strings.Contains(err.Error(), "DB_HOST") {
		t.Errorf("expected error to mention DB_HOST, got: %s", err.Error())
	}
}

func TestLoadConfig_FrontendURLDefault(t *testing.T) {
	setRequiredEnvVars(t)

	cfg, err := LoadConfig()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if cfg.FrontendURL != "http://localhost:3000" {
		t.Errorf("expected default FrontendURL, got %s", cfg.FrontendURL)
	}
}

func TestLoadConfig_FrontendURLCustom(t *testing.T) {
	setRequiredEnvVars(t)
	t.Setenv("FRONTEND_URL", "https://myapp.com")

	cfg, err := LoadConfig()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if cfg.FrontendURL != "https://myapp.com" {
		t.Errorf("expected custom FrontendURL, got %s", cfg.FrontendURL)
	}
}

func TestLoadConfig_CookieSecure(t *testing.T) {
	setRequiredEnvVars(t)

	cfg, _ := LoadConfig()
	if cfg.CookieSecure {
		t.Error("expected CookieSecure=false by default")
	}

	t.Setenv("COOKIE_SECURE", "true")
	cfg, _ = LoadConfig()
	if !cfg.CookieSecure {
		t.Error("expected CookieSecure=true when COOKIE_SECURE=true")
	}
}
