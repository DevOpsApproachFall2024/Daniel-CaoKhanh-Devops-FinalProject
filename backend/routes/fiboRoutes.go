package routes

import (
	"fmt"
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

	// Write plain text response
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte(result))
}
