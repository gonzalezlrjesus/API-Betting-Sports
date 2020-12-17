package models

import (
	u "github.com/gonzalezlrjesus/API-Betting-Sports/utils"

	"github.com/jinzhu/gorm"
)

// Retiro struct
type Retiro struct {
	gorm.Model
	Amount                   float64 `json:"amount"`
	Clientidentificationcard string  `json:"clientidentificationcard"`
}

// Dowithdrawal Client db
func (retiro *Retiro) Dowithdrawal(clientIDCard string) map[string]interface{} {
	retiro.Clientidentificationcard = clientIDCard
	if resp, ok := retiro.validateRetiro(); !ok {
		return resp
	}

	temp, err := ExistClientinCoinsDB(retiro.Clientidentificationcard)
	if err == gorm.ErrRecordNotFound {
		return nil
	}

	GetDB().Create(retiro)

	response := u.Message(true, "deposit has been created")
	response["update-coins"] = temp.DecreaseCoins(retiro.Amount)
	response["withdrawal"] = retiro
	return response
}

// GetAllWithdrawal Client
func GetAllWithdrawal(clientIDCard *string) map[string]interface{} {

	//check deposits ALL in DB
	temp := &[]Retiro{}
	err := GetDB().Table("retiros").Where("clientidentificationcard = ?", *clientIDCard).Order("id").Find(temp).Error
	if err == gorm.ErrRecordNotFound {
		return nil
	}

	response := u.Message(true, "Get all retiros")
	response["all-retiros"] = temp
	return response
}

// UpdateIdentificationClientRetiro  all record in table retiros
func UpdateIdentificationClientRetiro(clientIDCard, newIdentification string) map[string]interface{} {

	temp := &[]Retiro{}
	err := GetDB().Table("retiros").Where("clientidentificationcard = ?", clientIDCard).Update("clientidentificationcard", newIdentification).Error
	if err == gorm.ErrRecordNotFound {
		return nil
	}

	response := u.Message(true, "client all retiros has updated its identification")
	response["allretiros"] = temp
	return response
}

// ---------------------------Validations------------------------------

func (retiro *Retiro) validateRetiro() (map[string]interface{}, bool) {

	if retiro.Amount <= 0 {
		return u.Message(false, "Amount is required"), false
	}

	_, err := ExistClientIdentificationDB(retiro.Clientidentificationcard)
	if err == gorm.ErrRecordNotFound {
		return u.Message(false, "Client not exist"), false
	}

	return u.Message(false, "Requirement passed"), true
}
