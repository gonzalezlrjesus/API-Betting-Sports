package models

import (
	u "github.com/gonzalezlrjesus/API-Betting-Sports/utils"

	"github.com/jinzhu/gorm"
)

// Deposit struct
type Deposit struct {
	gorm.Model
	Amount                   float64 `json:"amount"`
	Clientidentificationcard string  `json:"clientidentificationcard"`
	FormaPago                string  `json:"formapago"`
	Serial                   uint    `json:"serial"`
}

// AddDepositClient Client db
func (deposit *Deposit) AddDepositClient() map[string]interface{} {

	if resp, ok := deposit.ValidateDeposit(); !ok {
		return resp
	}

	temp := &Coins{Clientidentificationcard: deposit.Clientidentificationcard}

	//check client in Coins table DB
	if !ExistClientinCoinsDB(temp.Clientidentificationcard) {
		return nil
	}

	GetDB().Create(deposit)
	response := u.Message(true, "deposit has been created")
	response["updateCoins"] = temp.UpdateCoins(deposit.Amount)
	response["deposit"] = deposit
	return response
}

// AddGananciaClient Client db
func AddGananciaClient(seudonimo string, montoganado int64, formapago string, serial uint) map[string]interface{} {

	//check client-seudonimo in DB
	temp, err := ExistClientSeudonimonDB(seudonimo)
	if err == gorm.ErrRecordNotFound {
		return nil
	}

	depositGanador := &Deposit{Amount: float64(montoganado),
		Clientidentificationcard: temp.Identificationcard,
		FormaPago:                formapago,
		Serial:                   serial}

	tempCoins := &Coins{Clientidentificationcard: depositGanador.Clientidentificationcard}

	//check client in Coins table DB
	if !ExistClientinCoinsDB(tempCoins.Clientidentificationcard) {
		return nil
	}

	GetDB().Create(depositGanador)
	response := u.Message(true, "deposit has been created")
	response["updateCoins"] = tempCoins.UpdateCoins(depositGanador.Amount)
	response["depositGanador"] = depositGanador
	return response
}

// GetAllDepositsClient Client db
func GetAllDepositsClient(idClient *string) map[string]interface{} {

	temp := &[]Deposit{}

	//check deposits ALL in DB
	err := GetDB().Table("deposits").Where("ClientIdentificationcard = ?", *idClient).Find(temp).Error
	if err == gorm.ErrRecordNotFound {
		return nil
	}

	response := u.Message(true, "All Deposits")
	response["depositList"] = temp
	return response
}

// UpdateIdentificationClientDeposit update  client ID in all deposits
func UpdateIdentificationClientDeposit(Clientidentificationcard, newIdentification string) map[string]interface{} {

	temp := &[]Deposit{}

	//check client identificacion in deposit table ALL in DB
	err := GetDB().Table("deposits").Where("ClientIdentificationcard = ?", Clientidentificationcard).
		Update("clientidentificationcard", newIdentification).Error
	if err == gorm.ErrRecordNotFound {
		return nil
	}

	response := u.Message(true, "client deposits has updated his identification")
	response["alldeposits"] = temp
	return response
}

// DeleteDepositsClient delete all client deposits in database
func DeleteDepositsClient(Clientidentificationcard string) bool {

	temp := &[]Deposit{}

	//check Client Deposits in DB
	err := GetDB().Table("deposits").Where("ClientIdentificationcard LIKE ?", Clientidentificationcard).
		Delete(temp).Error
	if err == gorm.ErrRecordNotFound {
		return false
	}

	return true
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

	//check client in DB
	_, err := ExistClientIdentificationDB(deposit.Clientidentificationcard)
	if err == gorm.ErrRecordNotFound {
		return u.Message(false, "Client exist no in DB"), false
	}

	return u.Message(false, "Requirement passed"), true
}
