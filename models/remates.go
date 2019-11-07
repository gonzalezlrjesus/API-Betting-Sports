package models

import (
	u "API-Betting-Sports/utils"
	"fmt"
	"time"
	"github.com/jinzhu/gorm"
)

// Cuando el proyecto este mas avanzado deberia llevar un control de la tabla que se esta jugando
// Todo eso para repartir el dinero mas eficazmente cuando existan varias tablas para la misma carrera
// Necesito tener un control a cual tabla pertenece el remate para repartir el dineros

// Remates struct
type Remates struct {
	gorm.Model
	Idracing    string `json:"idracing"`
	Idhorse     int    `json:"idhorse"`
	Numberhorse int    `json:"numberhorse"`
	Seudonimo   string `json:"seudonimo"`
	Amount      int64  `json:"amount"`
	Horsename   string `json:"horsename"`
}

// CreateRemates Remates db
func CreateRemates(idracing string, idhorse int, numberhorse int, seudonimo string, amount int64, horsename string) map[string]interface{} {
	remateGanador := &Remates{Idracing: idracing, Idhorse: idhorse, Numberhorse: numberhorse, Seudonimo: seudonimo, Amount: amount, Horsename: horsename}

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
		response["time"] = time.Now()
		return response
	}
	response := u.Message(true, "EMPTY")
	response["time"] = time.Now()
	return response
}
