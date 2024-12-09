package main

import (
	"backend.com/go-backend/src/cmd"
	"backend.com/go-backend/src/routers"
)

func main() {
	cmd.ConnectDatabase()

	router := routers.SetupRouter()
	router.Run(":8080")
}
