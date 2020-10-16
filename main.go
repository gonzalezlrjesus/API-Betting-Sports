package main

import (
	"log"
	"net/http"

	"github.com/gonzalezlrjesus/API-Betting-Sports/config"
	"github.com/gonzalezlrjesus/API-Betting-Sports/routes"
)

func main() {

	port := config.GetPort()
	c := config.GetCORS()

	routes := routes.Routes()
	handler := c.Handler(routes)

	err := http.ListenAndServe(":"+port, handler) //Launch the app.
	if err != nil {
		log.Println(err)
	}
}
