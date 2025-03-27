package main

// Main function
func main() {
	// Set up the server
	server := Server()
	// Run the server
	server.Run(":8080")
	// go func() {
	// 	if err := server.Run(":8080"); err != nil {
	// 		log.Fatalf("failed to run server")
	// 	}
	// }()
}
