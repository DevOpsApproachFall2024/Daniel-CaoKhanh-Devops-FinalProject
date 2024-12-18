package routes

import (
	"net/http"
)

func RickHandler(w http.ResponseWriter, r *http.Request) {
	// Serve the HTML for the embedded YouTube video
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
                title="YouTube video player" frameborder="0"
                allow="accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture; web-share"
                referrerpolicy="strict-origin-when-cross-origin" allowfullscreen></iframe>
    </body>
    </html>`

	// Write the HTML response
	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte(videoHTML))
}
