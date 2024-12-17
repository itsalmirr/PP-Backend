package config

import (
	"fmt"
	"os"

	"backend.com/go-backend/src/models"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		panic("Failed to load .env file!")
	}

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)
	database, err := gorm.Open(postgres.New(postgres.Config{
		DriverName: "pgx",
		DSN:        dsn,
	}), &gorm.Config{
		PrepareStmt: false,
	})
	if err != nil {
		panic("Failed to connect to database!")
	}

	// migrate models
	database.AutoMigrate(&models.Realtor{}, &models.User{})
	fmt.Println("Database connected!")
	DB = database // Assign the database connection to the global variable
}

func SessionStorage() redis.Store {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		panic("Failed to load .env file!")
	}
	store, _ := redis.NewStore(10, "tcp", os.Getenv("REDIS_URL"), "", []byte(os.Getenv("SESSION_SECRET")))

	store.Options(sessions.Options{
		Path:     "/api/v1/users/me",
		MaxAge:   3 * 24 * 60 * 60,
		HttpOnly: true,
		Secure:   true,
		SameSite: 3,
	})
	return store
}
