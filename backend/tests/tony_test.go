package tests

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/DevOpsApproachFall2024/Daniel-CaoKhanh-Devops-FinalProject/backend/routes"
)

func TestTonyRoute(t *testing.T) {
	req, err := http.NewRequest("GET", "/tony", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(routes.TonyHandler)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Expected status code %d but got %d", http.StatusOK, status)
	}
	expected := `<iframe width="560" height="315" src="https://www.youtube.com/embed/aAkMkVFwAoo?si=Cs3VG4r8d5Vw6l3P" title="YouTube video player" frameborder="0" allow="accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture; web-share" referrerpolicy="strict-origin-when-cross-origin" allowfullscreen></iframe>`
	if rr.Body.String() != expected {
		t.Errorf("Expected response body to be %s but got %s", expected, rr.Body.String())
	}
}
