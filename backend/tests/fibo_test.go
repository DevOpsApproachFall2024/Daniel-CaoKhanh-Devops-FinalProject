package tests

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/DevOpsApproachFall2024/Daniel-CaoKhanh-Devops-FinalProject/backend/routes"
)

func TestFibonacciHandler(t *testing.T) {
	tests := []struct {
		number     string
		expected   string
		statusCode int
	}{
		{
			number:     "5",
			expected:   "0 1 1 2 3 ",
			statusCode: http.StatusOK,
		},
		{
			number:     "8",
			expected:   "0 1 1 2 3 5 8 13 ",
			statusCode: http.StatusOK,
		},
		{
			number:     "0",
			expected:   "",
			statusCode: http.StatusOK,
		},
		{
			number:     "invalid",
			expected:   "Invalid number",
			statusCode: http.StatusBadRequest,
		},
	}

	for _, tt := range tests {
		t.Run(tt.number, func(t *testing.T) {
			req, err := http.NewRequest("GET", "/fibonacci/"+tt.number, nil)
			if err != nil {
				t.Fatal(err)
			}

			// Create a ResponseRecorder to capture the response
			rr := httptest.NewRecorder()

			// Call the FibonacciHandler with the request and response recorder
			routes.FibonacciHandler(rr, req)

			// Check the status code
			if status := rr.Code; status != tt.statusCode {
				t.Errorf("expected status %v, got %v", tt.statusCode, status)
			}

			// Check the response body
			actualResponse := strings.TrimSpace(rr.Body.String()) // Trim extra spaces or newlines
			expectedResponse := strings.TrimSpace(tt.expected)    // Trim expected response

			if actualResponse != expectedResponse {
				t.Errorf("expected body %v, got %v", expectedResponse, actualResponse)
			}
		})
	}
}
