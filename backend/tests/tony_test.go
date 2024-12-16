package tests

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/DevOpsApproachFall2024/Daniel-CaoKhanh-Devops-FinalProject/backend/routes"
)

func TestTonyRoute(t *testing.T) {
	// Create a new request for the /tony route
	req, err := http.NewRequest("GET", "/tony", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a ResponseRecorder to record the response
	rr := httptest.NewRecorder()

	// Call the TonyHandler
	handler := http.HandlerFunc(routes.TonyHandler)
	handler.ServeHTTP(rr, req)

	// Check that the status code is 200 OK
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Expected status code %d but got %d", http.StatusOK, status)
	}

	// Define the expected response body
	expected := `<!DOCTYPE html>
<html lang="en">
<head>
	<meta charset="UTF-8">
	<meta name="viewport" content="width=device-width, initial-scale=1.0">
	<title>YouTube Video</title>
</head>
<body>
	<h1>Embedded YouTube Video</h1>
	<iframe width="560" height="315" src="https://www.youtube.com/embed/aAkMkVFwAoo?si=1gcu44ZcUckuSuh6" title="YouTube video player" frameborder="0" allow="accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture; web-share" referrerpolicy="strict-origin-when-cross-origin" allowfullscreen></iframe>
</body>
</html>`

	// Check that the response body is what we expect
	if rr.Body.String() != expected {
		t.Errorf("Expected response body to be %v but got %v", expected, rr.Body.String())
	}
}
