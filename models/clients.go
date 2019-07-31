package models

import (
	"fmt"
	"os"
	u "test-golang/utils"

	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
)

// Client struct
type Client struct {
	gorm.Model
	Name               string `json:"name"`
	Identificationcard string `json:"identificationcard"`
	Email              string `json:"email"`
	Phone              string `json:"phone"`
	Banknumber         string `json:"banknumber"`
	Bankname           string `json:"bankname"`
	State              string `json:"state"`
	Token              string `json:"token";sql:"-"`
}

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

	//check for errors and duplicate Identificationcard
	err := GetDB().Table("clients").Where("Identificationcard = ?", client.Identificationcard).First(temp).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return u.Message(false, "Connection error. Please retry"), false
	}

	if temp.Identificationcard != "" {
		return u.Message(false, "Identificationcard already in use by another user."), false
	}

	return u.Message(false, "Requirement passed"), true
}

// CreateClient Client db
func (client *Client) CreateClient() map[string]interface{} {

	if resp, ok := client.ValidateClient(); !ok {
		return resp
	}

	GetDB().Create(client)

	if client.ID <= 0 {
		return u.Message(false, "Failed to create client, connection error.")
	}

	//Create new JWT token for the newly registered client
	tk := &Token{UserID: client.ID}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, _ := token.SignedString([]byte(os.Getenv("token_password")))
	client.Token = tokenString

	response := u.Message(true, "client has been created")
	response["client"] = client
	return response
}

// LoginClient function user
func LoginClient(identificationcard string) map[string]interface{} {

	client := &Client{}
	err := GetDB().Table("clients").Where("Identificationcard = ?", identificationcard).First(client).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return u.Message(false, "Cedula not found")
		}
		return u.Message(false, "Connection error. Please retry")
	}

	//Create JWT token
	tk := &Token{UserID: client.ID}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, err := token.SignedString([]byte(os.Getenv("token_password")))
	if err != nil {
		fmt.Println("sas", err)
		fmt.Println(err)
	}
	client.Token = tokenString //Store the token in the response

	resp := u.Message(true, "Logged In")
	resp["client"] = client
	return resp
}

// GetClient from DB
func GetClient(idClient *string) (*Client, string) {

	acc := &Client{}
	err := GetDB().Table("clients").Where("Identificationcard = ?", *idClient).First(acc).Error
	if err != nil {
		fmt.Println(err)
		return nil, "Failed"
	}
	if acc.Identificationcard == "" { //User not found!
		return nil, "Failed"
	}
	return acc, "success"
}

// GetClients all db
func GetClients() []*Client {

	clients := make([]*Client, 0)

	err := GetDB().Table("clients").Find(&clients).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return clients
}

// UpdateClient from DB
func (client *Client) UpdateClient(idClient *string) map[string]interface{} {

	if resp, ok := client.ValidateClientParams(idClient); !ok {
		return resp
	}

	//Identificationcard must be unique
	temp := &Client{Identificationcard: *idClient}
	fmt.Println(temp.Identificationcard)

	//check for errors and duplicate Identificationcard
	err := GetDB().Table("clients").Where("Identificationcard = ?", temp.Identificationcard).First(temp).Error
	if err == gorm.ErrRecordNotFound {
		fmt.Println(err)
		return nil
	}

	temp.Email = client.Email
	temp.Identificationcard = client.Identificationcard
	temp.Name = client.Name
	temp.Phone = client.Phone
	temp.Banknumber = client.Banknumber
	temp.Bankname = client.Bankname
	temp.State = client.State

	GetDB().Save(&temp)

	response := u.Message(true, "client has been updated")
	response["client"] = client
	return response

}

// ValidateClientParams struct that Front-End to Back-End
func (client *Client) ValidateClientParams(idClient *string) (map[string]interface{}, bool) {
	temp := &Client{}

	if client.Identificationcard == "" {
		return u.Message(false, "Identificationcard is required"), false
	}

	if client.Name == "" {
		return u.Message(false, "Name is required"), false
	}

	err := GetDB().Table("clients").Where("Identificationcard = ?", *idClient).First(temp).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return u.Message(false, "Connection error. Please retry"), false
	}
	fmt.Println("validation")
	fmt.Println(err)
	if err == gorm.ErrRecordNotFound {
		return u.Message(false, "Not found client"), false
	}

	temp2 := &Client{}
	err2 := GetDB().Table("clients").Where("Identificationcard = ?", client.Identificationcard).First(temp2).Error
	if err2 != nil && err2 != gorm.ErrRecordNotFound {
		return u.Message(false, "Connection error. Please retry"), false
	}

	if err2 != gorm.ErrRecordNotFound && temp2.Identificationcard != *idClient {
		return u.Message(false, "there is a client with this identification"), false
	}

	fmt.Println(temp.Identificationcard)
	return u.Message(false, "Requirement passed"), true
}

// DeleteClient from DB
func DeleteClient(idClient *string) bool {

	temp := &Client{}
	// Select records and delete it
	err := GetDB().Table("clients").Where("Identificationcard= ?", *idClient).First(temp).Error

	if err != nil || err == gorm.ErrRecordNotFound {
		return false
	}
	GetDB().Delete(temp)
	return true

}
