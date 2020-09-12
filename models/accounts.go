package models

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
	"os"
	u "restAPI_contacts/utils"
	"strings"
)

type Token struct {
	UserId uint
	jwt.StandardClaims
}

type Account struct {
	gorm.Model
	Email string `json:"email"`
	Password string `json:"password"`
	Token string `json:"token";sql:"-"`
}

func (account *Account) Validate() (map[string] interface{}, bool) {

	if !strings.Contains(account.Email, "@") {
		return u.Message(false, "Validate Email..."), false
	}

	if len(account.Password) < 6 {
		return u.Message(false, "Validate Password..."), false
	}

	temp := &Account{}
	err := GetDB().Table("accounts").Where("email = ?", account.Email).First(temp).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		return u.Message(false, "Connection error..."), false
	}

	if temp.Email != "" {
		return u.Message(false, "Email already used by another user..."), false
	}

	return u.Message(false, "Requirement passed"), true
}

func (account *Account) Create() (map[string] interface{}) {

	if resp, ok := account.Validate(); !ok {
		return resp
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(account.Password), bcrypt.DefaultCost)
	account.Password = string(hashedPassword)

	GetDB().Create(account)

	if account.ID <= 0 {
		return u.Message(false, "Failed to create account...")
	}

	tk := &Token{UserId: account.ID}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, _ := token.SignedString([]byte(os.Getenv("token_password")))
	account.Token = tokenString

	account.Password = ""

	response := u.Message(true, "Account created...")
	response["account"] = account
	return response
}

func Login(email, password string) (map[string]interface{}) {
	account := &Account{}
	err := GetDB().Table("accounts").Where("email = ?", email).First(account).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return u.Message(false, "Email adress not found...")
		}
		return u.Message(false, "Connection lost...")
	}

	err = bcrypt.CompareHashAndPassword([]byte(account.Password), []byte(password))
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return u.Message(false, "Invalid password...")
	}

	account.Password = ""

	tk := &Token{UserId: account.ID}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, _ := token.SignedString([]byte(os.Getenv("token_password")))
	account.Token = tokenString

	resp := u.Message(true, "Logged in")
	resp["account"] = account
	return resp
}

func GetUser(u uint) *Account {
	account := &Account{}
	GetDB().Table("accounts").Where("id = ?", u).First(account)
	if account.Email == "" {
		return nil
	}

	account.Password = ""
	return account
}

