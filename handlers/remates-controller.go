package handlers

import (
	"net/http"

	"github.com/gonzalezlrjesus/API-Betting-Sports/models"
	u "github.com/gonzalezlrjesus/API-Betting-Sports/utils"

	"github.com/gorilla/mux"
)

// RematesRacing racings
var RematesRacing = func(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	idRacing := vars["idRacing"]

	data := models.GetRemates(&idRacing)
	resp := u.Message(true, "success")
	resp["data"] = data
	u.Respond(w, resp, 200)
}

// GetRematesFor list Racings
var GetRematesFor = func(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idRacing := vars["idRacing"]

	data := models.GetRemates(&idRacing)
	resp := u.Message(true, "success")
	resp["data"] = data
	u.Respond(w, resp, 200)
}
