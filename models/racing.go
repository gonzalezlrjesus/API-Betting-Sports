package models

import (
	"math"

	u "github.com/gonzalezlrjesus/API-Betting-Sports/utils"

	"time"

	"github.com/jinzhu/gorm"
)

// Racing struct
type Racing struct {
	gorm.Model
	Eventid              uint      `json:"eventid"`
	Starttime            time.Time `json:"starttime"`
	Horsenumbers         uint      `json:"horsenumbers"`
	Auctiontime          time.Time `json:"auctiontime"`
	Alerttime            time.Time `json:"alerttime"`
	Stateracing          string    `json:"stateracing"`
	Idcaballoganador     int       `json:"idcaballoganador"`
	Nombrecaballoganador string    `json:"nombrecaballoganador"`
}

// CreateRaces .
func CreateRaces(racesList []Racing, idEvent uint) map[string]interface{} {

	_, err := ExistEventID(idEvent)
	if err == gorm.ErrRecordNotFound {
		return u.Message(false, "Event exist no in DB")
	}

	for i := range racesList {
		racesList[i].Eventid = idEvent
		GetDB().Create(&racesList[i])
	}

	response := u.Message(true, "Racings has been created")
	response["races-list"] = racesList
	return response
}

// UpdateRaces .
func UpdateRaces(Arrayracing []Racing, idEvent uint) map[string]interface{} {

	_, err := ExistEventID(idEvent)
	if err == gorm.ErrRecordNotFound {
		return u.Message(false, "Event exist no in DB")
	}

	for i := range Arrayracing {
		temp, err := ExistRaceID(Arrayracing[i].ID)
		if err == gorm.ErrRecordNotFound {
			return nil
		}

		temp = &Arrayracing[i]
		GetDB().Save(&temp)
	}

	response := u.Message(true, "Racings has been updated")
	return response
}

// CloseRacing .
func CloseRacing(idRacing uint) map[string]interface{} {

	temp, err := ExistRaceID(idRacing)
	if err == gorm.ErrRecordNotFound {
		return nil
	}

	temp.Stateracing = "CLOSED"
	GetDB().Save(&temp)

	response := u.Message(true, "Racings has become to closed")
	return response
}

// TimeisEqualStartTime in DB
func TimeisEqualStartTime(idRacing uint) bool {

	temp, err := ExistRaceID(idRacing)
	if err == gorm.ErrRecordNotFound {
		return false
	}

	// Hours
	hs := temp.Starttime.Sub(time.Now()).Hours()
	// Minutes
	hs, mf := math.Modf(hs)
	ms := mf * 60
	// Seconds
	ms, sf := math.Modf(ms)
	ss := sf * 60

	// fmt.Println(hs, "hours", ms, "minutes", int(ss), "seconds")
	// fmt.Println(ms == 0, "minutes == 0")
	// fmt.Println(int(ss) == 0, "ss == 0")

	if ms == 0 && int(ss) == 0 {
		CloseRacing(idRacing)
		return true
	}
	return false
}

// GetRace .
func GetRace(idRacing uint) map[string]interface{} {

	temp, err := ExistRaceID(idRacing)
	if err == gorm.ErrRecordNotFound {
		return u.Message(true, "Race not exist")
	}

	response := u.Message(true, "Get Race")
	response["racing"] = temp
	return response
}

// FindRaceByEventID find
func FindRaceByEventID(idEvent, idRacing uint) map[string]interface{} {

	temp, err := searchRacingByIDandEventID(idRacing, idEvent)
	if err == gorm.ErrRecordNotFound {
		return u.Message(true, "Race not exist")
	}

	response := u.Message(true, "Get Race")
	response["racing"] = temp
	response["time"] = time.Now()
	return response
}

// RepartirGanancias find
func RepartirGanancias(idRace uint, idHorse int) map[string]interface{} {
	// RACE
	tempRacing, err := ExistRaceID(idRace)
	if err == gorm.ErrRecordNotFound {
		return u.Message(true, "Race not exist")
	}

	// REMATES
	tempREMATES, err := SearchRemateByRaceIDAndHorseID(idRace, idHorse)
	if err == gorm.ErrRecordNotFound {
		return u.Message(true, "Remate not exist")
	}

	// TABLAS
	tempTablas, err := SearchTablaByRaceID(idRace)
	if err == gorm.ErrRecordNotFound {
		return u.Message(true, "Tabla not exist")
	}

	AddGananciaClient(tempREMATES.Seudonimo, tempTablas.Montoganador, "Ganancia", tempTablas.ID)
	AddGananciaClient("CASA", tempTablas.Montocasa, "CASA", tempTablas.ID)

	tempTablas.UpdateStateTabla()

	tempRacing.Idcaballoganador = idHorse
	GetDB().Save(&tempRacing)
	response := u.Message(true, "Ganancias Repartidas")
	return response
}

// GetRacings all Racings by eventID
func GetRacings(idEvent uint) map[string]interface{} {

	racings, err := searchAllRacesByEventID(idEvent)
	if err != nil {
		return nil
	}

	response := u.Message(true, "Get all races")
	response["races"] = racings
	return response
}

// DeleteRacing from DB
func DeleteRacing(idEvent, idRacing uint) bool {

	_, err := ExistEventID(idEvent)
	if err == gorm.ErrRecordNotFound {
		return false
	}

	temp, err := searchRacingByIDandEventID(idRacing, idEvent)
	if err == gorm.ErrRecordNotFound {
		return false
	}

	DeleteAllHorses(temp.ID)
	GetDB().Delete(temp)

	return true
}

// DeleteAllRacesByEventID .
func DeleteAllRacesByEventID(eventID uint) bool {

	racings, err := searchAllRacesByEventID(eventID)
	if err != nil {
		return false
	}

	for _, race := range *racings {
		DeleteAllHorses(race.ID)
	}

	tempDelete := &[]Racing{}
	err = GetDB().Table("racings").Where("eventid = ?", eventID).Delete(tempDelete).Error
	if err == gorm.ErrRecordNotFound {
		return false
	}

	return true
}

// ---------------------------Validations------------------------------

// ExistRaceID .
func ExistRaceID(raceID uint) (*Racing, error) {
	temp := &Racing{}
	err := GetDB().Table("racings").Where("id = ?", raceID).First(&Racing{}).Error
	return temp, err
}

func searchRacingByIDandEventID(idRacing, idEvent uint) (*Racing, error) {
	temp := &Racing{}
	err := GetDB().Table("racings").Where("id = ? AND eventid = ?", idRacing, idEvent).First(&Racing{}).Error
	return temp, err
}

func searchAllRacesByEventID(eventID uint) (*[]Racing, error) {
	temp := &[]Racing{}
	err := GetDB().Table("racings").Where("eventid = ?", eventID).Find(&[]Racing{}).Error
	return temp, err
}
