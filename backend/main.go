package main

import (
	"log"
	"net/http"

	"github.com/DevOpsApproachFall2024/Daniel-CaoKhanh-Devops-FinalProject/backend/routes"
)

func main() {
	// route for Tony's video
	http.HandleFunc("/tony", routes.TonyHandler)
	// route for Fibonacci sequence
	http.HandleFunc("/fibonacci/", routes.FibonacciHandler)
	// route to get rick rolled
	http.HandleFunc("/rick", routes.RickHandler)

	// log and start the server
	log.Println("Server running on port 8080...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
