package main

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"ppgroup.m0chi.com/internal/config"
	"ppgroup.m0chi.com/internal/routers"
	"ppgroup.m0chi.com/internal/services"
)

func Server() *gin.Engine {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	configVars := config.LoadConfig()
	ctx := context.Background()

	// Connect to database
	db, err := config.ConnectDatabase(ctx, configVars)
	if err != nil {
		panic("failed to connect database" + err.Error())
	}

	// Run migrations
	if err := db.Migrate(ctx); err != nil {
		panic("failed to run migrations: " + err.Error())
	}

	// Initialize ImageService with Cloudinary
	imageService := services.NewImageService(
		configVars.CloudinaryCloudName,
		configVars.CloudinaryAPIKey,
		configVars.CloudinaryAPISecret,
	)

	// Setup router
	router := routers.SetupRouter(configVars, db, imageService)

	return router
}
