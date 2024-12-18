package routes

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func FibonacciHandler(w http.ResponseWriter, r *http.Request) {
	// Get the number from the URL
	numberStr := strings.TrimPrefix(r.URL.Path, "/fibonacci/")

	// Convert the string to an integer
	number, err := strconv.Atoi(numberStr)
	if err != nil || number < 0 {
		http.Error(w, "Invalid number", http.StatusBadRequest)
		return
	}

	// Generate the Fibonacci sequence
	var sequence []int
	a, b := 0, 1
	for i := 0; i < number; i++ {
		sequence = append(sequence, a)
		a, b = b, a+b
	}

	// Convert sequence to string
	var result string
	for _, num := range sequence {
		result += fmt.Sprintf("%d ", num)
	}

	// Write plain text response and check for errors
	_, err = w.Write([]byte(result))
	if err != nil {
		log.Printf("Error writing response: %v", err)
	}
}
