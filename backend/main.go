package main

import (
	"log"
	"net/http"

	"github.com/DevOpsApproachFall2024/Daniel-CaoKhanh-Devops-FinalProject/backend/routes"
)

func main() {
	// Define the /tony route
	http.HandleFunc("/tony", routes.TonyHandler)

	// Start the server on port 8080
	log.Println("Starting server on :8080...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
