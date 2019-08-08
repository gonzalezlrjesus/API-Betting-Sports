package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"test-golang/models"
	u "test-golang/utils"

	"github.com/gorilla/mux"
)

// paramsID struct params
type paramsID struct {
	idEvent  string
	idRacing string
}

// CreateRacing Racing
var CreateRacing = func(w http.ResponseWriter, r *http.Request) {

	racing := &models.Racing{}
	vars := mux.Vars(r)
	idEvent := vars["idEvent"]

	err := json.NewDecoder(r.Body).Decode(racing)
	if err != nil {
		u.Respond(w, u.Message(false, "Error while decoding request body"))
		return
	}
	racing.Eventid = idEvent
	resp := racing.CreateRacing()
	u.Respond(w, resp)
}

// UpdateRacing Racing
var UpdateRacing = func(w http.ResponseWriter, r *http.Request) {
	params := &paramsID{}
	vars := mux.Vars(r)
	params.idEvent = vars["idEvent"]
	params.idRacing = vars["idRacing"]

	racing := &models.Racing{}

	err := json.NewDecoder(r.Body).Decode(racing) //decode the request body into struct and failed if any error occur
	if err != nil {
		u.Respond(w, u.Message(false, "Invalid request"))
		return
	}

	data := racing.UpdateRacing(&params.idEvent, &params.idRacing)
	resp := u.Message(true, "Success")
	resp["data"] = data
	u.Respond(w, resp)
}

// DeleteRacing Racing
var DeleteRacing = func(w http.ResponseWriter, r *http.Request) {
	params := &paramsID{}
	vars := mux.Vars(r)
	params.idEvent = vars["idEvent"]
	params.idRacing = vars["idRacing"]

	data := models.DeleteRacing(&params.idEvent, &params.idRacing)
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
	params := &paramsID{}
	vars := mux.Vars(r)
	params.idEvent = vars["idEvent"]
	params.idRacing = vars["idRacing"]

	data := models.GetOneRacing(&params.idEvent, &params.idRacing)
	resp := u.Message(true, "GetSpecificRacing Success")
	resp["data"] = data
	u.Respond(w, resp)
}
