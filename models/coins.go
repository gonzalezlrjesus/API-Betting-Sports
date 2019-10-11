package models

import (
	"fmt"
	u "API-Betting-Sports/utils"

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

// DecreaseCoins of client in DB
func (coins *Coins) DecreaseCoins(rematesAmount float64) map[string]interface{} {

	coins.Amount = ((coins.Amount) - rematesAmount)
	fmt.Printf("%f\n", coins.Amount)
	GetDB().Save(&coins)

	response := u.Message(true, "client has updated your coins")
	response["coins"] = coins
	return response

}

// UpdateIdentificationCoinClient client identification in coins table of DB
func UpdateIdentificationCoinClient(Clientidentificationcard, newIdentification string) map[string]interface{} {

	temp := &Coins{}

	//check Client identificaciont in DB
	err := GetDB().Table("coins").Where("clientidentificationcard = ?", Clientidentificationcard).First(temp).Error
	if err == gorm.ErrRecordNotFound {
		fmt.Println("client has not found : ", err)
		return u.Message(true, "client has NOT updated your coins")
	}

	temp.Clientidentificationcard = newIdentification
	GetDB().Save(&temp)

	response := u.Message(true, "client has updated your identification")
	response["coins"] = temp
	return response

}

// GetCoinsClient client
func GetCoinsClient(idClient *string) map[string]interface{} {
	temp := &Coins{}

	//check deposits ALL in DB
	err := GetDB().Table("coins").Where("clientidentificationcard = ?", *idClient).First(temp).Error
	if err == gorm.ErrRecordNotFound {
		fmt.Println("GetCoinsClient : ", err)
		return u.Message(true, "Client no exist")
	}

	response := u.Message(true, "Get Coins")
	response["coins"] = temp
	return response
}

// DeleteCoinsClient client deposit in database
func DeleteCoinsClient(Clientidentificationcard string) bool {

	temp := &[]Deposit{}

	//check Client coins in DB
	err := GetDB().Table("coins").Where("clientidentificationcard LIKE ?", Clientidentificationcard).Delete(temp).Error
	if err == gorm.ErrRecordNotFound {
		fmt.Println("Client coins ALL : ", err)
		return false
	}

	return true
}

// ---------------------------Validations------------------------------

// ValidateCoins struct that Front-End to Back-End
func (coins *Coins) ValidateCoins() (map[string]interface{}, bool) {

	if coins.Clientidentificationcard == "" {
		return u.Message(false, "User is not recognized"), false
	}

	return u.Message(false, "Requirement passed"), true
}
