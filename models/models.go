package models

// HelloRequest represents the request body for POST requests
type HelloRequest struct {
	Name string `json:"name"`
}

// HelloResponse represents the response structure
type HelloResponse struct {
	Message string `json:"message"`
}

