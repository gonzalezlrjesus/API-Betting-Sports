package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gonzalezlrjesus/API-Betting-Sports/models"
	u "github.com/gonzalezlrjesus/API-Betting-Sports/utils"

	"github.com/gorilla/mux"
)

// CreateHorse add new Horses to db
var CreateHorse = func(w http.ResponseWriter, r *http.Request) {

	newsList := make([]models.Horse, 0)
	err := json.NewDecoder(r.Body).Decode(&newsList)
	if err != nil {
		u.Respond(w, u.Message(false, "Error while decoding request body"), 400)
		return
	}

	resp := models.CreateHorses(newsList)
	u.Respond(w, resp, 201)
}

// UpdateHorse Horse data
var UpdateHorse = func(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idRacing := vars["idRacing"]

	horse := &models.Horse{}
	err := json.NewDecoder(r.Body).Decode(horse)
	if err != nil {
		u.Respond(w, u.Message(false, "Invalid request"), 400)
		return
	}

	data := horse.UpdateHorse(u.ConverStringToUint(idRacing), vars["idHorse"])
	resp := u.Message(true, "Success")
	resp["data"] = data
	u.Respond(w, resp, 200)
}

// DeleteHorse Delete Horse
var DeleteHorse = func(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idRacing := vars["idRacing"]
	idHorse := vars["idHorse"]
	data := models.DeleteHorse(u.ConverStringToUint(idRacing), u.ConverStringToUint(idHorse))
	resp := u.Message(true, strconv.FormatBool(data))
	u.Respond(w, resp, 200)
}

// GetHorses list Horses
var GetHorses = func(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	data := models.GetHorses(vars["idRacing"])
	resp := u.Message(true, "success")
	resp["data"] = data
	u.Respond(w, resp, 200)
}

// WithdrawHorseBefore Horse Before auction
var WithdrawHorseBefore = func(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idHorse := vars["idHorse"]
	idRacing := vars["idRacing"]

	data := models.WithdrawHorse(u.ConverStringToUint(idHorse), u.ConverStringToUint(idRacing))
	resp := u.Message(true, "success")
	resp["data"] = data
	u.Respond(w, resp, 200)
}
