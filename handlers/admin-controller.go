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
	err := json.NewDecoder(r.Body).Decode(admin)
	if err != nil {
		u.Respond(w, u.Message(false, "Invalid request"), 400)
		return
	}

	resp := admin.Create() //Create admin
	u.Respond(w, resp, 201)
}

// Authenticate admin user
var Authenticate = func(w http.ResponseWriter, r *http.Request) {

	admin := &models.Admin{}
	err := json.NewDecoder(r.Body).Decode(admin)
	if err != nil {
		u.Respond(w, u.Message(false, "Invalid request"), 400)
		return
	}

	resp := models.Login(admin.Email, admin.Password)
	u.Respond(w, resp, 200)
}

// GetAdminFor list admin user
var GetAdminFor = func(w http.ResponseWriter, r *http.Request) {

	data := models.GetAdmins()
	resp := u.Message(true, "success")
	resp["data"] = data

	u.Respond(w, resp, 200)
}
