package models

import (
	"fmt"
	u "test-golang/utils"

	"github.com/jinzhu/gorm"
)

// Coins struct
type Coins struct {
	gorm.Model
	Clientidentificationcard string  `json:"clientidentificationcard"`
	Amount                   float64 `json:"amount"`
}

// CreateAmount Client db
func CreateAmount(clientIdentificationcard string) map[string]interface{} {
	coins := &Coins{Clientidentificationcard: clientIdentificationcard, Amount: 0}

	if resp, ok := coins.ValidateCoins(); !ok {
		return resp
	}

	GetDB().Create(coins)

	response := u.Message(true, "Coins add to account")
	response["coins"] = coins
	return response
}

// UpdateCoins of client in DB
func (coins *Coins) UpdateCoins(amountDeposit float64) map[string]interface{} {

	coins.Amount = ((coins.Amount) + amountDeposit)

	GetDB().Save(&coins)

	response := u.Message(true, "client has updated your coins")
	response["coins"] = coins
	return response

}

// GetCoinsClient client
func GetCoinsClient(idClient *string) map[string]interface{} {
	temp := &Coins{}

	//check deposits ALL in DB
	err := GetDB().Table("coins").Where("ClientIdentificationcard = ?", *idClient).First(temp).Error
	if err == gorm.ErrRecordNotFound {
		fmt.Println("GetCoinsClient : ", err)
		return u.Message(true, "Client no exist")
	}

	response := u.Message(true, "Get Coins")
	response["coins"] = temp
	return response
}

// ---------------------------Validations------------------------------

// ValidateCoins struct that Front-End to Back-End
func (coins *Coins) ValidateCoins() (map[string]interface{}, bool) {

	if coins.Clientidentificationcard == "" {
		return u.Message(false, "User is not recognized"), false
	}

	return u.Message(false, "Requirement passed"), true
}
