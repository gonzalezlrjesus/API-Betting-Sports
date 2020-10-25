package models

import (
	"os"
	"strconv"
	"time"

	u "github.com/gonzalezlrjesus/API-Betting-Sports/utils"

	"github.com/jinzhu/gorm"
)

// Event struct
type Event struct {
	gorm.Model
	Dateevent           time.Time `json:"dateevent"`
	Hipodromo           string    `json:"hipodromo"`
	Racingnumbers       uint      `json:"racingnumbers"`
	Minimumamount       uint      `json:"minimumamount"`
	Profitpercentage    uint      `json:"profitpercentage"`
	Auctionnumber       uint      `json:"auctionnumber"`
	TimebetweenAuctions uint      `json:"timebetweenAuctions"`
	Horsenotauction     string    `json:"horsenotauction"`
}

// CreateEvent Event db
func (event *Event) CreateEvent() map[string]interface{} {

	if resp, ok := event.ValidateEvent(); !ok {
		return resp
	}

	tempAucnumber, _ := strconv.ParseUint(os.Getenv("Auctionnumber"), 10, 64)
	tempimebetweenAuctions, _ := strconv.ParseUint(os.Getenv("TimebetweenAuctions"), 10, 64)

	event.Auctionnumber = uint(tempAucnumber)
	event.TimebetweenAuctions = uint(tempimebetweenAuctions)
	event.Horsenotauction = os.Getenv("Horsenotauction")
	GetDB().Create(event)

	response := u.Message(true, "Event add to system")
	response["event"] = event
	return response
}

// UpdateEvent .
func (event *Event) UpdateEvent(idEvent uint) map[string]interface{} {

	if resp, ok := event.ValidateEventParams(idEvent); !ok {
		return resp
	}

	temp, err := existEventID(idEvent)
	if err == gorm.ErrRecordNotFound {
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
func GetOneEvent(idEvent uint) map[string]interface{} {
	temp, err := existEventID(idEvent)
	if err == gorm.ErrRecordNotFound {
		return u.Message(true, "Event no exist")
	}

	response := u.Message(true, "Get Event")
	response["event"] = temp
	return response
}

// GetEvents all events of table events
func GetEvents() map[string]interface{} {

	events := make([]*Event, 0)

	err := GetDB().Table("events").Order("dateevent").Find(&events).Error
	if err != nil {
		return nil
	}

	response := u.Message(true, "Get Event")
	response["events"] = events
	response["time"] = time.Now()
	return response
}

// DeleteEvent from DB
func DeleteEvent(idEvent uint) bool {
	temp, err := existEventID(idEvent)
	if err != nil || err == gorm.ErrRecordNotFound {
		return false
	}

	// DeleteRacings(idEvent)
	DeleteRacings(temp.ID)
	// Delete it
	GetDB().Delete(temp)

	return true
}

// ---------------------------Validations------------------------------

// ValidateEvent struct that Front-End to Back-End
func (event *Event) ValidateEvent() (map[string]interface{}, bool) {

	if event.Dateevent.IsZero() {
		return u.Message(false, "event date is empty"), false
	}

	tempForm, errAux := existEventDate(event.Dateevent)
	if errAux != nil && errAux != gorm.ErrRecordNotFound {
		return u.Message(false, "Connection error. Please retry"), false
	}

	if !tempForm.Dateevent.IsZero() {
		return u.Message(false, "there is event with this event date"), false
	}

	return u.Message(false, "Requirement passed"), true
}

// ValidateEventParams struct Params for Update Event
func (event *Event) ValidateEventParams(idEvent uint) (map[string]interface{}, bool) {

	if event.Dateevent.IsZero() {
		return u.Message(false, "Date Event is empty"), false
	}

	temp, err := existEventID(idEvent)
	if err == gorm.ErrRecordNotFound {
		return u.Message(false, "Not found ID Event Param"), false
	}

	tempForm, errAux := existEventDate(event.Dateevent)
	if errAux != nil && errAux != gorm.ErrRecordNotFound {
		return u.Message(false, "Connection error. Please retry"), false
	}

	if errAux != gorm.ErrRecordNotFound && tempForm.Dateevent != temp.Dateevent {
		return u.Message(false, "there is event with this DateEvent"), false
	}

	return u.Message(false, "Requirement passed"), true
}

func existEventID(idEvent uint) (*Event, error) {
	tempEvent := &Event{}
	err := GetDB().Table("events").Where("id = ?", idEvent).First(tempEvent).Error
	return tempEvent, err
}

func existEventDate(eventDate time.Time) (*Event, error) {
	tempEvent := &Event{}
	err := GetDB().Table("events").Where("dateevent = ?", eventDate).First(tempEvent).Error
	return tempEvent, err
}
