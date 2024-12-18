package routes

import (
	"log"
	"net/http"
)

func RickHandler(w http.ResponseWriter, r *http.Request) {
	// Full HTML content with embedded YouTube video
	videoHTML := `
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
			title="YouTube video player" frameborder="0" allow="accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture; web-share" 
			referrerpolicy="strict-origin-when-cross-origin" allowfullscreen></iframe>
		</body>
		</html>
	`

	// Write the HTML response and check for errors
	_, err := w.Write([]byte(videoHTML))
	if err != nil {
		log.Printf("Error writing response: %v", err)
	}
}
