package main

import (
	"backend.com/go-backend/src/models"
	"backend.com/go-backend/src/routers"
)

func main() {
	models.ConnectDatabase()

	router := routers.SetupRouter()
	router.Run(":8080")
}
