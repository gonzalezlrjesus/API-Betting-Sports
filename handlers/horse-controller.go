package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"test-golang/models"
	u "test-golang/utils"

	"github.com/gorilla/mux"
)

// ParamsID struct params
type ParamsID struct {
	idRacing string
	idHorse  string
}

// CreateHorse add a new Horse to db
var CreateHorse = func(w http.ResponseWriter, r *http.Request) {

	horse := &models.Horse{}
	vars := mux.Vars(r)
	idRacing := vars["idRacing"]

	err := json.NewDecoder(r.Body).Decode(horse)
	if err != nil {
		u.Respond(w, u.Message(false, "Error while decoding request body"))
		return
	}
	tempUint64, _ := strconv.ParseUint(idRacing, 10, 32)
	horse.Racingid = uint(tempUint64)
	resp := horse.CreateHorse()
	u.Respond(w, resp)
}

// UpdateHorse Horse data
var UpdateHorse = func(w http.ResponseWriter, r *http.Request) {
	params := &ParamsID{}
	vars := mux.Vars(r)
	params.idRacing = vars["idRacing"]
	params.idHorse = vars["idHorse"]

	horse := &models.Horse{}

	err := json.NewDecoder(r.Body).Decode(horse) //decode the request body into struct and failed if any error occur
	if err != nil {
		u.Respond(w, u.Message(false, "Invalid request"))
		return
	}

	data := horse.UpdateHorse(&params.idRacing, &params.idHorse)
	resp := u.Message(true, "Success")
	resp["data"] = data
	u.Respond(w, resp)
}

// DeleteHorse Delete Horse
var DeleteHorse = func(w http.ResponseWriter, r *http.Request) {
	params := &ParamsID{}
	vars := mux.Vars(r)
	params.idRacing = vars["idRacing"]
	params.idHorse = vars["idHorse"]

	data := models.DeleteHorse(&params.idRacing, &params.idHorse)
	resp := u.Message(true, strconv.FormatBool(data))
	u.Respond(w, resp)
}

// GetHorses list Horses
var GetHorses = func(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idRacing := vars["idRacing"]

	data := models.GetHorses(&idRacing)
	resp := u.Message(true, "success")
	resp["data"] = data
	u.Respond(w, resp)
}
