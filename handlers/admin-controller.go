package handlers

import (
	"encoding/json"
	"net/http"

	"API-Betting-Sports/models"
	u "API-Betting-Sports/utils"
)

// CreateAdmin user
var CreateAdmin = func(w http.ResponseWriter, r *http.Request) {

	admin := &models.Admin{}
	err := json.NewDecoder(r.Body).Decode(admin) //decode the request body into struct and failed if any error occur
	
	if err != nil {
		u.Respond(w, u.Message(false, "Invalid request"))
		return
	}

	resp := admin.Create() //Create admin
	u.Respond(w, resp)
}

// Authenticate var export
var Authenticate = func(w http.ResponseWriter, r *http.Request) {

	admin := &models.Admin{}

	err := json.NewDecoder(r.Body).Decode(admin) //decode the request body into struct and failed if any error occur
	if err != nil {
		u.Respond(w, u.Message(false, "Invalid request"))
		return
	}

	resp := models.Login(admin.Email, admin.Password)
	u.Respond(w, resp)
}

// GetAdminFor list system admin
var GetAdminFor = func(w http.ResponseWriter, r *http.Request) {

	data := models.GetAdmins()
	resp := u.Message(true, "success")
	resp["data"] = data
	u.Respond(w, resp)
}
