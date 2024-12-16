package routes

import (
	"net/http"
)

// TonyHandler serves the embedded YouTube video iframe HTML
func TonyHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`<iframe width="560" height="315" src="https://www.youtube.com/embed/aAkMkVFwAoo?si=Cs3VG4r8d5Vw6l3P" title="YouTube video player" frameborder="0" allow="accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture; web-share" referrerpolicy="strict-origin-when-cross-origin" allowfullscreen></iframe>`))
}
