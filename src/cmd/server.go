package cmd

import (
	"backend.com/go-backend/src/config"
	"backend.com/go-backend/src/routers"
	"github.com/gin-gonic/gin"
)

func Server() *gin.Engine {
	config.ConnectDatabase()
	router := routers.SetupRouter()
	return router
}
