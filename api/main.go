package main

import (
	"log"
	"net/http"
	"time"
)

type Config struct {
	Server struct {
		Address string
	}
}

func loadConfig() (*Config, error) {
	// For simplicity, we are hardcoding the config here.
	// In a real application, you might load this from a file or environment variables.
	cfg := &Config{}
	cfg.Server.Address = ":8080"
	return cfg, nil
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		log.Printf("Started %s %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
		log.Printf("Completed %s in %v", r.URL.Path, time.Since(start))
	})
}

func main() {
	// Load Config
	cfg, err := loadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// Initialize Router
	router := http.NewServeMux()

	// Add logging middleware
	loggedRouter := loggingMiddleware(router)

	// Start Server
	log.Printf("Starting server on %s...", cfg.Server.Address)
	if err := http.ListenAndServe(cfg.Server.Address, loggedRouter); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
