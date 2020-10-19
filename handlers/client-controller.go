package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gonzalezlrjesus/API-Betting-Sports/models"
	u "github.com/gonzalezlrjesus/API-Betting-Sports/utils"

	"github.com/gorilla/mux"
)

// CreateClient .
var CreateClient = func(w http.ResponseWriter, r *http.Request) {

	client := &models.Client{}
	err := json.NewDecoder(r.Body).Decode(client)
	if err != nil {
		u.Respond(w, u.Message(false, "Invalid request"), 400)
		return
	}

	resp := client.CreateClient()
	u.Respond(w, resp, 201)
}

// AuthenticateClient .
var AuthenticateClient = func(w http.ResponseWriter, r *http.Request) {

	client := &models.Client{}
	err := json.NewDecoder(r.Body).Decode(client)
	if err != nil {
		u.Respond(w, u.Message(false, "Invalid request"), 400)
		return
	}

	resp := models.LoginClient(client.Password, client.Seudonimo)
	u.Respond(w, resp, 200)
}

// GetClientsFor list clients
var GetClientsFor = func(w http.ResponseWriter, r *http.Request) {

	data := models.GetClients()
	resp := u.Message(true, "success")
	resp["data"] = data

	u.Respond(w, resp, 200)
}

// GetSpecificClient .
var GetSpecificClient = func(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idClient := vars["idClient"]

	data := models.GetClient(&idClient)
	resp := u.Message(true, "success")
	resp["data"] = data
	u.Respond(w, resp, 200)
}

// UpdateClient .
var UpdateClient = func(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idClient := vars["idClient"]
	client := &models.Client{}

	err := json.NewDecoder(r.Body).Decode(client)
	if err != nil {
		u.Respond(w, u.Message(false, "Invalid request"), 400)
		return
	}

	data := client.UpdateClient(&idClient)
	resp := u.Message(true, "Success")
	resp["data"] = data
	u.Respond(w, resp, 200)
}

// StateClient .
var StateClient = func(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idClient := vars["idClient"]

	data := models.UpdateStateClient(&idClient)
	resp := u.Message(true, "Success")
	resp["data"] = data
	u.Respond(w, resp, 200)
}

// DeleteClient client
var DeleteClient = func(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idClient := vars["idClient"]

	data := models.DeleteClient(&idClient)
	resp := u.Message(true, strconv.FormatBool(data))
	u.Respond(w, resp, 200)
}
