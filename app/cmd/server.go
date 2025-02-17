package cmd

import (
	"backend.com/go-backend/app/config"
	"backend.com/go-backend/app/routers"
	"github.com/gin-gonic/gin"
)

func Server() *gin.Engine {
	config.ConnectDatabase()
	router := routers.SetupRouter(config.SessionStorage())

	return router
}
