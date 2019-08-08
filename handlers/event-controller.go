package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"test-golang/models"
	u "test-golang/utils"

	"github.com/gorilla/mux"
)

// CreateEvent Event
var CreateEvent = func(w http.ResponseWriter, r *http.Request) {

	event := &models.Event{}
	err := json.NewDecoder(r.Body).Decode(event) //decode the request body into struct and failed if any error occur
	fmt.Println("err:", err)
	if err != nil {
		u.Respond(w, u.Message(false, "Invalid request"))
		return
	}

	resp := event.CreateEvent() //Create client
	u.Respond(w, resp)
}

// UpdateEvent find client
var UpdateEvent = func(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idEvent := vars["idEvent"]
	event := &models.Event{}

	err := json.NewDecoder(r.Body).Decode(event) //decode the request body into struct and failed if any error occur
	if err != nil {
		u.Respond(w, u.Message(false, "Invalid request"))
		return
	}

	data := event.UpdateEvent(&idEvent)
	resp := u.Message(true, "Success")
	resp["data"] = data
	u.Respond(w, resp)
}

// DeleteEvent Event
var DeleteEvent = func(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idEvent := vars["idEvent"]

	data := models.DeleteEvent(&idEvent)
	resp := u.Message(true, strconv.FormatBool(data))
	u.Respond(w, resp)
}

// GetEventsFor list Events
var GetEventsFor = func(w http.ResponseWriter, r *http.Request) {

	data := models.GetEvents()
	resp := u.Message(true, "success")
	resp["data"] = data
	u.Respond(w, resp)
}

// GetSpecificEvent find and show Event
var GetSpecificEvent = func(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idEvent := vars["idEvent"]

	data := models.GetOneEvent(&idEvent)
	resp := u.Message(true, "GetSpecificEvent Success")
	resp["data"] = data
	u.Respond(w, resp)
}
