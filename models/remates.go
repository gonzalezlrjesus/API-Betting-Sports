package models

import (
	u "API-Betting-Sports/utils"

	"github.com/jinzhu/gorm"
)

// Remates struct
type Remates struct {
	gorm.Model
	Idracing  string   `json:"idracing"`
	Idhorse   int    `json:"idhorse"`
	Seudonimo string `json:"seudonimo"`
	Amount    int64  `json:"amount"`
}

// CreateRemates Remates db
func CreateRemates(idracing string, idhorse int, seudonimo string, amount int64) map[string]interface{} {
	remateGanador := &Remates{Idracing: idracing, Idhorse: idhorse, Seudonimo: seudonimo, Amount: amount}

	GetDB().Create(remateGanador)

	response := u.Message(true, "Remate added")
	response["remate"] = remateGanador
	return response
}

// GetRemates Remates db
func GetRemates(idracing *string) map[string]interface{} {
	// remateGanador := &Remates{Idracing: idracing}

	// GetDB().Create(remateGanador)

	response := u.Message(true, "Remate added")
	// response["remate"] = remateGanador
	return response
}
