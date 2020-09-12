package controllers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"restAPI_contacts/models"
	u "restAPI_contacts/utils"
	"strconv"
)

var CreateContact = func(response http.ResponseWriter, r *http.Request) {
	user := r.Context().Value("user") . (uint)
	contact := &models.Contact{}

	err := json.NewDecoder(r.Body).Decode(contact)
	if err != nil {
		u.Respond(response, u.Message(false, "Failture while decoding..."))
		return
	}

	contact.UserId = user
	resp := contact.Create()
	u.Respond(response, resp)
}

var GetContactsFor = func(response http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		u.Respond(response, u.Message(false, "Error in your request..."))
		return
	}

	data := models.GetContact(uint(id))
	resp := u.Message(true, "Success")
	resp["data"] = data
	u.Respond(response, resp)
}
