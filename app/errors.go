package app

import (
	"net/http"
	u "restAPI_contacts/utils"
)

var NotFoundHandler = func(next http.Handler) http.Handler {
	return http.HandlerFunc(func(response http.ResponseWriter, r *http.Request) {
		response.WriteHeader(http.StatusNotFound)
		u.Respond(response, u.Message(false, "This resource not found."))
		next.ServeHTTP(response, r)
	})
}
