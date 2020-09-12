package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"os"
	"restAPI_contacts/app"
	"restAPI_contacts/controllers"
)

func main() {
    router := mux.NewRouter()

    router.HandleFunc("/api/user/new", controllers.CreateAccount).Methods("POST")
    router.HandleFunc("/api/user/login", controllers.Authenticate).Methods("POST")
	router.HandleFunc("/api/me/contacts", controllers.GetContactsFor).Methods("GET")
	router.HandleFunc("/api/contacts/new", controllers.CreateContact).Methods("POST")

    router.Use(app.JwtAuthentication)
    
    port := os.Getenv("PORT")
    if port == "" {
        port = "8000"
    }
    
    fmt.Println(port)
    
    err := http.ListenAndServe(":" + port, router)
    
    if err != nil {
        fmt.Print(err)
    }
}