package models

import (
	u "github.com/gonzalezlrjesus/API-Betting-Sports/utils"
	"fmt"
	"time"

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

// CreateHorseModel add a new race horse array to db in the table Horse
func CreateHorseModel(Arrayhorse []Horse, idRacing uint) map[string]interface{} {

	if resp, ok := ValidateHorse(idRacing); !ok {
		return resp
	}

	for i := range Arrayhorse {
		Arrayhorse[i].Racingid = idRacing
		// if resp, ok := Arrayhorse[i].ValidateRacing(); !ok {
		// 	return resp
		// }

		err := GetDB().Create(&Arrayhorse[i]).Error
		fmt.Println("err:, ", err)
	}

	response := u.Message(true, "Horse has been created")
	response["horse"] = Arrayhorse
	return response
}

// UpdateHorse in DB
func (horse *Horse) UpdateHorse(idRacing, idHorse *string) map[string]interface{} {

	if resp, ok := horse.ValidateHorseParams(idRacing); !ok {
		return resp
	}

	temp := &Horse{}

	err := GetDB().Table("horses").Where("id = ?", *idHorse).First(temp).Error
	if err == gorm.ErrRecordNotFound {
		fmt.Println(err)
		return u.Message(true, "Horse not exist")
	}

	temp.Horsename = horse.Horsename

	GetDB().Save(&temp)

	response := u.Message(true, "Horse has been updated")
	response["horse"] = temp
	return response

}

// GetHorses all Horses of Horses table
func GetHorses(idRacing *string) map[string]interface{} {

	horses := &[]Horse{}

	err := GetDB().Table("horses").Where("Racingid = ?", *idRacing).Find(&horses).Error
	if err != nil {
		// fmt.Println(err)
		return nil
	}

	response := u.Message(true, "Get all Horse")
	response["allHorses"] = horses
	return response
}

// DeleteHorse Delete a Horse from DB
func DeleteHorse(idRacing, idHorse *string) bool {

	tempRacing := &Racing{}
	// Search Racing
	errRacing := GetDB().Table("racings").Where("id= ?", *idRacing).First(tempRacing).Error

	if errRacing != nil || errRacing == gorm.ErrRecordNotFound {
		return false
	}

	tempHorse := &Horse{}
	// Search Horse
	errHorse := GetDB().Table("horses").Where("id= ? AND Racingid = ?", *idHorse, *idRacing).First(tempHorse).Error

	if errHorse != nil || errHorse == gorm.ErrRecordNotFound {
		return false
	}

	// Delete it
	GetDB().Delete(tempHorse)
	return true
}

// DeleteAllHorses delete all horses of race.
func DeleteAllHorses(idRacing uint) bool {

	temp := &[]Horse{}

	//check and delete all horses
	err := GetDB().Table("horses").Where("Racingid = ?", idRacing).Delete(temp).Error
	if err == gorm.ErrRecordNotFound {
		fmt.Println("Delete all horses : ", err)
		return false
	}

	return true
}

// WithdrawHorse withdraw Horse .
func WithdrawHorse(idHorse uint, idRacing uint) map[string]interface{} {

	temp := &Horse{}
	tempRacing := &Racing{}
	errRacing := GetDB().Table("racings").Where("id = ?", idRacing).First(tempRacing).Error
	if errRacing == gorm.ErrRecordNotFound {
		fmt.Println("Racings Not found : ", errRacing)
		return u.Message(true, "Racings has not found")
	}
	// Debo validar si el server time es mayor o menor que tempRacing.Starttime
	// Si es menor retiro caballo , color rojo en frontend cliente y no se remata
	// si es mayor retiro caballo, pregunto si apostaron por el el si es positivo
	// devuelvo dinero a apostador y modifico monto total a pagar
	// si no se aposto retiro caballo y modifico monto
	//

	//check and retirar caballo especifico
	err := GetDB().Table("horses").Where("id = ? AND racingid = ?", idHorse, idRacing).First(temp).Error
	fmt.Println("err : ", err)
	if err == gorm.ErrRecordNotFound {
		fmt.Println("Not found : ", err)
		return u.Message(true, "Horse has not found")
	}

	temp.State = "withdrawed"

	GetDB().Save(&temp)

	SegAntesRemate := tempRacing.Auctiontime.Sub(time.Now()).Seconds()
	SegDespuesCarrera := tempRacing.Starttime.Sub(time.Now()).Seconds()

	if SegAntesRemate > 0 {
		fmt.Println("No ha inicidado  Remate: ", SegAntesRemate)
	}

	if (SegAntesRemate < 0) == true && (SegDespuesCarrera > 0) == true {
		fmt.Println("Estan Rematando: ", SegDespuesCarrera)
	}

	if (SegAntesRemate < 0) == true && (SegDespuesCarrera < 0) == true {
		fmt.Println("Termino Remate y Carrera: ", SegDespuesCarrera)
	}

	response := u.Message(true, "Horse has been withdrawed")
	response["horse"] = temp
	return response

}

// ---------------------------Validations------------------------------

// ValidateHorse struct that Front-End to Back-End
func ValidateHorse(racingID uint) (map[string]interface{}, bool) {

	// if horse.Horsename == "" {
	// 	return u.Message(false, "Horse name is required"), false
	// }

	temp := &Racing{}
	//check Racing in DB
	err := GetDB().Table("racings").Where("id = ?", racingID).First(temp).Error
	if err == gorm.ErrRecordNotFound {
		fmt.Println("Racing", err)
		return u.Message(false, "Racing exist no in DB"), false
	}

	return u.Message(false, "Requirement passed"), true
}

// ValidateHorseParams struct Params for Update Racing
func (horse *Horse) ValidateHorseParams(idRacing *string) (map[string]interface{}, bool) {

	if horse.Horsename == "" {
		return u.Message(false, "Horse name is required"), false
	}

	tempIDRacing := &Racing{}

	// Search Racing in DB
	erridRacing := GetDB().Table("racings").Where("id = ?", *idRacing).First(tempIDRacing).Error
	if erridRacing == gorm.ErrRecordNotFound {
		fmt.Println(erridRacing)
		return u.Message(false, "Not found ID Racing Param"), false
	}

	return u.Message(false, "Requirement passed"), true
}
