package tests

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/DevOpsApproachFall2024/Daniel-CaoKhanh-Devops-FinalProject/backend/routes"
)

func TestRickRoute(t *testing.T) {
	// Create a test request to the /rick route
	req := httptest.NewRequest("GET", "/rick", nil)
	rec := httptest.NewRecorder()

	// Call the RickHandler
	routes.RickHandler(rec, req)

	// Expected HTML response
	expectedHTML := `
    <!DOCTYPE html>
    <html lang="en">
    <head>
        <meta charset="UTF-8">
        <meta name="viewport" content="width=device-width, initial-scale=1.0">
        <title>YouTube Video</title>
    </head>
    <body>
        <h1>Embedded YouTube Video</h1>
        <iframe width="560" height="315" src="https://www.youtube.com/embed/BBJa32lCaaY?si=GMsiK3ZLo66i-F04"
                title="YouTube video player" frameborder="0"
                allow="accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture; web-share"
                referrerpolicy="strict-origin-when-cross-origin" allowfullscreen></iframe>
    </body>
    </html>`

	// Get the result from the recorder
	result := strings.TrimSpace(rec.Body.String())
	expected := strings.TrimSpace(expectedHTML)

	// Compare the expected and actual response body
	if result != expected {
		t.Errorf("Expected response body to be %v but got %v", expected, result)
	}

	// Check that the status code is 200 OK
	if rec.Result().StatusCode != http.StatusOK {
		t.Errorf("Expected status code 200, but got %v", rec.Result().StatusCode)
	}
}
