package routes

import (
	w "API-Betting-Sports/auth"
	"API-Betting-Sports/handlers"
	"API-Betting-Sports/models"

	"github.com/gorilla/mux"
)

// Routes all routes
func Routes() *mux.Router {
	go models.Manager.Start()
	router := mux.NewRouter()

	// -------------------Routes Admin---------------------------------------

	// Create new Admin
	router.HandleFunc("/api/admin", handlers.CreateAdmin).Methods("POST")
	// Create new Client
	router.HandleFunc("/api/admin/clients", handlers.CreateClient).Methods("POST")
	// Create new Event
	router.HandleFunc("/api/admin/event", handlers.CreateEvent).Methods("POST")
	// Create new Racing
	router.HandleFunc("/api/admin/event/{idEvent}/racing", handlers.CreateRacing).Methods("POST")
	// Create new a Horse
	router.HandleFunc("/api/admin/racing/{idRacing}/horse", handlers.CreateHorse).Methods("POST")

	// To view all admin
	router.HandleFunc("/api/admin", handlers.GetAdminFor).Methods("GET")
	// To View all clients
	router.HandleFunc("/api/admin/clients", handlers.GetClientsFor).Methods("GET")
	// To view all Deposit Client
	router.HandleFunc("/api/admin/clients/{idClient}/deposit", handlers.HistorialDeposits).Methods("GET")
	// To View all Events
	router.HandleFunc("/api/admin/event", handlers.GetEventsFor).Methods("GET")
	// To View all Racings
	router.HandleFunc("/api/admin/event/{idEvent}/racing", handlers.GetRacingsFor).Methods("GET")
	// To View all Auction Number of specify Racing Component
	router.HandleFunc("/api/admin/components/{idComponent}/auctionnumber", handlers.GetAuctionNumbers).Methods("GET")
	// To View all Horses
	router.HandleFunc("/api/admin/racing/{idRacing}/horse", handlers.GetHorses).Methods("GET")

	// To a specific client
	router.HandleFunc("/api/admin/clients/{idClient}", handlers.GetSpecificClient).Methods("GET")
	// To a specific Coin Client
	router.HandleFunc("/api/admin/clients/{idClient}/coin", handlers.CoinClient).Methods("GET")
	// To a specific event
	router.HandleFunc("/api/admin/event/{idEvent}", handlers.GetSpecificEvent).Methods("GET")
	// To a specific Racing
	router.HandleFunc("/api/admin/racing/{idRacing}", handlers.GetSpecificRacing).Methods("GET")
	// To a specific Racing within Event
	router.HandleFunc("/api/admin/racing/{idRacing}/{idEvent}", handlers.GetSpecificRacingWithEvent).Methods("GET")
	// To a specific Racing Components
	router.HandleFunc("/api/admin/racing/{idRacing}/components", handlers.GetRacingComponents).Methods("GET")
	// To a specific remates Racings
	router.HandleFunc("/api/admin/{idRacing}/remates", handlers.GetRematesFor).Methods("GET")
	// To a specific Tablas Remates
	router.HandleFunc("/api/admin/{idRacing}/tablas", handlers.GetTablas).Methods("GET")

	// Update Client
	router.HandleFunc("/api/admin/clients/{idClient}", handlers.UpdateClient).Methods("PUT")
	// Update State Blocked or Enable to access Client
	router.HandleFunc("/api/admin/clients/{idClient}/state", handlers.StateClient).Methods("PUT")
	// Update Event
	router.HandleFunc("/api/admin/event/{idEvent}", handlers.UpdateEvent).Methods("PUT")
	// Update Racing
	router.HandleFunc("/api/admin/event/{idEvent}/racing", handlers.UpdateRacing).Methods("PUT")
	// Update RacingComponents
	router.HandleFunc("/api/admin/racing/{idRacing}/components", handlers.UpdateRacingComponents).Methods("PUT")
	// Update Horse
	router.HandleFunc("/api/admin/racing/{idRacing}/horse/{idHorse}", handlers.UpdateHorse).Methods("PUT")

	// Delete Cliente
	router.HandleFunc("/api/admin/clients/{idClient}", handlers.DeleteClient).Methods("DELETE")
	// Delete Event
	router.HandleFunc("/api/admin/event/{idEvent}", handlers.DeleteEvent).Methods("DELETE")
	// Delete Racing
	router.HandleFunc("/api/admin/event/{idEvent}/racing/{idRacing}", handlers.DeleteRacing).Methods("DELETE")
	// Delete a auction number
	router.HandleFunc("/api/admin/event/components/{idComponent}/auctionnumber/{idauctionnumber}", handlers.DeleteAuctionNumber).Methods("DELETE")
	// Delete a Horse
	router.HandleFunc("/api/admin/racing/{idRacing}/horse/{idHorse}", handlers.DeleteHorse).Methods("DELETE")

	// Add a deposit to money in client account
	router.HandleFunc("/api/admin/clients/{idClient}/deposit", handlers.AddDeposit).Methods("POST")
	// Add a auction number
	router.HandleFunc("/api/admin/components/{idComponent}/auctionnumber", handlers.AddAuctionNumber).Methods("POST")

	// Login Admin
	router.HandleFunc("/api/admin/login", handlers.Authenticate).Methods("POST")

	// -------------------Routes Client---------------------------------------

	// Login Client
	router.HandleFunc("/api/clients/login", handlers.AuthenticateClient).Methods("POST")

	// ------------------ WEBSOCKET ----------------------------

	// websocket to auctions
	router.HandleFunc("/ws", handlers.WsPage)

	// ------------------Middleware Auth Token JWT----------------------------
	router.Use(w.JwtAuthentication) //attach JWT auth middleware

	return router
}
