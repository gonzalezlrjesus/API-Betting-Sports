package models

import (
	"fmt"
	u "test-golang/utils"

	"github.com/jinzhu/gorm"
)

// RacingComponents struct
type RacingComponents struct {
	gorm.Model
	Racingid        uint   `json:"racingid"`
	Stateracing     string `json:"stateracing"`
	Minimumamount   uint   `json:"minimunamount"`
	Percentage      uint   `json:"percentage"`
	Timeduration    uint   `json:"timeduration"`
	Horsenamestate  bool   `json:"horsenamestate"`
	Horsenotauction string `json:"horsenotauction"`
	Jockeynamestate bool   `json:"jockeynamestate"`
}

// CreateRacingComponents Racing Components db
func CreateRacingComponents(RacingID uint) map[string]interface{} {
	racingComponents := &RacingComponents{
		Racingid:        RacingID,
		Stateracing:     "En Espera",
		Minimumamount:   1000.00,
		Percentage:      20,
		Timeduration:    5,
		Horsenamestate:  true,
		Horsenotauction: "CASA",
		Jockeynamestate: true,
	}

	GetDB().Create(racingComponents)

	response := u.Message(true, "racing Components is add to Racing")
	response["racingComponents"] = racingComponents
	return response
}

// UpdateRacingComponents update all racing components DB
func (racingcomponents *RacingComponents) UpdateRacingComponents(RacingID string) map[string]interface{} {

	temp := &RacingComponents{}

	//check racing components in DB
	err := GetDB().Table("racing_components").Where("racingid = ?", RacingID).First(temp).Error
	if err == gorm.ErrRecordNotFound {
		fmt.Println("Racing has not found : ", err)
		return u.Message(true, "Racing Components has NOT updated")
	}

	temp.Stateracing = racingcomponents.Stateracing
	temp.Minimumamount = racingcomponents.Minimumamount
	temp.Percentage = racingcomponents.Percentage
	temp.Timeduration = racingcomponents.Timeduration
	temp.Horsenamestate = racingcomponents.Horsenamestate
	temp.Horsenotauction = racingcomponents.Horsenotauction
	temp.Jockeynamestate = racingcomponents.Jockeynamestate

	GetDB().Save(&temp)

	response := u.Message(true, "racing components has updated")
	response["Components"] = temp
	return response

}

// GetRacingComponents Racing Components
func GetRacingComponents(RacingID *string) map[string]interface{} {

	tempRacing := &Racing{}

	//check racing in DB
	errRacing := GetDB().Table("racings").Where("id = ?", *RacingID).First(tempRacing).Error
	if errRacing == gorm.ErrRecordNotFound {
		fmt.Println("Racing has not found : ", errRacing)
		return u.Message(true, "Racing not exist")
	}

	temp := &RacingComponents{}

	//check racing components in DB
	err := GetDB().Table("racing_components").Where("racingid = ?", *RacingID).First(temp).Error
	if err == gorm.ErrRecordNotFound {
		fmt.Println("Racing has not found : ", err)
		return u.Message(true, "Racing Components not exist")
	}

	response := u.Message(true, "Get Racing Components")
	response["Components"] = temp
	return response
}

// DeleteRacingComponents racing compoenents in DB
func DeleteRacingComponents(RacingID uint) bool {

	temp := &RacingComponents{}

	//check racing components in DB
	err := GetDB().Table("racing_components").Where("racingid = ?", RacingID).First(temp).Error
	if err == gorm.ErrRecordNotFound {
		fmt.Println("Racing has not found : ", err)
		return false
	}
	GetDB().Delete(temp)
	DeleteAuctionNumbers(temp.ID)
	return true
}
