package models

import (
	u "github.com/gonzalezlrjesus/API-Betting-Sports/utils"

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

// UpdateCoins of client
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
	GetDB().Save(&coins)
	response := u.Message(true, "client has updated your coins")
	response["coins"] = coins
	return response
}

// UpdateIdentificationCoinClient client identification in coins table of DB
func UpdateIdentificationCoinClient(Clientidentificationcard, newIdentification string) map[string]interface{} {

	temp, err := ExistClientinCoinsDB(Clientidentificationcard)
	if err == gorm.ErrRecordNotFound {
		return u.Message(true, "Client not exist")
	}

	temp.Clientidentificationcard = newIdentification
	GetDB().Save(&temp)

	response := u.Message(true, "client has updated your identification")
	response["coins"] = temp
	return response

}

// GetCoinsClient client
func GetCoinsClient(Clientidentificationcard string) map[string]interface{} {

	temp, err := ExistClientinCoinsDB(Clientidentificationcard)
	if err == gorm.ErrRecordNotFound {
		return u.Message(true, "Client not exist")
	}

	response := u.Message(true, "Get Coins")
	response["coins"] = temp
	return response
}

// DeleteCoinsClient client deposit in database
func DeleteCoinsClient(Clientidentificationcard string) bool {

	//check client exist and delete it
	temp := &[]Deposit{}
	err := GetDB().Table("coins").Where("clientidentificationcard LIKE ?", Clientidentificationcard).Delete(temp).Error
	if err == gorm.ErrRecordNotFound {
		return false
	}

	return true
}

// ---------------------------Validations------------------------------

// ValidateCoins struct
func (coins *Coins) ValidateCoins() (map[string]interface{}, bool) {

	if coins.Clientidentificationcard == "" {
		return u.Message(false, "User is not recognized"), false
	}

	return u.Message(false, "Requirement passed"), true
}

// ExistClientinCoinsDB .
func ExistClientinCoinsDB(clientIDCard string) (*Coins, error) {
	temp := &Coins{}
	err := GetDB().Table("coins").Where("ClientIdentificationcard = ?", clientIDCard).First(temp).Error
	return temp, err
}
