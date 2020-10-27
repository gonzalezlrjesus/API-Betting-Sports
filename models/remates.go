package models

import (
	"time"

	u "github.com/gonzalezlrjesus/API-Betting-Sports/utils"
	"github.com/jinzhu/gorm"
)

// Cuando el proyecto este mas avanzado deberia llevar un control de la tabla que se esta jugando
// Todo eso para repartir el dinero mas eficazmente cuando existan varias tablas para la misma carrera
// Necesito tener un control a cual tabla pertenece el remate para repartir el dineros

// Remates struct
type Remates struct {
	gorm.Model
	Idracing    uint   `json:"idracing"`
	Idhorse     int    `json:"idhorse"`
	Numberhorse int    `json:"numberhorse"`
	Seudonimo   string `json:"seudonimo"`
	Amount      int64  `json:"amount"`
	Horsename   string `json:"horsename"`
}

// CreateRemates .
func CreateRemates(idracing uint, idhorse int, numberhorse int, seudonimo string, amount int64, horsename string) map[string]interface{} {
	remateGanador := &Remates{
		Idracing:    idracing,
		Idhorse:     idhorse,
		Numberhorse: numberhorse,
		Seudonimo:   seudonimo,
		Amount:      amount,
		Horsename:   horsename}

	GetDB().Create(remateGanador)

	response := u.Message(true, "Remate added")
	response["remate"] = remateGanador
	return response
}

// GetRemates .
func GetRemates(idracing uint) map[string]interface{} {

	remates, err := searchRematesByRacingID(idracing)
	if err != nil {
		return nil
	}

	response := u.Message(true, "Remates")
	response["remates"] = remates
	response["time"] = time.Now()
	return response
}

// SearchRemateByRaceIDAndHorseID .
func SearchRemateByRaceIDAndHorseID(idRacing uint, idHorse int) (*Remates, error) {
	temp := &Remates{}
	err := GetDB().Table("remates").Where("idracing = ? AND idhorse = ?", idRacing, idHorse).Find(temp).Error
	return temp, err
}

func searchRematesByRacingID(idracing uint) ([]*Remates, error) {
	remates := make([]*Remates, 0)
	err := GetDB().Table("remates").Where("idracing = ?", idracing).Find(&remates).Error
	return remates, err
}
