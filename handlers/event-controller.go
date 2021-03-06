package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gonzalezlrjesus/API-Betting-Sports/models"
	u "github.com/gonzalezlrjesus/API-Betting-Sports/utils"

	"github.com/gorilla/mux"
)

// CreateEvent Event
var CreateEvent = func(w http.ResponseWriter, r *http.Request) {

	event := &models.Event{}
	err := json.NewDecoder(r.Body).Decode(event)
	if err != nil {
		u.Respond(w, u.Message(false, "Invalid request"), 400)
		return
	}

	resp := event.CreateEvent()
	u.Respond(w, resp, 201)
}

// UpdateEvent find client
var UpdateEvent = func(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idEvent := vars["idEvent"]

	event := &models.Event{}
	err := json.NewDecoder(r.Body).Decode(event)
	if err != nil {
		u.Respond(w, u.Message(false, "Invalid request"), 400)
		return
	}

	data := event.UpdateEvent(u.ConverStringToUint(idEvent))
	resp := u.Message(true, "Success")
	resp["data"] = data
	u.Respond(w, resp, 200)
}

// DeleteEvent Event
var DeleteEvent = func(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idEvent := vars["idEvent"]

	data := models.DeleteEvent(u.ConverStringToUint(idEvent))
	resp := u.Message(true, strconv.FormatBool(data))
	u.Respond(w, resp, 200)
}

// GetEventsFor list Events
var GetEventsFor = func(w http.ResponseWriter, r *http.Request) {

	data := models.GetEvents()
	resp := u.Message(true, "success")
	resp["data"] = data
	u.Respond(w, resp, 200)
}

// GetSpecificEvent find and show Event
var GetSpecificEvent = func(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idEvent := vars["idEvent"]

	data := models.GetOneEvent(u.ConverStringToUint(idEvent))
	resp := u.Message(true, "success")
	resp["data"] = data
	u.Respond(w, resp, 200)
}
