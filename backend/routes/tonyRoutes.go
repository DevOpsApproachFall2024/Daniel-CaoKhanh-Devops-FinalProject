package routes

import (
	"net/http"
)

func TonyHandler(w http.ResponseWriter, r *http.Request) {
	html := `
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

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	_, err := w.Write([]byte(html))
	if err != nil {
		http.Error(w, "Failed to write response", http.StatusInternalServerError)
	}
}
