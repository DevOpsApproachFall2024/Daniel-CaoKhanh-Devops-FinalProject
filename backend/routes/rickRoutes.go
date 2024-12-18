package routes

import (
	"log"
	"net/http"
)

func RickHandler(w http.ResponseWriter, r *http.Request) {
	// HTML content with embedded YouTube video
	videoHTML := `
		<h1>Rick's YouTube Video</h1>
		<iframe width="560" height="315" src="https://www.youtube.com/embed/BBJa32lCaaY?si=GMsiK3ZLo66i-F04"
		title="YouTube video player" frameborder="0" allow="accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture; web-share" 
		referrerpolicy="strict-origin-when-cross-origin" allowfullscreen></iframe>
	`

	// Write response and check for errors
	_, err := w.Write([]byte(videoHTML))
	if err != nil {
		log.Printf("Error writing response: %v", err)
	}
}
