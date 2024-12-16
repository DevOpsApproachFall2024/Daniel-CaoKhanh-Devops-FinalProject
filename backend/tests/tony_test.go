package tests

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/DevOpsApproachFall2024/Daniel-CaoKhanh-Devops-FinalProject/backend/routes"
)

func TestTonyRoute(t *testing.T) {
	req := httptest.NewRequest("GET", "/tony", nil)
	rec := httptest.NewRecorder()

	routes.TonyHandler(rec, req)

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
        <iframe width="560" height="315" src="https://www.youtube.com/embed/aAkMkVFwAoo?si=1gcu44ZcUckuSuh6"
                title="YouTube video player" frameborder="0"
                allow="accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture; web-share"
                referrerpolicy="strict-origin-when-cross-origin" allowfullscreen></iframe>
    </body>
    </html>`

	result := strings.TrimSpace(rec.Body.String())
	expected := strings.TrimSpace(expectedHTML)

	if result != expected {
		t.Errorf("Expected response body to be %v but got %v", expected, result)
	}

	if rec.Result().StatusCode != http.StatusOK {
		t.Errorf("Expected status code 200, but got %v", rec.Result().StatusCode)
	}
}
