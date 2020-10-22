package models

import (
	"fmt"

	u "github.com/gonzalezlrjesus/API-Betting-Sports/utils"

	"github.com/jinzhu/gorm"
)

// AuctionNumber struct
type AuctionNumber struct {
	gorm.Model
	Racingcomponentid uint `json:"racingcomponentid"`
	Valueauction      uint `json:"valueauction"`
}

// AddAuctionNumber auction number for auction to horse db
func (auctionNumber *AuctionNumber) AddAuctionNumber() map[string]interface{} {
	if resp, ok := auctionNumber.ValidateNumber(); !ok {
		return resp
	}

	GetDB().Create(auctionNumber)

	response := u.Message(true, "Auction number has been created")
	response["addAuctionNumber"] = auctionNumber
	return response
}

// GetAuctionNumbers auction Numbers all
func GetAuctionNumbers(idComponent *string) map[string]interface{} {

	temp := &[]AuctionNumber{}

	//check all auction number of racing component
	err := GetDB().Table("auction_numbers").Where("racingcomponentid = ?", *idComponent).Find(temp).Error
	fmt.Println("err", err)
	if err == gorm.ErrRecordNotFound {
		fmt.Println("auction number ALL : ", err)
		return u.Message(true, "There is not component with this ID")
	}

	response := u.Message(true, "Get all auction number")
	response["auctionNumbers"] = temp
	return response
}

// DeleteAuctionNumber client database
func DeleteAuctionNumber(idComponent, idauctionnumber *string) bool {

	temp := &AuctionNumber{}

	//check auctionNumber in DB
	err := GetDB().Table("auction_numbers").Where("id= ? AND racingcomponentid = ?", *idauctionnumber, *idComponent).First(temp).Error
	fmt.Println("err", err)
	if err != nil {
		fmt.Println("auctionNumber : ", err)
		return false
	}
	GetDB().Delete(temp)
	return true
}

// DeleteAuctionNumbers delete all auction numbers
func DeleteAuctionNumbers(idComponent uint) bool {

	temp := &[]AuctionNumber{}

	// delete all auction numbers
	err := GetDB().Table("auction_numbers").Where("Racingcomponentid = ?", idComponent).Delete(temp).Error
	if err == gorm.ErrRecordNotFound {
		fmt.Println("Racingcomponentid  : ", err)
		return false
	}

	return true
}

// ---------------------------Validations------------------------------
// I must validate that there is not a number exactly to a number just get of user front

// ValidateNumber Validate number to auction
func (auctionNumber *AuctionNumber) ValidateNumber() (map[string]interface{}, bool) {

	if auctionNumber.Valueauction <= 0 {
		return u.Message(false, "Valueauction is required"), false
	}

	temp := &AuctionNumber{}
	//check Racing Component in DB
	err := GetDB().Table("racing_components").Where("id = ?", auctionNumber.Racingcomponentid).First(temp).Error
	if err == gorm.ErrRecordNotFound {
		fmt.Println("Racing Component", err)
		return u.Message(false, "Racing Component exist no in DB"), false
	}

	return u.Message(false, "Requirement passed"), true
}
