package main

import (
	"fmt"
	"net/http"
	"os"
	"test-golang/routes"
)

func main() {
	routes := routes.Routes()

	port := os.Getenv("PORT") //Get port from .env file, we did not specify any port so this should return an empty string when tested locally
	if port == "" {
		port = "8000" //localhost
	}

	fmt.Println(port)

	err := http.ListenAndServe(":"+port, routes) //Launch the app, visit localhost:8000/api
	if err != nil {
		fmt.Print(err)
	}
}
