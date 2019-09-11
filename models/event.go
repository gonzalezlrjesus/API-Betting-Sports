package models

import (
	u "API-Betting-Sports/utils"
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
)

// Event struct
type Event struct {
	gorm.Model
	Dateevent           time.Time 	`json:"dateevent"`
	Hipodromo      		string 		`json:"hipodromo"`
	Racingnumbers 		uint 		`json:"racingnumbers"`
	Minimumamount   	uint   		`json:"minimumamount"`
	Profitpercentage 	uint 		`json:"profitpercentage"`
	Auctionnumber 		uint 		`json:"auctionnumber"`
	TimebetweenAuctions uint 		`json:"timebetweenAuctions"`
	Horsenotauction 	string 		`json:"horsenotauction"`
}

// CreateEvent Event db
func (event *Event) CreateEvent() map[string]interface{} {

	if resp, ok := event.ValidateEvent(); !ok {
		return resp
	}

    event.Auctionnumber = 3
    event.TimebetweenAuctions = 5
    event.Horsenotauction = "casa"
	GetDB().Create(event)

	response := u.Message(true, "Event add to system")
	response["event"] = event
	return response
}

// UpdateEvent in DB I MUST CORRECT IT, BECAUSE I UPDATE NEW DATA***
func (event *Event) UpdateEvent(idEvent *string) map[string]interface{} {

	if resp, ok := event.ValidateEventParams(idEvent); !ok {
		return resp
	}

	temp := &Event{}

	err := GetDB().Table("events").Where("id = ?", *idEvent).First(temp).Error
	if err == gorm.ErrRecordNotFound {
		fmt.Println(err)
		return nil
	}

	temp.Dateevent = event.Dateevent
	temp.Hipodromo = event.Hipodromo
	temp.Minimumamount = event.Minimumamount
	temp.Profitpercentage = event.Profitpercentage

	GetDB().Save(&temp)

	response := u.Message(true, "Event has been updated")
	response["event"] = temp
	return response

}

// GetOneEvent event
func GetOneEvent(idEvent *string) map[string]interface{} {
	temp := &Event{}

	//check event specific in DB
	err := GetDB().Table("events").Where("id = ?", *idEvent).First(temp).Error
	if err == gorm.ErrRecordNotFound {
		fmt.Println("Event : ", err)
		return u.Message(true, "Event no exist")
	}

	response := u.Message(true, "Get Event")
	response["event"] = temp
	return response
}

// GetEvents all events of table events
func GetEvents() []*Event {

	events := make([]*Event, 0)

	err := GetDB().Table("events").Order("dateevent").Find(&events).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return events
}

// DeleteEvent from DB
func DeleteEvent(idEvent *string) bool {

	temp := &Event{}
	// Select Event
	err := GetDB().Table("events").Where("id= ?", *idEvent).First(temp).Error

	if err != nil || err == gorm.ErrRecordNotFound {
		return false
	}
	DeleteRacings(temp.ID)
	// Delete it
	GetDB().Delete(temp)
	// DeleteRacings(idEvent)
	return true
}

// ---------------------------Validations------------------------------

// ValidateEvent struct that Front-End to Back-End
func (event *Event) ValidateEvent() (map[string]interface{}, bool) {

	if event.Dateevent.IsZero() {
		return u.Message(false, "Date Event is empty"), false
	}

	// Data form
	tempForm := &Event{}
	errAux := GetDB().Table("events").Where("dateevent = ?", event.Dateevent).First(tempForm).Error
	if errAux != nil && errAux != gorm.ErrRecordNotFound {
		return u.Message(false, "Connection error. Please retry"), false
	}
	if !tempForm.Dateevent.IsZero() {
		return u.Message(false, "there is event with this Dateevent"), false
	}

	return u.Message(false, "Requirement passed"), true
}

// ValidateEventParams struct Params for Update Event
func (event *Event) ValidateEventParams(idEvent *string) (map[string]interface{}, bool) {

	if event.Dateevent.IsZero() {
		return u.Message(false, "Date Event is empty"), false
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

	// Data form
	tempForm := &Event{}
	errAux := GetDB().Table("events").Where("dateevent = ?", event.Dateevent).First(tempForm).Error
	if errAux != nil && errAux != gorm.ErrRecordNotFound {
		return u.Message(false, "Connection error. Please retry"), false
	}

	if errAux != gorm.ErrRecordNotFound && tempForm.Dateevent != tempIDevent.Dateevent {
		return u.Message(false, "there is event with this DateEvent"), false
	}

	return u.Message(false, "Requirement passed"), true
}
