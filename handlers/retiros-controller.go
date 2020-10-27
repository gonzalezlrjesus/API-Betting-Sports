package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gonzalezlrjesus/API-Betting-Sports/models"
	u "github.com/gonzalezlrjesus/API-Betting-Sports/utils"

	"github.com/gorilla/mux"
)

// Dowithdrawal from client coin
var Dowithdrawal = func(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	clientIDCard := vars["clientidentificationcard"]

	retiro := &models.Retiro{}
	err := json.NewDecoder(r.Body).Decode(retiro)
	if err != nil {
		u.Respond(w, u.Message(false, "Error while decoding request body"), 400)
		return
	}

	resp := retiro.Dowithdrawal(clientIDCard)
	u.Respond(w, resp, 200)
}

// HistorialWithdrawal from client account
var HistorialWithdrawal = func(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	clientIDCard := vars["clientidentificationcard"]

	data := models.GetAllWithdrawal(&clientIDCard)
	resp := u.Message(true, "success")
	resp["data"] = data
	u.Respond(w, resp, 200)
}
