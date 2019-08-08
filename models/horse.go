package models

import (
	"fmt"

	u "test-golang/utils"

	"github.com/jinzhu/gorm"
)

// Horse struct
type Horse struct {
	gorm.Model
	Racingid      uint   `json:"racingid"`
	Horsename     string `json:"horsename"`
	Jockeyname    string `json:"jockeyname"`
	Finalposition string `json:"finalposition"`
}

// CreateHorse add a new race horse to db in the table Horse
func (horse *Horse) CreateHorse() map[string]interface{} {

	if resp, ok := horse.ValidateHorse(); !ok {
		return resp
	}

	GetDB().Create(horse)

	response := u.Message(true, "Horse has been created")
	response["horse"] = horse
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
	temp.Jockeyname = horse.Jockeyname
	temp.Finalposition = horse.Finalposition

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
		fmt.Println(err)
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

// ---------------------------Validations------------------------------

// ValidateHorse struct that Front-End to Back-End
func (horse *Horse) ValidateHorse() (map[string]interface{}, bool) {

	if horse.Horsename == "" {
		return u.Message(false, "Horse name is required"), false
	}

	temp := &Racing{}
	//check Racing in DB
	err := GetDB().Table("racings").Where("id = ?", horse.Racingid).First(temp).Error
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
