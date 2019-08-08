package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"test-golang/models"
	u "test-golang/utils"

	"github.com/gorilla/mux"
)

// AddAuctionNumber to Auction Number table in DB
var AddAuctionNumber = func(w http.ResponseWriter, r *http.Request) {

	AuctionNumber := &models.AuctionNumber{}
	vars := mux.Vars(r)
	idComponent := vars["idComponent"]

	err := json.NewDecoder(r.Body).Decode(AuctionNumber)
	if err != nil {
		u.Respond(w, u.Message(false, "Error while decoding request body"))
		return
	}
	tempUint64, _ := strconv.ParseUint(idComponent, 10, 32)
	AuctionNumber.Racingcomponentid = uint(tempUint64)
	resp := AuctionNumber.AddAuctionNumber()
	u.Respond(w, resp)
}

// GetAuctionNumbers client
var GetAuctionNumbers = func(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	idComponent := vars["idComponent"]

	data := models.GetAuctionNumbers(&idComponent)
	resp := u.Message(true, "success")
	resp["data"] = data
	u.Respond(w, resp)
}

// DeleteAuctionNumber delete a auction Number
var DeleteAuctionNumber = func(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idComponent := vars["idComponent"]
	idauctionnumber := vars["idauctionnumber"]

	data := models.DeleteAuctionNumber(&idComponent, &idauctionnumber)
	resp := u.Message(true, strconv.FormatBool(data))
	u.Respond(w, resp)
}
