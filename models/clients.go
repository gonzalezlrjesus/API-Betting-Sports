package models

import (
	u "github.com/gonzalezlrjesus/API-Betting-Sports/utils"
	"fmt"
	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

// Client struct
type Client struct {
	gorm.Model
	Name               string `json:"name"`
	Identificationcard string `json:"identificationcard"`
	Seudonimo          string `json:"seudonimo"`
	Password 		   string `json:"password"`
	Email              string `json:"email"`
	Phone              string `json:"phone"`
	Banknumber         string `json:"banknumber"`
	Bankname           string `json:"bankname"`
	State              string `json:"state"`
	Token              string `json:"token";sql:"-"`
}

// CreateClient Client db
func (client *Client) CreateClient() map[string]interface{} {

	if resp, ok := client.ValidateClient(); !ok {
		return resp
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(client.Password), bcrypt.DefaultCost)
	client.Password = string(hashedPassword)

	GetDB().Create(client)

	if client.ID <= 0 {
		return u.Message(false, "Failed to create client, connection error.")
	}

	//Create new JWT token for the newly registered client
	tk := &Token{UserID: client.ID}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, _ := token.SignedString([]byte(os.Getenv("token_password")))
	client.Token = tokenString

	client.Password = "" //delete password

	response := u.Message(true, "client has been created")
	response["Amount"] = CreateAmount(client.Identificationcard)
	response["client"] = client
	return response
}

// LoginClient function user
func LoginClient(password string, seudonimo string) map[string]interface{} {

	client := &Client{}
	err := GetDB().Table("clients").Where("seudonimo = ? ", seudonimo).First(client).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			fmt.Println("er:", err)
			return u.Message(false, "seudonimo not found or not identificationcard")
		}
		return u.Message(false, "Connection error. Please retry")
	}

	err = bcrypt.CompareHashAndPassword([]byte(client.Password), []byte(password))
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword { //Password does not match!
		return u.Message(false, "seudonimo not found or not identificationcard")
	}
	//Worked! Logged In
	client.Password = ""

	//Create JWT token
	tk := &Token{UserID: client.ID}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, err := token.SignedString([]byte(os.Getenv("token_password")))
	if err != nil {

	}
	client.Token = tokenString //Store the token in the response

	resp := u.Message(true, "Logged In")
	resp["client"] = client
	return resp
}

// GetClient client that request for him identification
func GetClient(idClient *string) (*Client, string) {

	client := &Client{}
	err := GetDB().Table("clients").Where("id = ?", *idClient).First(client).Error
	if err != nil {
		return nil, "Failed"
	}

	if client.Identificationcard == "" { //User not found!
		return nil, "Failed"
	}
	return client, "success"
}

// GetClients all client of table clients
func GetClients() []*Client {

	clients := make([]*Client, 0)
	err := GetDB().Table("clients").Order("seudonimo").Find(&clients).Error
	if err != nil {
		// fmt.Println(err)
		return nil
	}

	return clients
}

// UpdateClient client in DB
func (client *Client) UpdateClient(idClient *string) map[string]interface{} {

	if resp, ok := client.ValidateClientParams(idClient); !ok {
		return resp
	}

	temp := &Client{}

	//check client in DB
	err := GetDB().Table("clients").Where("id = ?", idClient).First(temp).Error
	if err == gorm.ErrRecordNotFound {
		fmt.Println(err)
		return nil
	}

	temp.Email = client.Email
	temp.Name = client.Name
	temp.Seudonimo = client.Seudonimo

	if len(client.Password) > 0 { 
		hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(client.Password), bcrypt.DefaultCost)
		temp.Password = string(hashedPassword)
	}
	
	temp.Phone = client.Phone
	temp.Banknumber = client.Banknumber
	temp.Bankname = client.Bankname
	temp.State = client.State

	response := u.Message(true, "client has been updated")
	if temp.Identificationcard != client.Identificationcard {
		response["Coins"] = UpdateIdentificationCoinClient(temp.Identificationcard, client.Identificationcard)
		response["Deposit"] = UpdateIdentificationClientDeposit(temp.Identificationcard, client.Identificationcard)
		response["Retiros"] = UpdateIdentificationClientRetiro(temp.Identificationcard, client.Identificationcard)
		temp.Identificationcard = client.Identificationcard
	}
	GetDB().Save(&temp)
	response["client"] = client
	return response

}

// UpdateStateClient client in DB
func UpdateStateClient(idClient *string, setStateClient *Client) map[string]interface{} {

	temp := &Client{}

	//check client in DB
	err := GetDB().Table("clients").Where("id = ?", *idClient).First(temp).Error
	if err == gorm.ErrRecordNotFound {
		fmt.Println(err)
		return nil
	}

	if setStateClient.State == "Habilitado" {
		temp.State = "Bloqueado"
	} else {
		temp.State = "Habilitado"
	}

	GetDB().Save(&temp)

	response := u.Message(true, "client has updated your state")
	response["client"] = temp
	return response

}

// DeleteClient from DB  Delete client then delete All client deposits and client coin.
func DeleteClient(idClient *string) bool {

	temp := &Client{}
	// Select records
	err := GetDB().Table("clients").Where("id= ?", *idClient).First(temp).Error

	if err != nil || err == gorm.ErrRecordNotFound {
		return false
	}
	// Delete all client deposit
	DeleteDepositsClient(temp.Identificationcard)

	// Delete all client coins
	DeleteCoinsClient(temp.Identificationcard)

	// Delete it
	GetDB().Delete(temp)

	return true
}

// ---------------------------Validations------------------------------

// ValidateClient struct that Front-End to Back-End
func (client *Client) ValidateClient() (map[string]interface{}, bool) {

	if client.Name == "" {
		return u.Message(false, "Name is required"), false
	}
	//Identificationcard must be unique
	temp := &Client{}

	if client.Identificationcard == "" {
		return u.Message(false, "Identificationcard is required"), false
	}

	if client.Seudonimo == "" {
		return u.Message(false, "Seudonimo is required"), false
	}

	if len(client.Password) < 6 {
		return u.Message(false, "Password is required"), false
	}

	//check for errors and duplicate Identificationcard
	err := GetDB().Table("clients").Where("Identificationcard = ?", client.Identificationcard).First(temp).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return u.Message(false, "Connection error. Please retry"), false
	}

	if temp.Identificationcard != "" {
		return u.Message(false, "Identificationcard already in use by another user."), false
	}

	// Seudonimo must be unique
	tempSeudonimo := &Client{}

	//check for errors and duplicate Identificationcard
	errSeudonimo := GetDB().Table("clients").Where("seudonimo = ?", client.Seudonimo).First(tempSeudonimo).Error
	if errSeudonimo != nil && errSeudonimo != gorm.ErrRecordNotFound {
		return u.Message(false, "Connection error. Please retry"), false
	}

	if tempSeudonimo.Seudonimo != "" {
		return u.Message(false, "Seudonimo already in use by another user."), false
	}

	tempEmail := &Client{}

	//check for errors and duplicate Email
	errEmail := GetDB().Table("clients").Where("email = ?", client.Email).First(tempEmail).Error
	if errEmail != nil && errEmail != gorm.ErrRecordNotFound {
		return u.Message(false, "Connection error. Please retry"), false
	}

	if tempEmail.Email != "" {
		return u.Message(false, "Email already in use by another user."), false
	}

	return u.Message(false, "Requirement passed"), true
}

// ValidateClientParams struct Params for Update Client
func (client *Client) ValidateClientParams(idClient *string) (map[string]interface{}, bool) {
	tempParam := &Client{}

	if client.Identificationcard == "" {
		return u.Message(false, "Identificationcard is required"), false
	}

	if client.Name == "" {
		return u.Message(false, "Name is required"), false
	}

	if client.Seudonimo == "" {
		return u.Message(false, "Seudonimo is required"), false
	}

	// Data Param
	err := GetDB().Table("clients").Where("id = ?", *idClient).First(tempParam).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return u.Message(false, "Connection error. Please retry"), false
	}

	if err == gorm.ErrRecordNotFound {
		return u.Message(false, "Not found client ID Param"), false
	}

	// Data form
	tempForm := &Client{}
	errAux := GetDB().Table("clients").Where("Identificationcard = ?", client.Identificationcard).First(tempForm).Error
	if errAux != nil && errAux != gorm.ErrRecordNotFound {
		return u.Message(false, "Connection error. Please retry"), false
	}

	// Data form
	tempSeudonimo := &Client{}
	errSeudonimo := GetDB().Table("clients").Where("seudonimo = ?", client.Seudonimo).First(tempSeudonimo).Error
	if errSeudonimo != nil && errSeudonimo != gorm.ErrRecordNotFound {
		return u.Message(false, "Connection error. Please retry"), false
	}

	// Data form email
	tempEmail := &Client{}
	errEmail := GetDB().Table("clients").Where("email = ?", client.Email).First(tempEmail).Error
	if errEmail != nil && errEmail != gorm.ErrRecordNotFound {
		return u.Message(false, "Connection error. Please retry"), false
	}

	if errAux != gorm.ErrRecordNotFound && tempForm.Identificationcard != tempParam.Identificationcard {
		return u.Message(false, "there is client with this identification to send in Form"), false
	}
	if errSeudonimo != gorm.ErrRecordNotFound && tempSeudonimo.Seudonimo != tempParam.Seudonimo {
		return u.Message(false, "there is client with this Seudonimo to send in Form"), false
	}
	if errEmail != gorm.ErrRecordNotFound && tempEmail.Email != tempParam.Email {
		return u.Message(false, "there is client with this Email to send in Form"), false
	}

	return u.Message(false, "Requirement passed"), true
}
