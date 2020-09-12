package app

import (
    "context"
    "fmt"
    jwt "github.com/dgrijalva/jwt-go"
    "net/http"
    "os"
    "restAPI_contacts/models"
    u "restAPI_contacts/utils"
    "strings"
)

var JwtAuthentication = func(next http.Handler) http.Handler {
    return http.HandlerFunc (func(response http.ResponseWriter, r *http.Request) {
        notAuth := []string{"/api/user/new", "/api/user/login"}
        requestPath := r.URL.Path
        
        for _, value := range notAuth {
            if value == requestPath {
                next.ServeHTTP(response, r)
                return
            }
        }
        respond := make(map[string]interface{})
        tokenHeader := r.Header.Get("Authorization")
        
        if tokenHeader == "" {
            respond = u.Message(false, "Missing Token")
            response.WriteHeader(http.StatusForbidden)
            response.Header().Add("Content-Type", "application/json")
            u.Respond(response, respond)
            return
        }
        
        splitted := strings.Split(tokenHeader, " ")
        if len(splitted) != 2 {
            respond = u.Message(false, "Invalid auth token")
            response.WriteHeader(http.StatusForbidden)
            response.Header().Add("Content-Type", "application/json")
            u.Respond(response, respond)
            return
        }
        
        tokenPart := splitted[1]
        tk := &models.Token{}
        
        token, err := jwt.ParseWithClaims(tokenPart, tk, func(token *jwt.Token)(interface{}, error) {
          return []byte(os.Getenv("token_password")), nil  
        })
        
        if err != nil {
            respond = u.Message(false, "Malformed auth token")
            response.WriteHeader(http.StatusForbidden)
            response.Header().Add("Content-Type", "application/json")
            u.Respond(response, respond)
            return
        }
        
        if !token.Valid {
            respond = u.Message(false, "Token is not valid")
            response.WriteHeader(http.StatusForbidden)
            response.Header().Add("Content-Type", "application/json")
            u.Respond(response, respond)
            return
        }
        
        fmt.Sprintf("User %", tk.UserId) //Debug
        ctx := context.WithValue(r.Context(), "user", tk.UserId)
        r = r.WithContext(ctx)
        next.ServeHTTP(response, r)
        
    })
}