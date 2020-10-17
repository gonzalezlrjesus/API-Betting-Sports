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

// ParamsID struct params
type ParamsID struct {
	idRacing string
	idHorse  string
}

// CreateHorse add a new Horse to db
var CreateHorse = func(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	idRacing := vars["idRacing"]

	newsList := make([]models.Horse, 0)
	err := json.NewDecoder(r.Body).Decode(&newsList)
	if err != nil {
		u.Respond(w, u.Message(false, "Error while decoding request body"), 400)
		return
	}
	tempUint64, _ := strconv.ParseUint(idRacing, 10, 32)
	resp := models.CreateHorseModel(newsList, uint(tempUint64))
	u.Respond(w, resp, 201)
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
		u.Respond(w, u.Message(false, "Invalid request"), 400)
		return
	}

	data := horse.UpdateHorse(&params.idRacing, &params.idHorse)
	resp := u.Message(true, "Success")
	resp["data"] = data
	u.Respond(w, resp, 200)
}

// DeleteHorse Delete Horse
var DeleteHorse = func(w http.ResponseWriter, r *http.Request) {
	params := &ParamsID{}
	vars := mux.Vars(r)
	params.idRacing = vars["idRacing"]
	params.idHorse = vars["idHorse"]

	data := models.DeleteHorse(&params.idRacing, &params.idHorse)
	resp := u.Message(true, strconv.FormatBool(data))
	u.Respond(w, resp, 200)
}

// GetHorses list Horses
var GetHorses = func(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idRacing := vars["idRacing"]

	data := models.GetHorses(&idRacing)
	resp := u.Message(true, "success")
	resp["data"] = data
	u.Respond(w, resp, 200)
}

// WithdrawHorseBefore Withdraw Horse Before auction
var WithdrawHorseBefore = func(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idHorse := vars["idHorse"]
	idRacing := vars["idRacing"]
	fmt.Println(idRacing, idHorse)
	tempUint64, _ := strconv.ParseUint(idHorse, 10, 32)
	tempRacing, _ := strconv.ParseUint(idRacing, 10, 32)

	data := models.WithdrawHorse(uint(tempUint64), uint(tempRacing))
	resp := u.Message(true, "success")
	resp["data"] = data
	u.Respond(w, resp, 200)
}
