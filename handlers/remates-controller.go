package handlers

import (
	"net/http"

	"github.com/gonzalezlrjesus/API-Betting-Sports/models"
	u "github.com/gonzalezlrjesus/API-Betting-Sports/utils"

	"github.com/gorilla/mux"
)

// GetRematesFor by racingID
var GetRematesFor = func(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idRacing := vars["idRacing"]

	data := models.GetRemates(u.ConverStringToUint(idRacing))
	resp := u.Message(true, "success")
	resp["data"] = data
	u.Respond(w, resp, 200)
}
