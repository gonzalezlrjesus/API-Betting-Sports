package handlers

import (
	"net/http"
	"test-golang/models"
	u "test-golang/utils"

	"github.com/gorilla/mux"
)

// CoinClient client
var CoinClient = func(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	idClient := vars["idClient"]

	data := models.GetCoinsClient(&idClient)
	resp := u.Message(true, "success")
	resp["data"] = data
	u.Respond(w, resp)
}
