package main

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"ppgroup.i0sys.com/internal/config"
	"ppgroup.i0sys.com/internal/routers"
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

	// Setup router
	router := routers.SetupRouter(configVars, db)

	return router
}
