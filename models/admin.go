package models

import (
	"fmt"
	"os"
	"strings"
	u "API-Betting-Sports/utils"

	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

// Token JWT claims struct
type Token struct {
	UserID uint
	jwt.StandardClaims
}

// Admin struct
type Admin struct {
	gorm.Model
	Email    string `json:"email"`
	Password string `json:"password"`
	Token    string `json:"token";sql:"-"`
}

// Validate struct that Front-End to Back-End
func (admin *Admin) Validate() (map[string]interface{}, bool) {

	if !strings.Contains(admin.Email, "@") {
		return u.Message(false, "Email address is required"), false
	}

	if len(admin.Password) < 6 {
		return u.Message(false, "Password is required"), false
	}

	//Email must be unique
	temp := &Admin{}

	//check for errors and duplicate emails
	err := GetDB().Table("admins").Where("email = ?", admin.Email).First(temp).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return u.Message(false, "Connection error. Please retry"), false
	}
	if temp.Email != "" {
		return u.Message(false, "Email address already in use by another user."), false
	}

	return u.Message(false, "Requirement passed"), true
}

// Create accoutn db
func (admin *Admin) Create() map[string]interface{} {

	if resp, ok := admin.Validate(); !ok {
		return resp
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(admin.Password), bcrypt.DefaultCost)
	admin.Password = string(hashedPassword)

	GetDB().Create(admin)

	if admin.ID <= 0 {
		return u.Message(false, "Failed to create admin, connection error.")
	}

	//Create new JWT token for the newly registered admin
	tk := &Token{UserID: admin.ID}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, _ := token.SignedString([]byte(os.Getenv("token_password")))
	admin.Token = tokenString

	admin.Password = "" //delete password

	response := u.Message(true, "admin has been created")
	response["admin"] = admin
	return response
}

// Login function user
func Login(email, password string) map[string]interface{} {

	account := &Admin{}
	err := GetDB().Table("admins").Where("email = ?", email).First(account).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return u.Message(false, "Email address not found")
		}
		return u.Message(false, "Connection error. Please retry")
	}

	err = bcrypt.CompareHashAndPassword([]byte(account.Password), []byte(password))
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword { //Password does not match!
		return u.Message(false, "Invalid login credentials. Please try again")
	}
	//Worked! Logged In
	account.Password = ""

	//Create JWT token
	tk := &Token{UserID: account.ID}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, err := token.SignedString([]byte(os.Getenv("token_password")))
	if err != nil {
		// fmt.Println("sas", err)
		fmt.Println(err)
	}
	account.Token = tokenString //Store the token in the response

	resp := u.Message(true, "Logged In")
	resp["account"] = account
	return resp
}

// GetUser from DB
func GetUser(u uint) *Admin {

	acc := &Admin{}
	GetDB().Table("admins").Where("id = ?", u).First(acc)
	if acc.Email == "" { //User not found!
		return nil
	}

	acc.Password = ""
	return acc
}

// GetAdmins all db
func GetAdmins() []*Admin {

	admins := make([]*Admin, 0)
	err := GetDB().Table("admins").Find(&admins).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return admins
}
