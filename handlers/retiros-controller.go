package handlers

import (
	"encoding/json"
	"net/http"
	"github.com/gonzalezlrjesus/API-Betting-Sports/models"
	u "github.com/gonzalezlrjesus/API-Betting-Sports/utils"

	"github.com/gorilla/mux"
)

// Dowithdrawal to client
var Dowithdrawal = func(w http.ResponseWriter, r *http.Request) {

	retiro := &models.Retiro{}
	vars := mux.Vars(r)
	idClient := vars["idClient"]

	err := json.NewDecoder(r.Body).Decode(retiro)
	if err != nil {
		u.Respond(w, u.Message(false, "Error while decoding request body"))
		return
	}
	retiro.Clientidentificationcard = idClient
	resp := retiro.Dowithdrawal()
	u.Respond(w, resp)
}

// HistorialWithdrawal client
var HistorialWithdrawal = func(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	idClient := vars["idClient"]

	data := models.GetAllWithdrawal(&idClient)
	resp := u.Message(true, "success")
	resp["data"] = data
	u.Respond(w, resp)
}
