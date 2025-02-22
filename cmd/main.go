package main

// Main function
func main() {
	// Set up the server
	server := Server()
	// Run the server
	server.Run(":8080")
}
