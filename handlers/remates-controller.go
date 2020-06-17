package handlers

import (
	"API-Betting-Sports/models"
	u "API-Betting-Sports/utils"
	"net/http"

	"github.com/gorilla/mux"
)

// RematesRacing racings
var RematesRacing = func(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	idRacing := vars["idRacing"]

	data := models.GetRemates(&idRacing)
	resp := u.Message(true, "success")
	resp["data"] = data
	u.Respond(w, resp)
}

// GetRematesFor list Racings
var GetRematesFor = func(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idRacing := vars["idRacing"]

	data := models.GetRemates(&idRacing)
	resp := u.Message(true, "success")
	resp["data"] = data
	u.Respond(w, resp)
}
