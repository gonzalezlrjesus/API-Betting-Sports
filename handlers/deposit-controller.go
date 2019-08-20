package handlers

import (
	"encoding/json"
	"net/http"
	"API-Betting-Sports/models"
	u "API-Betting-Sports/utils"

	"github.com/gorilla/mux"
)

// AddDeposit to client
var AddDeposit = func(w http.ResponseWriter, r *http.Request) {

	deposit := &models.Deposit{}
	vars := mux.Vars(r)
	idClient := vars["idClient"]

	err := json.NewDecoder(r.Body).Decode(deposit)
	if err != nil {
		u.Respond(w, u.Message(false, "Error while decoding request body"))
		return
	}
	deposit.Clientidentificationcard = idClient
	resp := deposit.AddDepositClient()
	u.Respond(w, resp)
}

// HistorialDeposits client
var HistorialDeposits = func(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	idClient := vars["idClient"]

	data := models.GetAllDepositsClient(&idClient)
	resp := u.Message(true, "success")
	resp["data"] = data
	u.Respond(w, resp)
}
