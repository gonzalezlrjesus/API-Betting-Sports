package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"test-golang/models"
	u "test-golang/utils"

	"github.com/gorilla/mux"
)

// CreateClient client
var CreateClient = func(w http.ResponseWriter, r *http.Request) {

	client := &models.Client{}
	err := json.NewDecoder(r.Body).Decode(client) //decode the request body into struct and failed if any error occur

	if err != nil {
		u.Respond(w, u.Message(false, "Invalid request"))
		return
	}

	resp := client.CreateClient() //Create client
	u.Respond(w, resp)
}

// AuthenticateClient client
var AuthenticateClient = func(w http.ResponseWriter, r *http.Request) {

	client := &models.Client{}

	err := json.NewDecoder(r.Body).Decode(client) //decode the request body into struct and failed if any error occur
	if err != nil {
		u.Respond(w, u.Message(false, "Invalid request"))
		return
	}

	resp := models.LoginClient(client.Identificationcard)
	u.Respond(w, resp)
}

// GetClientsFor list clients
var GetClientsFor = func(w http.ResponseWriter, r *http.Request) {

	data := models.GetClients()
	resp := u.Message(true, "success")
	resp["data"] = data
	u.Respond(w, resp)
}

// GetSpecificClient find and show client
var GetSpecificClient = func(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idClient := vars["idClient"]

	data, respString := models.GetClient(&idClient)
	resp := u.Message(true, respString)
	resp["data"] = data
	u.Respond(w, resp)
}

// UpdateClient find client
var UpdateClient = func(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idClient := vars["idClient"]
	client := &models.Client{}

	err := json.NewDecoder(r.Body).Decode(client) //decode the request body into struct and failed if any error occur
	if err != nil {
		u.Respond(w, u.Message(false, "Invalid request"))
		return
	}

	data := client.UpdateClient(&idClient)
	resp := u.Message(true, "Success")
	resp["data"] = data
	u.Respond(w, resp)
}

// StateClient find client
var StateClient = func(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idClient := vars["idClient"]
	setStateClient := &models.Client{}

	err := json.NewDecoder(r.Body).Decode(setStateClient) //decode the request body into struct and failed if any error occur
	if err != nil {
		u.Respond(w, u.Message(false, "Invalid request"))
		return
	}

	data := models.UpdateStateClient(&idClient, setStateClient)
	resp := u.Message(true, "Success")
	resp["data"] = data
	u.Respond(w, resp)
}

// DeleteClient client
var DeleteClient = func(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idClient := vars["idClient"]

	data := models.DeleteClient(&idClient)
	resp := u.Message(true, strconv.FormatBool(data))
	u.Respond(w, resp)
}
