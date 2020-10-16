package handlers

import (
	"encoding/json"
	"net/http"
	"github.com/gonzalezlrjesus/API-Betting-Sports/models"
	u "github.com/gonzalezlrjesus/API-Betting-Sports/utils"

	"github.com/gorilla/mux"
)

// UpdateRacingComponents Update racing Components
var UpdateRacingComponents = func(w http.ResponseWriter, r *http.Request) {

	racingComponents := &models.RacingComponents{}
	vars := mux.Vars(r)
	idRacing := vars["idRacing"]

	err := json.NewDecoder(r.Body).Decode(racingComponents)
	if err != nil {
		u.Respond(w, u.Message(false, "Error while decoding request body"))
		return
	}

	resp := racingComponents.UpdateRacingComponents(idRacing)
	u.Respond(w, resp)
}

// GetRacingComponents Get racing components
var GetRacingComponents = func(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	idRacing := vars["idRacing"]

	data := models.GetRacingComponents(&idRacing)
	resp := u.Message(true, "success")
	resp["data"] = data
	u.Respond(w, resp)
}
