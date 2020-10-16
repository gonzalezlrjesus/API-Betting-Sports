package models

import (
	u "github.com/gonzalezlrjesus/API-Betting-Sports/utils"
	"fmt"

	"github.com/jinzhu/gorm"
)

// Retiro struct
type Retiro struct {
	gorm.Model
	Amount                   float64 `json:"amount"`
	Clientidentificationcard string  `json:"clientidentificationcard"`
}

// Dowithdrawal Client db
func (retiro *Retiro) Dowithdrawal() map[string]interface{} {

	if resp, ok := retiro.ValidateRetiro(); !ok {
		return resp
	}

	GetDB().Create(retiro)
	temp := &Coins{Clientidentificationcard: retiro.Clientidentificationcard}

	//check client_Coins in DB
	err := GetDB().Table("coins").Where("clientIdentificationcard = ?", temp.Clientidentificationcard).First(temp).Error
	if err == gorm.ErrRecordNotFound {
		fmt.Println("Client Coins : ", err)
		return nil
	}

	response := u.Message(true, "deposit has been created")
	response["updateCoins"] = temp.WithdrawalCoins(retiro.Amount)
	response["withdrawal"] = retiro
	return response
}

// GetAllWithdrawal Client db
func GetAllWithdrawal(idClient *string) map[string]interface{} {

	temp := &[]Deposit{}

	//check deposits ALL in DB
	err := GetDB().Table("retiros").Where("clientidentificationcard = ?", *idClient).Find(temp).Error
	if err == gorm.ErrRecordNotFound {
		fmt.Println("Retiros: ", err)
		return nil
	}

	response := u.Message(true, "Get all retiros")
	response["allretiros"] = temp
	return response
}

// UpdateIdentificationClientRetiro update identificacionCard of client all record in table deposits of db
func UpdateIdentificationClientRetiro(Clientidentificationcard, newIdentification string) map[string]interface{} {

	temp := &[]Retiro{}

	//check client identificacion in retiros table ALL in DB
	err := GetDB().Table("retiros").Where("clientidentificationcard = ?", Clientidentificationcard).Update("clientidentificationcard", newIdentification).Error

	if err == gorm.ErrRecordNotFound {
		fmt.Println("deposits ALL : ", err)
		return nil
	}

	response := u.Message(true, "client retiros has updated your identification")
	response["allretiros"] = temp
	return response
}

// ---------------------------Validations------------------------------

// ValidateRetiro struct that Front-End to Back-End
func (retiro *Retiro) ValidateRetiro() (map[string]interface{}, bool) {

	if retiro.Amount <= 0 {
		return u.Message(false, "Amount is required"), false
	}

	temp := &Client{}
	//check client in DB
	err := GetDB().Table("clients").Where("identificationcard = ?", retiro.Clientidentificationcard).First(temp).Error
	if err == gorm.ErrRecordNotFound {
		fmt.Println("Retiros", err)
		return u.Message(false, "Client exist no in DB"), false
	}

	return u.Message(false, "Requirement passed"), true
}
