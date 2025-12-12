package main

import (
	"fmt"
	"log"
	"net/http"

	"grinnin-golang-api/routes"
)

func main() {
	// Setup all application routes
	routes.SetupRoutes()

	// Start the server on port 8080
	port := ":8080"
	fmt.Printf("Server starting on port %s\n", port)
	fmt.Println("Endpoints:")
	fmt.Println("  GET  http://localhost:8080/hello")
	fmt.Println("  POST http://localhost:8080/hello")

	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatal(err)
	}
}
