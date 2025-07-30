package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/hyepartners-gmail/dyad-go-template/config"
)

func TestHealth(t *testing.T) {
	app := &App{Config: &config.Config{Environment: "test"}}
	req := httptest.NewRequest(http.MethodGet, "/health", nil)
	rec := httptest.NewRecorder()

	app.Health(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", rec.Code)
	}
}
