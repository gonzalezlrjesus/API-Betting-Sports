package models

import (
	u "API-Betting-Sports/utils"
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
)

// Racing struct
type Racing struct {
	gorm.Model
	Eventid      uint      `json:"eventid"`
	Starttime    time.Time `json:"starttime"`
	Horsenumbers uint      `json:"horsenumbers"`
	Auctiontime  time.Time `json:"auctiontime"`
	Alerttime    time.Time `json:"alerttime"`
	Stateracing  string    `json:"stateracing"`
}

// ListRacings struct test
type ListRacings struct {
	Arrayracing []Racing
}

// CreateRacingModel db
func CreateRacingModel(Arrayracing []Racing, idEvent uint) map[string]interface{} {
	for i := range Arrayracing {
		Arrayracing[i].Eventid = idEvent
		if resp, ok := Arrayracing[i].ValidateRacing(); !ok {
			return resp
		}

		err := GetDB().Create(&Arrayracing[i]).Error
		fmt.Println("err:, ", err)
	}

	response := u.Message(true, "Racings has been created")
	response["Arrayracing"] = Arrayracing
	return response
}

// UpdateRacingModel in DB
func UpdateRacingModel(Arrayracing []Racing, idEvent uint) map[string]interface{} {

	if resp, ok := ValidateEventRacingParams(&idEvent); !ok {
		return resp
	}

	for i := range Arrayracing {

		temp := &Racing{}
		err := GetDB().Table("racings").Where("id = ?", Arrayracing[i].ID).First(temp).Error
		if err == gorm.ErrRecordNotFound {
			fmt.Println(err)
			return nil
		}
		temp = &Arrayracing[i]

		err = GetDB().Save(&temp).Error
		fmt.Println("err:, ", err)
	}

	response := u.Message(true, "Racings has been updated")
	return response

}

// CloseRacing in DB
func CloseRacing(idRacing string) map[string]interface{} {

		temp := &Racing{}
		err := GetDB().Table("racings").Where("id = ?", idRacing).First(temp).Error
		if err == gorm.ErrRecordNotFound {
			fmt.Println(err)
			return nil
		}
		temp.Stateracing = "CLOSED"

		err = GetDB().Save(&temp).Error
		fmt.Println("err:, ", err)

	response := u.Message(true, "Racings has become to closed")
	return response

}


// GetOneRacing Racing
// func GetOneRacing(idEvent, idRacing *string) map[string]interface{} {
func GetOneRacing(idRacing *string) map[string]interface{} {
	temp := &Racing{}

	//check racing specific in DB
	// err := GetDB().Table("racings").Where("id= ? AND eventid = ?", *idRacing, *idEvent).First(temp).Error
	err := GetDB().Table("racings").Where("id= ?", *idRacing).First(temp).Error
	if err == gorm.ErrRecordNotFound {
		fmt.Println("Racing : ", err)
		return u.Message(true, "Racing no exist")
	}

	response := u.Message(true, "Get Racing")
	response["racing"] = temp
	return response
}

func FindRacingWithinEvent(idEvent *string, idRacing *string) map[string]interface{} {
    temp := &Racing{}

	//check racing specific in DB
	err := GetDB().Table("racings").Where("id= ? AND eventid = ?", *idRacing, *idEvent).First(temp).Error
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
	errRacing := GetDB().Table("racings").Where("id= ? AND eventid = ?", *idRacing, idEvent).First(tempRacing).Error

	if errRacing != nil || errRacing == gorm.ErrRecordNotFound {
		return false
	}

	// Delete it
	GetDB().Delete(tempRacing)
	DeleteRacingComponents(tempRacing.ID)
	DeleteAllHorses(tempRacing.ID)
	return true
}

// DeleteRacings delete all racings
func DeleteRacings(idComponent uint) bool {

	racings := &[]Racing{}

	errRacings := GetDB().Table("racings").Where("eventid = ?", idComponent).Find(&racings).Error
	if errRacings != nil {
		fmt.Println(errRacings)
		return false
	}

	tempDelete := &[]Racing{}

	// idEvent := *idComponent
	// tempUint64, _ := strconv.ParseUint(idEvent, 10, 32)

	// delete all racings
	err := GetDB().Table("racings").Where("eventid = ?", idComponent).Delete(tempDelete).Error
	if err == gorm.ErrRecordNotFound {
		fmt.Println("Racingcomponentid  : ", err)
		return false
	}
		fmt.Println("Error  : ", err)
	for _, JustRace := range *racings {
		// DeleteRacingComponents(JustRace.ID)
		DeleteAllHorses(JustRace.ID)
	}

	return true
}

// ---------------------------Validations------------------------------

// ValidateRacing struct that Front-End to Back-End
func (racing *Racing) ValidateRacing() (map[string]interface{}, bool) {

	// if racing.Hipodromo == "" {
	// 	return u.Message(false, "Hipodromo is required"), false
	// }

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
func ValidateEventRacingParams(idEvent *uint) (map[string]interface{}, bool) {

	tempIDevent := &Event{}

	// Search idEvent in DB
	erridEvent := GetDB().Table("events").Where("id = ?", *idEvent).First(tempIDevent).Error
	if erridEvent == gorm.ErrRecordNotFound {
		fmt.Println(erridEvent)
		return u.Message(false, "Not found ID Event Param"), false
	}

	return u.Message(false, "Requirement passed"), true
}
