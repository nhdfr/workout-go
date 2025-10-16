package api

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
)

func TestHandleWorkoutByID_OK(t *testing.T) {
	r := chi.NewRouter()
	wh := NewWorkoutHandler()
	r.Get("/workouts/{id}", wh.HandleWorkoutByID)

	req := httptest.NewRequest("GET", "/workouts/42", nil)
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	resp := w.Result()
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected status 200, got %d", resp.StatusCode)
	}

	// body check
	// we expect "this is the workout id 42\n"
	buf := w.Body.String()
	if buf != "this is the workout id 42\n" {
		t.Fatalf("unexpected body: %q", buf)
	}
}
