package main

import (
	"backend.com/go-backend/src/cmd"
)

func main() {
	server := cmd.Server()
	server.Run(":8080")
}
