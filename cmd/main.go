package main

// Main function
func main() {
	// Set up the server
	server := Server()
	// Run the server
	err := server.Run(":8080")
	if err != nil {
		println("Failed to run server: " + err.Error())
	}
}
