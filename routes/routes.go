package routes

import (
	"github.com/gonzalezlrjesus/API-Betting-Sports/handlers"
	"github.com/gonzalezlrjesus/API-Betting-Sports/models"

	w "github.com/gonzalezlrjesus/API-Betting-Sports/auth"

	"github.com/gorilla/mux"
)

// Routes all routes
func Routes() *mux.Router {
	go models.Manager.Start()
	router := mux.NewRouter()

	// -------------------Admin---------------------------------------

	// Create new Admin
	router.HandleFunc("/api/admin", handlers.CreateAdmin).Methods("POST")
	// Login Admin
	router.HandleFunc("/api/admin/login", handlers.Authenticate).Methods("POST")
	// To view all admin
	router.HandleFunc("/api/admin", handlers.GetAdminFor).Methods("GET")

	// -------------------Clients-------------------------------------

	// Create new Client from Admin APP
	router.HandleFunc("/api/admin/clients", handlers.CreateClient).Methods("POST")
	// Create new Client from Client APP
	router.HandleFunc("/api/admin/client/app", handlers.CreateClient).Methods("POST")
	// Login Client
	router.HandleFunc("/api/clients/login", handlers.AuthenticateClient).Methods("POST")
	// To View all clients
	router.HandleFunc("/api/admin/clients", handlers.GetClientsFor).Methods("GET")
	// To a specific client
	router.HandleFunc("/api/admin/clients/{idClient}", handlers.GetSpecificClient).Methods("GET")
	// Update Client from Admin APP
	router.HandleFunc("/api/admin/clients/{idClient}", handlers.UpdateClient).Methods("PUT")
	// Update State Blocked or Enable to access Client
	router.HandleFunc("/api/admin/clients/{idClient}/state", handlers.StateClient).Methods("PUT")
	// Delete Cliente
	router.HandleFunc("/api/admin/clients/{idClient}", handlers.DeleteClient).Methods("DELETE")

	// -------------------Event---------------------------------------

	// Create new Event
	router.HandleFunc("/api/admin/event", handlers.CreateEvent).Methods("POST")
	// To View all Events
	router.HandleFunc("/api/admin/event", handlers.GetEventsFor).Methods("GET")
	// To a specific event
	router.HandleFunc("/api/admin/event/{idEvent}", handlers.GetSpecificEvent).Methods("GET")
	// Update Event
	router.HandleFunc("/api/admin/event/{idEvent}", handlers.UpdateEvent).Methods("PUT")
	// Delete Event
	router.HandleFunc("/api/admin/event/{idEvent}", handlers.DeleteEvent).Methods("DELETE")

	// -------------------Racing--------------------------------------

	// Create new Racing
	router.HandleFunc("/api/admin/event/{idEvent}/racing", handlers.CreateRacing).Methods("POST")
	// To View all Racings
	router.HandleFunc("/api/admin/event/{idEvent}/racing", handlers.GetRacingsFor).Methods("GET")
	// To a specific Racing
	router.HandleFunc("/api/admin/racing/{idRacing}", handlers.GetSpecificRacing).Methods("GET")
	// To a specific Racing within Event
	router.HandleFunc("/api/admin/event/{idEvent}/racing/{idRacing}", handlers.GetSpecificRacingWithEvent).Methods("GET")
	// Repartir Ganancias
	router.HandleFunc("/api/admin/{idRacing}/{idHorse}", handlers.RepartirGanancias).Methods("GET")
	// Update Racing
	router.HandleFunc("/api/admin/event/{idEvent}/racing", handlers.UpdateRacing).Methods("PUT")
	// Delete Racing
	router.HandleFunc("/api/admin/event/{idEvent}/racing/{idRacing}", handlers.DeleteRacing).Methods("DELETE")

	// -------------------Horse---------------------------------------

	// Create new a Horse
	router.HandleFunc("/api/admin/racing/horse", handlers.CreateHorse).Methods("POST")
	// To View all Horses
	router.HandleFunc("/api/admin/racing/{idRacing}/horse", handlers.GetHorses).Methods("GET")
	//********************// Update Horse state retirar caballos*************************
	router.HandleFunc("/api/admin/racing/{idRacing}/horse/{idHorse}", handlers.WithdrawHorseBefore).Methods("PATCH")
	// Update Horse
	router.HandleFunc("/api/admin/racing/{idRacing}/horse/{idHorse}", handlers.UpdateHorse).Methods("PUT")
	// Delete a Horse
	router.HandleFunc("/api/admin/racing/{idRacing}/horse/{idHorse}", handlers.DeleteHorse).Methods("DELETE")

	// -------------------Withdrawal----------------------------------

	// Do a withdrawal to money in client account
	router.HandleFunc("/api/admin/clients/{clientidentificationcard}/withdrawal", handlers.Dowithdrawal).Methods("POST")
	// To view all withdrawal Client
	router.HandleFunc("/api/admin/clients/{clientidentificationcard}/withdrawal", handlers.HistorialWithdrawal).Methods("GET")

	// -------------------Coin----------------------------------------

	// To a specific Coin Client
	router.HandleFunc("/api/admin/clients/{idClient}/coin", handlers.CoinClient).Methods("GET")

	// -------------------Remates-------------------------------------

	// To a specific remates Racings
	router.HandleFunc("/api/admin/racing/{idRacing}/remates", handlers.GetRematesFor).Methods("GET")

	// -------------------Tablas--------------------------------------

	// To a specific Tablas Remates
	router.HandleFunc("/api/admin/{idRacing}/tablas", handlers.GetTablas).Methods("GET")

	// -------------------Deposit-------------------------------------

	// Add a deposit to money in client account
	router.HandleFunc("/api/admin/clients/{idClient}/deposit", handlers.AddDeposit).Methods("POST")
	// To view all Deposit Client
	router.HandleFunc("/api/admin/clients/{idClient}/deposit", handlers.HistorialDeposits).Methods("GET")

	// ------------------WEBSOCKET------------------------------------

	// websocket to auctions
	router.HandleFunc("/ws", handlers.WsPage)

	// ------------------Middleware Auth Token JWT--------------------

	//attach JWT auth middleware
	router.Use(w.JwtAuthentication)

	return router
}
