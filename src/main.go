package main

import (
	"backend.com/go-backend/src/cmd"
)

// Main function
func main() {
	// Set up the server
	server := cmd.Server()
	// Run the server
	server.Run(":8080")
}
