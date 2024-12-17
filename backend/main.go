package main

import (
	"log"
	"net/http"

	"github.com/DevOpsApproachFall2024/Daniel-CaoKhanh-Devops-FinalProject/backend/routes"
)

func main() {
	http.HandleFunc("/tony", routes.TonyHandler)

	log.Println("Server running on port 8080...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
