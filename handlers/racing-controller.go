package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gonzalezlrjesus/API-Betting-Sports/models"
	u "github.com/gonzalezlrjesus/API-Betting-Sports/utils"

	"github.com/gorilla/mux"
)

// CreateRacing Racing
var CreateRacing = func(w http.ResponseWriter, r *http.Request) {
	newsList := make([]models.Racing, 0)
	vars := mux.Vars(r)
	idEvent := vars["idEvent"]

	err := json.NewDecoder(r.Body).Decode(&newsList)
	if err != nil {
		u.Respond(w, u.Message(false, "Error while decoding request body"), 400)
		return
	}

	tempUint64, _ := strconv.ParseUint(idEvent, 10, 32)
	resp := models.CreateRacingModel(newsList, uint(tempUint64))
	u.Respond(w, resp, 201)
}

// UpdateRacing Racing
var UpdateRacing = func(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idEvent := vars["idEvent"]
	tempUint64, _ := strconv.ParseUint(idEvent, 10, 32)

	newsList := make([]models.Racing, 0)
	err := json.NewDecoder(r.Body).Decode(&newsList)
	if err != nil {
		u.Respond(w, u.Message(false, "Error while decoding request body"), 400)
		return
	}

	data := models.UpdateRacingModel(newsList, uint(tempUint64))
	resp := u.Message(true, "Success")
	resp["data"] = data
	u.Respond(w, resp, 200)
}

// DeleteRacing Racing
var DeleteRacing = func(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idEvent := vars["idEvent"]
	idRacing := vars["idRacing"]

	data := models.DeleteRacing(&idEvent, &idRacing)
	resp := u.Message(true, strconv.FormatBool(data))
	u.Respond(w, resp, 200)
}

// GetRacingsFor list Racings
var GetRacingsFor = func(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idEvent := vars["idEvent"]

	data := models.GetRacings(&idEvent)
	resp := u.Message(true, "success")
	resp["data"] = data
	u.Respond(w, resp, 200)
}

// GetSpecificRacing find and show Racing
var GetSpecificRacing = func(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idRacing := vars["idRacing"]

	data := models.GetOneRacing(&idRacing)
	resp := u.Message(true, "GetSpecificRacing Success")
	resp["data"] = data
	u.Respond(w, resp, 200)
}

// GetSpecificRacingWithEvent find and show Racing within event
var GetSpecificRacingWithEvent = func(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idEvent := vars["idEvent"]
	idRacing := vars["idRacing"]

	data := models.FindRacingWithinEvent(&idEvent, &idRacing)
	resp := u.Message(true, "GetSpecificRacing Success")
	resp["data"] = data
	u.Respond(w, resp, 200)
}

// RepartirGanancias repartir ganancias al ganador
var RepartirGanancias = func(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idRacing := vars["idRacing"]
	idHorse := vars["idHorse"]

	tempUint64Horse, _ := strconv.ParseUint(idHorse, 10, 32)

	data := models.RepartirGanancias(idRacing, int(tempUint64Horse))
	resp := u.Message(true, "GetSpecificRacing Success")
	resp["data"] = data
	u.Respond(w, resp, 200)
}
