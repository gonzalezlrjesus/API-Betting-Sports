package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gonzalezlrjesus/API-Betting-Sports/models"
	u "github.com/gonzalezlrjesus/API-Betting-Sports/utils"

	"github.com/gorilla/mux"
)

// AddDeposit to client
var AddDeposit = func(w http.ResponseWriter, r *http.Request) {

	deposit := &models.Deposit{}
	vars := mux.Vars(r)
	idClient := vars["idClient"]

	err := json.NewDecoder(r.Body).Decode(deposit)
	if err != nil {
		u.Respond(w, u.Message(false, "Error while decoding request body"), 400)
		return
	}
	deposit.Clientidentificationcard = idClient
	resp := deposit.AddDepositClient()
	u.Respond(w, resp, 200)
}

// HistorialDeposits client
var HistorialDeposits = func(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	idClient := vars["idClient"]

	data := models.GetAllDepositsClient(&idClient)
	resp := u.Message(true, "success")
	resp["data"] = data
	u.Respond(w, resp, 200)
}
