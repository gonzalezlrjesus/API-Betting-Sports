package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gonzalezlrjesus/API-Betting-Sports/models"
	u "github.com/gonzalezlrjesus/API-Betting-Sports/utils"
)

// CreateAdmin user
var CreateAdmin = func(w http.ResponseWriter, r *http.Request) {

	admin := &models.Admin{}
	err := json.NewDecoder(r.Body).Decode(admin) //decode the request body into struct and failed if any error occur
	if err != nil {
		u.Respond(w, u.Message(false, "Invalid request"), 400)
		return
	}

	resp := admin.Create() //Create admin
	u.Respond(w, resp, 201)
}

// Authenticate var export
var Authenticate = func(w http.ResponseWriter, r *http.Request) {

	admin := &models.Admin{}

	err := json.NewDecoder(r.Body).Decode(admin) //decode the request body into struct and failed if any error occur
	if err != nil {
		u.Respond(w, u.Message(false, "Invalid request"), 400)
		return
	}

	resp := models.Login(admin.Email, admin.Password)
	u.Respond(w, resp, 200)
}

// GetAdminFor list system admin
var GetAdminFor = func(w http.ResponseWriter, r *http.Request) {

	data := models.GetAdmins()
	resp := u.Message(true, "success")
	resp["data"] = data
	u.Respond(w, resp, 200)
}
