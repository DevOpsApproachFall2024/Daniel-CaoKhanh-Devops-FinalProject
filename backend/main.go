package main

import (
	"log"
	"net/http"

	"github.com/DevOpsApproachFall2024/Daniel-CaoKhanh-Devops-FinalProject/backend/routes"
	"github.com/gorilla/handlers"
)

func main() {
	// Create a new router
	mux := http.NewServeMux()

	// Define your routes
	mux.HandleFunc("/tony", routes.TonyHandler)
	mux.HandleFunc("/fibonacci/", routes.FibonacciHandler)
	mux.HandleFunc("/rick", routes.RickHandler)

	// Set up CORS options
	corsOptions := handlers.AllowedOrigins([]string{"*"})
	corsHeaders := handlers.AllowedHeaders([]string{"Content-Type", "X-Requested-With"})
	corsMethods := handlers.AllowedMethods([]string{"GET", "POST", "OPTIONS"})

	// Log and start the server with CORS enabled
	log.Println("Server running on port 8080...")
	err := http.ListenAndServe(":8080", handlers.CORS(corsOptions, corsHeaders, corsMethods)(mux))
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
