package main

import (
	"backend.com/go-backend/internal/config"
	"backend.com/go-backend/internal/routers"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

// config struct for database connection
var configVars = config.LoadConfig()

func Server() *gin.Engine {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	config.ConnectDatabase(configVars)
	router := routers.SetupRouter(configVars)

	return router
}
