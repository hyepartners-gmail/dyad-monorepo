package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/hyepartners-gmail/dyad-go-template/config"
	"github.com/hyepartners-gmail/dyad-go-template/internal/handler"
	"github.com/hyepartners-gmail/dyad-go-template/internal/middleware"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	app := &handler.App{Config: cfg}

	mux := http.NewServeMux()
	mux.HandleFunc("/", app.Home)
	mux.HandleFunc("/health", app.Health)

	port := cfg.Port
	if port == "" {
		port = "8080"
	}
	addr := fmt.Sprintf(":%s", port)

	log.Printf("Environment: %s | Listening on %s\n", cfg.Environment, addr)

	// Wrap with logging + CORS middleware
	if err := http.ListenAndServe(addr, middleware.WithCommonMiddleware(mux)); err != nil {
		log.Fatalf("Server error: %v", err)
	}
}
