package models

import (
	"fmt"

	u "test-golang/utils"

	"github.com/jinzhu/gorm"
)

// Racing struct
type Racing struct {
	gorm.Model
	Eventid   string `json:"eventid"`
	Hipodromo string `json:"hipodromo"`
}

// CreateRacing db
func (racing *Racing) CreateRacing() map[string]interface{} {

	if resp, ok := racing.ValidateRacing(); !ok {
		return resp
	}

	GetDB().Create(racing)
	// this are comment because here i must write create Component racing
	// temp := &Racing{Eventid: deposit.Clientidentificationcard}

	// //check client_Coins in DB
	// err := GetDB().Table("coins").Where("ClientIdentificationcard = ?", temp.Clientidentificationcard).First(temp).Error
	// if err == gorm.ErrRecordNotFound {
	// 	fmt.Println("Client Coins : ", err)
	// 	return nil
	// }

	response := u.Message(true, "Racing has been created")
	// response["updateCoins"] = temp.UpdateCoins(deposit.Amount)
	response["racing"] = racing
	return response
}

// UpdateRacing in DB
func (racing *Racing) UpdateRacing(idEvent, idRacing *string) map[string]interface{} {

	if resp, ok := racing.ValidateEventRacingParams(idEvent, idRacing); !ok {
		return resp
	}

	temp := &Racing{}

	err := GetDB().Table("racings").Where("id = ?", *idRacing).First(temp).Error
	if err == gorm.ErrRecordNotFound {
		fmt.Println(err)
		return nil
	}

	temp.Hipodromo = racing.Hipodromo

	GetDB().Save(&temp)

	response := u.Message(true, "Racing has been updated")
	response["racing"] = temp
	return response

}

// GetOneRacing Racing
func GetOneRacing(idEvent, idRacing *string) map[string]interface{} {
	temp := &Racing{}

	//check racing specific in DB
	err := GetDB().Table("racings").Where("id= ? AND eventid >= ?", *idRacing, *idEvent).First(temp).Error
	if err == gorm.ErrRecordNotFound {
		fmt.Println("Racing : ", err)
		return u.Message(true, "Racing no exist")
	}

	response := u.Message(true, "Get Racing")
	response["racing"] = temp
	return response
}

// GetRacings all Racings of table Racings
func GetRacings(idEvent *string) []*Racing {

	racings := make([]*Racing, 0)

	err := GetDB().Table("racings").Where("eventid = ?", *idEvent).Find(&racings).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return racings
}

// DeleteRacing from DB
func DeleteRacing(idEvent, idRacing *string) bool {

	tempEvent := &Event{}
	// Search Event
	errEvent := GetDB().Table("events").Where("id= ?", *idEvent).First(tempEvent).Error

	if errEvent != nil || errEvent == gorm.ErrRecordNotFound {
		return false
	}

	tempRacing := &Racing{}
	// Search Racing
	errRacing := GetDB().Table("racings").Where("id= ? AND eventid >= ?", *idRacing, idEvent).First(tempRacing).Error

	if errRacing != nil || errRacing == gorm.ErrRecordNotFound {
		return false
	}

	// Delete it
	GetDB().Delete(tempRacing)
	return true
}

// ---------------------------Validations------------------------------

// ValidateRacing struct that Front-End to Back-End
func (racing *Racing) ValidateRacing() (map[string]interface{}, bool) {

	if racing.Hipodromo == "" {
		return u.Message(false, "Hipodromo is required"), false
	}

	temp := &Racing{}
	//check Event in DB
	err := GetDB().Table("events").Where("id = ?", racing.Eventid).First(temp).Error
	if err == gorm.ErrRecordNotFound {
		fmt.Println("Event", err)
		return u.Message(false, "Event exist no in DB"), false
	}

	return u.Message(false, "Requirement passed"), true
}

// ValidateEventRacingParams struct Params for Update Racing
func (racing *Racing) ValidateEventRacingParams(idEvent, idRacing *string) (map[string]interface{}, bool) {

	if racing.Hipodromo == "" {
		return u.Message(false, "Hipodromo is required"), false
	}

	tempIDevent := &Event{}

	// Search idEvent in DB
	erridEvent := GetDB().Table("events").Where("id = ?", *idEvent).First(tempIDevent).Error
	if erridEvent == gorm.ErrRecordNotFound {
		fmt.Println(erridEvent)
		return u.Message(false, "Not found ID Event Param"), false
	}

	if erridEvent == gorm.ErrRecordNotFound {
		return u.Message(false, "Not found ID Event Param"), false
	}

	tempIDRacing := &Racing{}

	// Search idEvent in DB
	erridRacing := GetDB().Table("racings").Where("id = ?", *idRacing).First(tempIDRacing).Error
	if erridRacing == gorm.ErrRecordNotFound {
		fmt.Println(erridRacing)
		return u.Message(false, "Not found ID Racing Param"), false
	}

	if erridRacing == gorm.ErrRecordNotFound {
		return u.Message(false, "Not found ID Racing Param"), false
	}

	return u.Message(false, "Requirement passed"), true
}
