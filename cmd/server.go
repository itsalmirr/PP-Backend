package main

import (
	"backend.com/go-backend/internal/config"
	"backend.com/go-backend/internal/routers"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func Server() *gin.Engine {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	config.ConnectDatabase()
	router := routers.SetupRouter(config.SessionStorage())

	return router
}
