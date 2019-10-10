package main

import (
	"API-Betting-Sports/routes"
	"fmt"

	"net/http"
	"os"

	"github.com/rs/cors"
)

func main() {

	routes := routes.Routes()

	port := os.Getenv("PORT") //Get port from .env file, we did not specify any port so this should return an empty string when tested locally
	if port == "" {
		port = "8000" //localhost
	}
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:4200", "*", "ws://localhost:4200"},
		AllowCredentials: true,
		AllowedMethods:   []string{"POST", "GET", "OPTIONS", "PUT", "DELETE"},
		AllowedHeaders:   []string{"Accept", "Content-Type", "Origin", "X-Requested-With", "Content-Length", "Accept-Encoding", "Cache-Control", "Authorization"},
		// Enable Debugging for testing, consider disabling in production
		Debug: true,
	})

	handler := c.Handler(routes)

	err := http.ListenAndServe(":"+port, handler) //Launch the app, visit localhost:8000/api
	if err != nil {
		fmt.Print(err)
	}
}
