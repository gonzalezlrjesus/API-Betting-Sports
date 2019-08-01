package routes

import (
	w "test-golang/auth"
	"test-golang/handlers"

	"github.com/gorilla/mux"
)

// Routes all routes
func Routes() *mux.Router {
	router := mux.NewRouter()

	// -------------------Routes Admin---------------------------------------

	// Create new Admin
	router.HandleFunc("/api/admin", handlers.CreateAdmin).Methods("POST")
	// Create new Client
	router.HandleFunc("/api/admin/clients", handlers.CreateClient).Methods("POST")

	// To view all admin
	router.HandleFunc("/api/admin", handlers.GetAdminFor).Methods("GET")
	// To View all clients
	router.HandleFunc("/api/admin/clients", handlers.GetClientsFor).Methods("GET")

	// To a specific client
	router.HandleFunc("/api/admin/clients/{idClient}", handlers.GetSpecificClient).Methods("GET")

	// Update Client
	router.HandleFunc("/api/admin/clients/{idClient}", handlers.UpdateClient).Methods("PUT")

	// Update State Blocked or Enable to access Client
	router.HandleFunc("/api/admin/clients/{idClient}/state", handlers.StateClient).Methods("PUT")

	// Delete Cliente
	router.HandleFunc("/api/admin/clients/{idClient}", handlers.DeleteClient).Methods("DELETE")

	// Add a deposit to money in client account
	router.HandleFunc("/api/admin/clients/{idClient}/deposit", handlers.AddDeposit).Methods("POST")

	// To view all Deposit Client
	router.HandleFunc("/api/admin/clients/{idClient}/deposit", handlers.HistorialDeposits).Methods("GET")

	// To view Coin Client
	router.HandleFunc("/api/admin/clients/{idClient}/coin", handlers.CoinClient).Methods("GET")

	// Login Admin
	router.HandleFunc("/api/admin/login", handlers.Authenticate).Methods("POST")

	// -------------------Routes Client---------------------------------------

	// Login Client
	router.HandleFunc("/api/clients/login", handlers.AuthenticateClient).Methods("POST")

	// ------------------Middleware Auth Token JWT----------------------------
	router.Use(w.JwtAuthentication) //attach JWT auth middleware

	return router
}
