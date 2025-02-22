package config

import "os"

type Config struct {
	DBHost            string
	DBUser            string
	DBPassword        string
	DBName            string
	DBPort            string
	GoogleClientID    string
	GoogleSecret      string
	GoogleCallbackURL string
	GitHubClientID    string
	GitHubSecret      string
	GitHubCallbackURL string
	RedisURL          string
	SessionKey        string
	SessionSecret     string
}

func LoadConfig() *Config {
	return &Config{
		DBHost:            getEnv("DB_HOST"),
		DBUser:            getEnv("DB_USER"),
		DBPassword:        getEnv("DB_PASSWORD"),
		DBName:            getEnv("DB_NAME"),
		DBPort:            getEnv("DB_PORT"),
		GoogleClientID:    getEnv("GOOGLE_CLIENT_ID"),
		GoogleSecret:      getEnv("GOOGLE_SECRET"),
		GoogleCallbackURL: getEnv("GOOGLE_CALLBACK_URL"),
		GitHubClientID:    getEnv("GITHUB_CLIENT_ID"),
		GitHubSecret:      getEnv("GITHUB_CLIENT_SECRET"),
		GitHubCallbackURL: getEnv("GITHUB_CALLBACK_URL"),
		RedisURL:          getEnv("REDIS_URL"),
		SessionKey:        getEnv("SESSION_KEY"),
		SessionSecret:     getEnv("SESSION_SECRET"),
	}
}

func getEnv(key string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		panic("Environment variable " + key + " not found")
	}
	return value
}
