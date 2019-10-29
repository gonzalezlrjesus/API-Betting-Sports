package models

import (
	u "API-Betting-Sports/utils"
	"encoding/json"
	"fmt"

	"github.com/jinzhu/gorm"
)

// Tablas struct
type Tablas struct {
	gorm.Model
	Idracing         string `json:"idracing"`
	Montototal       int64  `json:"montototal"`
	Montocasa        int64  `json:"montocasa"`
	Montoganador     int64  `json:"montoganador"`
	Posiciontabla    uint   `json:"posiciontabla"`
	Porcentajeevento uint   `json:"posicionevento"`
	Estado           string `json:"estado"`
}

// CreateTablas Tablas db
func CreateTablas(idracing string, montoTotal int64) map[string]interface{} {

	testGetRacing := GetOneRacing(&idracing)
	jsonMessage2, _ := json.Marshal(testGetRacing["racing"])
	fmt.Println(string(jsonMessage2))
	s := string(jsonMessage2)
	data := Racing{}
	json.Unmarshal([]byte(s), &data)

	// GetOneEventMonto event Params: idEvent uint
	testGetEvent := GetOneEventMonto(data.Eventid)

	jsonMessage3, _ := json.Marshal(testGetEvent["event"])
	sw := string(jsonMessage3)
	dataEvent := Event{}
	json.Unmarshal([]byte(sw), &dataEvent)

	var gananciaCasa int64 = (int64(dataEvent.Profitpercentage) * montoTotal) / 100
	var gananciaGanador int64 = montoTotal - gananciaCasa
	tablaAlmacenar := &Tablas{
		Idracing:         idracing,
		Montototal:       montoTotal,
		Montocasa:        gananciaCasa,
		Montoganador:     gananciaGanador,
		Posiciontabla:    1,
		Porcentajeevento: dataEvent.Profitpercentage,
		Estado:           "ESPERANDO",
	}

	GetDB().Create(tablaAlmacenar)

	response := u.Message(true, "Tabla added")
	response["tabla"] = tablaAlmacenar
	return response
}

// GetTablas Tablas db
func GetTablas(idracing *string) map[string]interface{} {

	tablas := make([]*Tablas, 0)

	err := GetDB().Table("tablas").Where("idracing = ?", *idracing).Find(&tablas).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}

	if len(tablas) > 0 {
		response := u.Message(true, "Tablas")
		response["tablas"] = tablas
		return response
	}
	response := u.Message(true, "EMPTY")
	return response
}
