package handlers

import (
	"API-Betting-Sports/models"
	"net/http"

	"github.com/gorilla/websocket"
	uuid "github.com/satori/go.uuid"
)

// WsPage func
func WsPage(res http.ResponseWriter, req *http.Request) {
	conn, error := (&websocket.Upgrader{CheckOrigin: func(r *http.Request) bool { return true }}).Upgrade(res, req, nil)

	if error != nil {
		http.NotFound(res, req)
		return
	}

	client := &models.Clientmodel{Id: uuid.Must(uuid.NewV4()).String(), Socket: conn, Send: make(chan []byte)}
	models.Manager.Register <- client

	go client.Read()
	go client.Write()
}
