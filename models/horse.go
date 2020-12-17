package models

import (
	"fmt"
	"log"
	"time"

	u "github.com/gonzalezlrjesus/API-Betting-Sports/utils"

	"github.com/jinzhu/gorm"
)

// Horse struct
type Horse struct {
	gorm.Model
	Racingid  uint   `json:"racingid"`
	Horsename string `json:"horsename"`
	Numero    uint   `json:"numero"`
	State     string `json:"state"`
}

// CreateHorses add new race horses
func CreateHorses(Arrayhorse []Horse) map[string]interface{} {

	for i := range Arrayhorse {
		GetDB().Create(&Arrayhorse[i])
	}

	response := u.Message(true, "Horses has been created")
	response["horses"] = Arrayhorse
	return response
}

// UpdateHorse in DB
func (horse *Horse) UpdateHorse(idRacing uint, idHorse string) map[string]interface{} {

	if resp, ok := horse.ValidateHorseParams(idRacing); !ok {
		return resp
	}

	temp := &Horse{}
	err := GetDB().Table("horses").Where("id = ?", idHorse).First(temp).Error
	if err == gorm.ErrRecordNotFound {
		return u.Message(true, "Horse not exist")
	}

	temp.Horsename = horse.Horsename
	GetDB().Save(&temp)

	response := u.Message(true, "Horse has been updated")
	response["horse"] = temp
	return response
}

// GetHorses all Horses
func GetHorses(idRacing string) map[string]interface{} {

	horses := &[]Horse{}
	err := GetDB().Table("horses").Where("Racingid = ?", idRacing).Order("id").Find(horses).Error
	if err != nil {
		return nil
	}

	response := u.Message(true, "Get all Horse")
	response["allHorses"] = horses
	return response
}

// DeleteHorse from DB
func DeleteHorse(idRacing uint, idHorse uint) bool {

	_, err := ExistRaceID(idRacing)
	if err == gorm.ErrRecordNotFound {
		return false
	}

	tempHorse, err := existSpecificHorse(idHorse, idRacing)
	if err == gorm.ErrRecordNotFound {
		return false
	}

	// Delete it
	GetDB().Delete(tempHorse)
	return true
}

// DeleteAllHorses delete all horses of race.
func DeleteAllHorses(idRacing uint) bool {

	temp := &[]Horse{}
	err := GetDB().Table("horses").Where("Racingid = ?", idRacing).Delete(temp).Error
	if err == gorm.ErrRecordNotFound {
		return false
	}

	return true
}

// WithdrawHorse .
func WithdrawHorse(idHorse uint, idRacing uint) map[string]interface{} {

	tempRacing, err := ExistRaceID(idRacing)
	if err == gorm.ErrRecordNotFound {
		return u.Message(true, "Race not exist")
	}

	// Debo validar si el server time es mayor o menor que tempRacing.Starttime
	// Si es menor retiro caballo , color rojo en frontend cliente y no se remata
	// si es mayor retiro caballo, pregunto si apostaron por el el si es positivo
	// devuelvo dinero a apostador y modifico monto total a pagar
	// si no se aposto retiro caballo y modifico monto
	// Validar que se cerro el remate
	//

	tempHorse, err := existSpecificHorse(idHorse, idRacing)
	if err == gorm.ErrRecordNotFound {
		return u.Message(true, "Horse has not found")
	}

	SegAntesRemate := tempRacing.Auctiontime.Sub(time.Now()).Seconds()
	SegDespuesCarrera := tempRacing.Starttime.Sub(time.Now()).Seconds()

	if (SegAntesRemate < 0) && (SegDespuesCarrera > 0) {
		fmt.Println("Estan Rematando: ", SegDespuesCarrera)
		response := u.Message(true, "Horse has not been withdrawn")
		response["horse"] = tempHorse
		return response
	}

	tempHorse.State = "withdrawed"

	GetDB().Save(&tempHorse)

	if SegAntesRemate > 0 {
		log.Println("No ha inicidado  Remate: ", SegAntesRemate)
	}

	if (SegAntesRemate < 0) && (SegDespuesCarrera < 0) {
		log.Println("Termino Remate y Carrera: ", SegDespuesCarrera)

		// REMATES
		tempREMATES, err := SearchRemateByRaceIDAndHorseID(idRacing, int(idHorse))
		if err == gorm.ErrRecordNotFound {
			return u.Message(true, "Remate not exist")
		}

		//check client-seudonimo in DB
		tempClient, err := ExistClientSeudonimonDB(tempREMATES.Seudonimo)
		if err == gorm.ErrRecordNotFound {
			return u.Message(true, "Client not exist")
		}

		tempDeposit := &Deposit{Amount: float64(tempREMATES.Amount), Clientidentificationcard: tempClient.Identificationcard, FormaPago: "reintegro"}
		tempDeposit.AddDepositClient()
		UpdateMontos(idRacing, tempREMATES.Amount)
	}

	response := u.Message(true, "Horse has been withdrawed")
	response["horse"] = tempHorse
	return response
}

// ---------------------------Validations------------------------------

// ValidateHorseParams .
func (horse *Horse) ValidateHorseParams(idRacing uint) (map[string]interface{}, bool) {

	if horse.Horsename == "" {
		return u.Message(false, "Horse name is required"), false
	}

	_, err := ExistRaceID(idRacing)
	if err == gorm.ErrRecordNotFound {
		return u.Message(false, "Not found ID Racing Param"), false
	}

	return u.Message(false, "Requirement passed"), true
}

func existSpecificHorse(idHorse, idRacing uint) (*Horse, error) {
	temp := &Horse{}
	err := GetDB().Table("horses").Where("id = ? AND racingid = ?", idHorse, idRacing).First(temp).Error
	return temp, err
}
