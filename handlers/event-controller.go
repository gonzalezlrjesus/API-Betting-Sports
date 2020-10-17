package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gonzalezlrjesus/API-Betting-Sports/models"
	u "github.com/gonzalezlrjesus/API-Betting-Sports/utils"

	"github.com/gorilla/mux"
)

// CreateEvent Event
var CreateEvent = func(w http.ResponseWriter, r *http.Request) {

	event := &models.Event{}
	err := json.NewDecoder(r.Body).Decode(event) //decode the request body into struct and failed if any error occur
	fmt.Println("err:", err)
	if err != nil {
		u.Respond(w, u.Message(false, "Invalid request"), 400)
		return
	}

	resp := event.CreateEvent() //Create client
	u.Respond(w, resp, 201)
}

// UpdateEvent find client
var UpdateEvent = func(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idEvent := vars["idEvent"]
	event := &models.Event{}

	err := json.NewDecoder(r.Body).Decode(event) //decode the request body into struct and failed if any error occur
	if err != nil {
		u.Respond(w, u.Message(false, "Invalid request"), 400)
		return
	}

	data := event.UpdateEvent(&idEvent)
	resp := u.Message(true, "Success")
	resp["data"] = data
	u.Respond(w, resp, 200)
}

// DeleteEvent Event
var DeleteEvent = func(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idEvent := vars["idEvent"]

	data := models.DeleteEvent(&idEvent)
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

	data := models.GetOneEvent(&idEvent)
	resp := u.Message(true, "GetSpecificEvent Success")
	resp["data"] = data
	u.Respond(w, resp, 200)
}
