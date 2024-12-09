package main

import (
	"backend.com/go-backend/src/routers"
)

func main() {
	router := routers.SetupRouter()
	router.Run(":8080")
}
