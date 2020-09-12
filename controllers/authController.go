package controllers

import (
	"encoding/json"
	"net/http"
	"restAPI_contacts/models"
	u "restAPI_contacts/utils"
)

var CreateAccount = func(response http.ResponseWriter, r *http.Request) {

	account := &models.Account{}
	err := json.NewDecoder(r.Body).Decode(account)
	if err != nil {
		u.Respond(response, u.Message(false, "Invalid Request"))
		return
	}
	resp := account.Create()
	u.Respond(response, resp)
}

var Authenticate = func(response http.ResponseWriter, r *http.Request) {
	account := &models.Account{}
	err := json.NewDecoder(r.Body).Decode(account)
	if err != nil {
		u.Respond(response, u.Message(false, "Invalid Request"))
		return
	}
	resp := models.Login(account.Email, account.Password)
	u.Respond(response, resp)
}
