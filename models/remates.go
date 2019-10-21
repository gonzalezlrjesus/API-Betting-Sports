package models

import (
	u "API-Betting-Sports/utils"
	"fmt"

	"github.com/jinzhu/gorm"
)

// Remates struct
type Remates struct {
	gorm.Model
	Idracing  string `json:"idracing"`
	Idhorse   int    `json:"idhorse"`
	Seudonimo string `json:"seudonimo"`
	Amount    int64  `json:"amount"`
	Horsename string `json:"horsename"`
}

// CreateRemates Remates db
func CreateRemates(idracing string, idhorse int, seudonimo string, amount int64, horsename string) map[string]interface{} {
	remateGanador := &Remates{Idracing: idracing, Idhorse: idhorse, Seudonimo: seudonimo, Amount: amount, Horsename: horsename}

	GetDB().Create(remateGanador)

	response := u.Message(true, "Remate added")
	response["remate"] = remateGanador
	return response
}

// GetRemates Remates db
func GetRemates(idracing *string) map[string]interface{} {

	remates := make([]*Remates, 0)

	err := GetDB().Table("remates").Where("idracing = ?", *idracing).Find(&remates).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}

	if len(remates) > 0 {
		response := u.Message(true, "Remates added")
		response["remates"] = remates
		return response
	}
	response := u.Message(true, "EMPTY")
	return response
}
