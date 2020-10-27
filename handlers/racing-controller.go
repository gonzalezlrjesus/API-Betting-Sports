package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gonzalezlrjesus/API-Betting-Sports/models"
	u "github.com/gonzalezlrjesus/API-Betting-Sports/utils"

	"github.com/gorilla/mux"
)

// CreateRacing .
var CreateRacing = func(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idEvent := vars["idEvent"]

	raceList := make([]models.Racing, 0)
	err := json.NewDecoder(r.Body).Decode(&raceList)
	if err != nil {
		u.Respond(w, u.Message(false, "Error while decoding request body"), 400)
		return
	}

	resp := models.CreateRaces(raceList, u.ConverStringToUint(idEvent))
	u.Respond(w, resp, 201)
}

// UpdateRacing .
var UpdateRacing = func(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idEvent := vars["idEvent"]

	raceList := make([]models.Racing, 0)
	err := json.NewDecoder(r.Body).Decode(&raceList)
	if err != nil {
		u.Respond(w, u.Message(false, "Error while decoding request body"), 400)
		return
	}

	data := models.UpdateRaces(raceList, u.ConverStringToUint(idEvent))
	resp := u.Message(true, "Success")
	resp["data"] = data
	u.Respond(w, resp, 200)
}

// DeleteRacing .
var DeleteRacing = func(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idEvent := vars["idEvent"]
	idRacing := vars["idRacing"]

	data := models.DeleteRacing(u.ConverStringToUint(idEvent), u.ConverStringToUint(idRacing))
	resp := u.Message(true, strconv.FormatBool(data))
	u.Respond(w, resp, 200)
}

// GetRacingsFor .
var GetRacingsFor = func(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idEvent := vars["idEvent"]

	data := models.GetRacings(u.ConverStringToUint(idEvent))
	resp := u.Message(true, "success")
	resp["data"] = data
	u.Respond(w, resp, 200)
}

// GetSpecificRacing .
var GetSpecificRacing = func(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idRacing := vars["idRacing"]

	data := models.GetRace(u.ConverStringToUint(idRacing))
	resp := u.Message(true, "Success")
	resp["data"] = data
	u.Respond(w, resp, 200)
}

// GetSpecificRacingWithEvent find and show Racing within event
var GetSpecificRacingWithEvent = func(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idEvent := vars["idEvent"]
	idRacing := vars["idRacing"]

	data := models.FindRaceByEventID(u.ConverStringToUint(idEvent), u.ConverStringToUint(idRacing))
	resp := u.Message(true, "Success")
	resp["data"] = data
	u.Respond(w, resp, 200)
}

// RepartirGanancias .
var RepartirGanancias = func(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idRacing := vars["idRacing"]
	idHorse := vars["idHorse"]

	data := models.RepartirGanancias(u.ConverStringToUint(idRacing), u.ConverStringToInt(idHorse))
	resp := u.Message(true, "Success")
	resp["data"] = data
	u.Respond(w, resp, 200)
}
