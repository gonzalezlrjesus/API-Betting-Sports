package config

import (
	"os"

	"github.com/rs/cors"
)

// GetPort .
func GetPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000" //localhost
	}
	return port
}

// GetCORS .
func GetCORS() *cors.Cors {
	return cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:4200", "*", "ws://localhost:4200"},
		AllowCredentials: true,
		AllowedMethods:   []string{"POST", "GET", "OPTIONS", "PUT", "DELETE"},
		AllowedHeaders:   []string{"Accept", "Content-Type", "Origin", "X-Requested-With", "Content-Length", "Accept-Encoding", "Cache-Control", "Authorization"},
		// Enable Debugging for testing, consider disabling in production
		Debug: false,
	})
}
