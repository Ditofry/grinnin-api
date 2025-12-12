package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"grinnin-golang-api/models"
)

// HelloHandler handles both GET and POST requests to the /hello endpoint
func HelloHandler(w http.ResponseWriter, r *http.Request) {
	// Set content type to JSON
	w.Header().Set("Content-Type", "application/json")

	switch r.Method {
	case http.MethodGet:
		handleHelloGet(w)

	case http.MethodPost:
		handleHelloPost(w, r)

	default:
		// Method not allowed
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

// handleHelloGet handles GET requests to return "Hello World"
func handleHelloGet(w http.ResponseWriter) {
	response := models.HelloResponse{Message: "Hello World"}
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Error encoding response", http.StatusInternalServerError)
		return
	}
}

// handleHelloPost handles POST requests to return personalized greeting
func handleHelloPost(w http.ResponseWriter, r *http.Request) {
	var req models.HelloRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Validate that name is provided
	if req.Name == "" {
		http.Error(w, "Name is required", http.StatusBadRequest)
		return
	}

	response := models.HelloResponse{Message: fmt.Sprintf("Hello %s", req.Name)}
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Error encoding response", http.StatusInternalServerError)
		return
	}
}

