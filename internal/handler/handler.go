package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Home handles root requests and returns a simple JSON response.
func (a *App) Home(w http.ResponseWriter, r *http.Request) {
	msg := fmt.Sprintf("Welcome to the Dyad Go Template (%s)", a.Config.Environment)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": msg})
}
