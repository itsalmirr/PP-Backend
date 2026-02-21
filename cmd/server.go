package main

import (
	"context"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"ppgroup.ppgroup.com/internal/config"
	"ppgroup.ppgroup.com/internal/routers"
	"ppgroup.ppgroup.com/internal/services"
)

func Server() (*gin.Engine, *config.Database, error) {
	// .env is optional in production (env vars set externally)
	_ = godotenv.Load()

	configVars, err := config.LoadConfig()
	if err != nil {
		return nil, nil, fmt.Errorf("loading config: %w", err)
	}

	ctx := context.Background()

	db, err := config.ConnectDatabase(ctx, configVars)
	if err != nil {
		return nil, nil, fmt.Errorf("connecting to database: %w", err)
	}

	if err := db.Migrate(ctx); err != nil {
		db.Close()
		return nil, nil, fmt.Errorf("running migrations: %w", err)
	}

	imageService, err := services.NewImageService(
		configVars.CloudinaryCloudName,
		configVars.CloudinaryAPIKey,
		configVars.CloudinaryAPISecret,
	)
	if err != nil {
		db.Close()
		return nil, nil, fmt.Errorf("initializing image service: %w", err)
	}

	router, err := routers.SetupRouter(configVars, db, imageService)
	if err != nil {
		db.Close()
		return nil, nil, fmt.Errorf("setting up router: %w", err)
	}

	return router, db, nil
}
