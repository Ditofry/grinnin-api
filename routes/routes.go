package routes

import (
	"net/http"

	"grinnin-golang-api/handlers"
)

// SetupRoutes configures all application routes
func SetupRoutes() {
	http.HandleFunc("/hello", handlers.HelloHandler)
}

