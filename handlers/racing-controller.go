package handlers

import (
	"API-Betting-Sports/models"
	u "API-Betting-Sports/utils"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// CreateRacing Racing
var CreateRacing = func(w http.ResponseWriter, r *http.Request) {
	newsList := make([]models.Racing, 0)
	err := json.NewDecoder(r.Body).Decode(&newsList)
	if err != nil {
		u.Respond(w, u.Message(false, "Error while decoding request body"))
		return
	}
	vars := mux.Vars(r)
	idEvent := vars["idEvent"]
	tempUint64, _ := strconv.ParseUint(idEvent, 10, 32)
	resp := models.CreateRacingModel(newsList, uint(tempUint64))
	u.Respond(w, resp)
}

// UpdateRacing Racing
var UpdateRacing = func(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idEvent := vars["idEvent"]
	tempUint64, _ := strconv.ParseUint(idEvent, 10, 32)

	newsList := make([]models.Racing, 0)
	err := json.NewDecoder(r.Body).Decode(&newsList)
	if err != nil {
		u.Respond(w, u.Message(false, "Error while decoding request body"))
		return
	}

	data := models.UpdateRacingModel(newsList, uint(tempUint64))
	resp := u.Message(true, "Success")
	resp["data"] = data
	u.Respond(w, resp)
}

// DeleteRacing Racing
var DeleteRacing = func(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idEvent := vars["idEvent"]
	idRacing := vars["idRacing"]

	data := models.DeleteRacing(&idEvent, &idRacing)
	resp := u.Message(true, strconv.FormatBool(data))
	u.Respond(w, resp)
}

// GetRacingsFor list Racings
var GetRacingsFor = func(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idEvent := vars["idEvent"]

	data := models.GetRacings(&idEvent)
	resp := u.Message(true, "success")
	resp["data"] = data
	u.Respond(w, resp)
}

// GetSpecificRacing find and show Racing
var GetSpecificRacing = func(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idEvent := vars["idEvent"]
	idRacing := vars["idRacing"]

	data := models.GetOneRacing(&idEvent, &idRacing)
	resp := u.Message(true, "GetSpecificRacing Success")
	resp["data"] = data
	u.Respond(w, resp)
}
