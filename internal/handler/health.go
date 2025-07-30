package handler

import (
	"encoding/json"
	"net/http"
)

// Health responds to health checks
func (a *App) Health(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
}
