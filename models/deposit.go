package models

import (
	"fmt"
	u "test-golang/utils"

	"github.com/jinzhu/gorm"
)

// Deposit struct
type Deposit struct {
	gorm.Model
	Amount                   float64 `json:"amount"`
	Clientidentificationcard string  `json:"clientidentificationcard"`
	FormaPago                string  `json:"formapago"`
	Serial                   string  `json:"serial"`
}

// AddDepositClient Client db
func (deposit *Deposit) AddDepositClient() map[string]interface{} {

	if resp, ok := deposit.ValidateDeposit(); !ok {
		return resp
	}

	GetDB().Create(deposit)
	temp := &Coins{Clientidentificationcard: deposit.Clientidentificationcard}

	//check client_Coins in DB
	err := GetDB().Table("coins").Where("ClientIdentificationcard = ?", temp.Clientidentificationcard).First(temp).Error
	if err == gorm.ErrRecordNotFound {
		fmt.Println("Client Coins : ", err)
		return nil
	}

	response := u.Message(true, "deposit has been created")
	response["updateCoins"] = temp.UpdateCoins(deposit.Amount)
	response["deposit"] = deposit
	return response
}

// GetAllDepositsClient Client db
func GetAllDepositsClient(idClient *string) map[string]interface{} {

	temp := &[]Deposit{}

	//check deposits ALL in DB
	err := GetDB().Table("deposits").Where("ClientIdentificationcard = ?", *idClient).Find(temp).Error
	if err == gorm.ErrRecordNotFound {
		fmt.Println("deposits ALL : ", err)
		return nil
	}

	response := u.Message(true, "deposit has been created")
	response["alldeposits"] = temp
	return response
}

// ---------------------------Validations------------------------------

// ValidateDeposit struct that Front-End to Back-End
func (deposit *Deposit) ValidateDeposit() (map[string]interface{}, bool) {

	if deposit.Amount <= 0 {
		return u.Message(false, "Amount is required"), false
	}

	if deposit.FormaPago == "" {
		return u.Message(false, "FormaPago is required"), false
	}

	temp := &Client{}
	//check client in DB
	err := GetDB().Table("clients").Where("Identificationcard = ?", deposit.Clientidentificationcard).First(temp).Error
	if err == gorm.ErrRecordNotFound {
		fmt.Println("DEPosit", err)
		return u.Message(false, "Client exist no in DB"), false
	}

	return u.Message(false, "Requirement passed"), true
}
