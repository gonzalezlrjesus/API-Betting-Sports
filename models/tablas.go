package models

import (
	u "github.com/gonzalezlrjesus/API-Betting-Sports/utils"

	"github.com/jinzhu/gorm"
)

// Tablas struct
type Tablas struct {
	gorm.Model
	Idracing         uint   `json:"idracing"`
	Montototal       int64  `json:"montototal"`
	Montocasa        int64  `json:"montocasa"`
	Montoganador     int64  `json:"montoganador"`
	Posiciontabla    uint   `json:"posiciontabla"`
	Porcentajeevento uint   `json:"posicionevento"`
	Estado           string `json:"estado"`
}

// CreateTablas .
func CreateTablas(idracing uint, montoTotal int64) map[string]interface{} {

	race, _ := ExistRaceID(idracing)
	event, _ := ExistEventID(race.Eventid)

	var gananciaCasa int64 = (int64(event.Profitpercentage) * montoTotal) / 100
	var gananciaGanador int64 = montoTotal - gananciaCasa
	tablaAlmacenar := &Tablas{
		Idracing:         idracing,
		Montototal:       montoTotal,
		Montocasa:        gananciaCasa,
		Montoganador:     gananciaGanador,
		Posiciontabla:    1,
		Porcentajeevento: event.Profitpercentage,
		Estado:           "ESPERANDO",
	}

	GetDB().Create(tablaAlmacenar)

	response := u.Message(true, "Tabla added")
	response["tabla"] = tablaAlmacenar
	return response
}

// GetTablas .
func GetTablas(idracing uint) map[string]interface{} {

	tablas, err := searchAllTablasByRaceID(idracing)
	if err != nil {
		return nil
	}

	response := u.Message(true, "Tablas")
	response["tablas"] = tablas
	return response
}

// UpdateStateTabla tablas
func (tabla *Tablas) UpdateStateTabla() map[string]interface{} {
	tabla.Estado = "PAGADO"
	GetDB().Save(&tabla)
	response := u.Message(true, "Tabla has been updated")
	response["tabla"] = tabla
	return response
}

// SearchTablaByRaceID .
func SearchTablaByRaceID(idRace uint) (*Tablas, error) {
	temp := &Tablas{}
	err := GetDB().Table("tablas").Where("idracing = ?", idRace).Find(temp).Error
	return temp, err
}

func searchAllTablasByRaceID(idRace uint) ([]*Tablas, error) {
	tablas := make([]*Tablas, 0)
	err := GetDB().Table("tablas").Where("idracing = ?", idRace).Find(tablas).Error
	return tablas, err
}
