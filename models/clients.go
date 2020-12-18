package models

import (
	"strings"

	u "github.com/gonzalezlrjesus/API-Betting-Sports/utils"

	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

// Client struct
type Client struct {
	gorm.Model
	Name               string `json:"name"`
	Identificationcard string `json:"identificationcard"`
	Seudonimo          string `json:"seudonimo"`
	Password           string `json:"password"`
	Email              string `json:"email"`
	Phone              string `json:"phone"`
	Banknumber         string `json:"banknumber"`
	Bankname           string `json:"bankname"`
	State              string `json:"state"`
	Token              string `json:"token";sql:"-"`
}

// CreateClient .
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

	client.Token = createToken(client.ID)
	client.Password = "" //delete password

	response := u.Message(true, "client has been created")
	response["Amount"] = CreateAmount(client.Identificationcard)
	response["client"] = client
	return response
}

// LoginClient function user
func LoginClient(password string, seudonimo string) map[string]interface{} {

	client, err := ExistClientSeudonimonDB(seudonimo)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return u.Message(false, "seudonimo not found or not identificationcard")
		}
		return u.Message(false, "Connection error. Please retry")
	}

	err = bcrypt.CompareHashAndPassword([]byte(client.Password), []byte(password))
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword { //Password does not match!
		return u.Message(false, "seudonimo not found or not identificationcard")
	}

	client.Token = createToken(client.ID)
	client.Password = ""

	resp := u.Message(true, "Logged In")
	resp["client"] = client
	return resp
}

// GetClient client that request for him identification
func GetClient(idClient *string) []*Client {

	client := make([]*Client, 0)
	GetDB().Table("clients").Where("id = ?", *idClient).Order("id").Find(&client)

	for _, s := range client {
		s.Password = ""
	}
	return client
}

// GetClients all client of table clients
func GetClients() []*Client {

	clients := make([]*Client, 0)
	GetDB().Table("clients").Order("id").Find(&clients)

	for _, s := range clients {
		s.Password = ""
	}

	return clients
}

// UpdateClient client in DB
func (client *Client) UpdateClient(idClient *string) map[string]interface{} {

	if resp, ok := client.ValidateClientParams(idClient); !ok {
		return resp
	}

	temp, err := ExistClientDB(idClient)
	if err == gorm.ErrRecordNotFound {
		return nil
	}

	temp.Email = client.Email
	temp.Name = client.Name
	temp.Seudonimo = client.Seudonimo
	temp.Phone = client.Phone
	temp.Banknumber = client.Banknumber
	temp.Bankname = client.Bankname
	temp.State = client.State
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(client.Password), bcrypt.DefaultCost)
	temp.Password = string(hashedPassword)

	response := u.Message(true, "client has been updated")
	if temp.Identificationcard != client.Identificationcard {
		response["Coins"] = UpdateIdentificationCoinClient(temp.Identificationcard, client.Identificationcard)
		response["Deposit"] = UpdateIdentificationClientDeposit(temp.Identificationcard, client.Identificationcard)
		response["Retiros"] = UpdateIdentificationClientRetiro(temp.Identificationcard, client.Identificationcard)
		temp.Identificationcard = client.Identificationcard
	}
	GetDB().Save(&temp)
	temp.Password = ""
	response["client"] = temp
	return response

}

// UpdateStateClient client in DB
func UpdateStateClient(idClient *string) map[string]interface{} {

	temp, err := ExistClientDB(idClient)
	if err == gorm.ErrRecordNotFound {
		return nil
	}

	if temp.State == "Habilitado" {
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

	temp, err := ExistClientDB(idClient)
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

	if resp, ok := client.validateFields(); !ok {
		return resp, ok
	}

	tempIdentification, tempSeudonimo, tempEmail, resp, ok := client.validateEqualUserData()
	if !ok {
		return resp, ok
	}

	if tempIdentification.Identificationcard != "" {
		return u.Message(false, "Identificationcard already in use by another user."), false
	}

	if tempSeudonimo.Seudonimo != "" {
		return u.Message(false, "Seudonimo already in use by another user."), false
	}

	if tempEmail.Email != "" {
		return u.Message(false, "Email already in use by another user."), false
	}

	return u.Message(false, "Requirement passed"), true
}

// ValidateClientParams struct Params for Update Client
func (client *Client) ValidateClientParams(idClient *string) (map[string]interface{}, bool) {

	if resp, ok := client.validateFields(); !ok {
		return resp, ok
	}

	tempParam, err := ExistClientDB(idClient)
	if err != nil && err != gorm.ErrRecordNotFound {
		return u.Message(false, "Connection error. Please retry"), false
	}

	if err == gorm.ErrRecordNotFound {
		return u.Message(false, "Not found client ID Param"), false
	}

	// Validate users with the same: email, identification, and seudonimo

	tempIdentification, tempSeudonimo, tempEmail, resp, ok := client.validateEqualUserData()
	if !ok {
		return resp, ok
	}

	if tempIdentification.Identificationcard != "" && tempIdentification.Identificationcard != tempParam.Identificationcard {
		return u.Message(false, "there is client with this identification to send in Form"), false
	}
	if tempSeudonimo.Seudonimo != "" && tempSeudonimo.Seudonimo != tempParam.Seudonimo {
		return u.Message(false, "there is client with this Seudonimo to send in Form"), false
	}
	if tempEmail.Email != "" && tempEmail.Email != tempParam.Email {
		return u.Message(false, "there is client with this Email to send in Form"), false
	}

	return u.Message(false, "Requirement passed"), true
}

func (client *Client) validateFields() (map[string]interface{}, bool) {

	if client.Identificationcard == "" {
		return u.Message(false, "Identificationcard is required"), false
	}

	if client.Name == "" {
		return u.Message(false, "Name is required"), false
	}

	if client.Seudonimo == "" {
		return u.Message(false, "Seudonimo is required"), false
	}

	if client.State == "" {
		return u.Message(false, "State is required"), false
	}

	if !strings.Contains(client.Email, "@") {
		return u.Message(false, "Email address is required"), false
	}

	if len(client.Password) < 6 {
		return u.Message(false, "Password is required"), false
	}

	return u.Message(false, "Requirement passed"), true
}

func (client *Client) validateEqualUserData() (*Client, *Client, *Client, map[string]interface{}, bool) {
	// tempIdentification form
	tempForm, errAux := ExistClientIdentificationDB(client.Identificationcard)
	if errAux != nil && errAux != gorm.ErrRecordNotFound {
		return nil, nil, nil, u.Message(false, "Connection error. Please retry"), false
	}

	// tempSeudonimo form
	tempSeudonimo, errSeudonimo := ExistClientSeudonimonDB(client.Seudonimo)
	if errSeudonimo != nil && errSeudonimo != gorm.ErrRecordNotFound {
		return nil, nil, nil, u.Message(false, "Connection error. Please retry"), false
	}

	// tempEmail form
	tempEmail := &Client{}
	errEmail := GetDB().Table("clients").Where("email = ?", client.Email).First(tempEmail).Error
	if errEmail != nil && errEmail != gorm.ErrRecordNotFound {
		return nil, nil, nil, u.Message(false, "Connection error. Please retry"), false
	}

	return tempForm, tempSeudonimo, tempEmail, u.Message(false, "Requirement passed"), true
}

// ExistClientDB .
func ExistClientDB(idClient *string) (*Client, error) {
	temp := &Client{}
	err := GetDB().Table("clients").Where("id = ?", *idClient).First(temp).Error
	return temp, err
}

// ExistClientIdentificationDB .
func ExistClientIdentificationDB(clientIdentification string) (*Client, error) {
	temp := &Client{}
	err := GetDB().Table("clients").Where("Identificationcard = ?", clientIdentification).First(temp).Error
	return temp, err
}

// ExistClientSeudonimonDB .
func ExistClientSeudonimonDB(seudonimo string) (*Client, error) {
	temp := &Client{}
	err := GetDB().Table("clients").Where("seudonimo = ?", seudonimo).First(temp).Error
	return temp, err
}
