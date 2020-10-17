package auth

import (
	"context"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gonzalezlrjesus/API-Betting-Sports/models"
	u "github.com/gonzalezlrjesus/API-Betting-Sports/utils"

	jwt "github.com/dgrijalva/jwt-go"
)

// JwtAuthentication validation authentification
var JwtAuthentication = func(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//List of endpoints that doesn't require auth
		notAuth := []string{
			"/api/admin",
			"/api/admin/login",
			"/api/admin/client/app",
			"/api/clients/login",
			"/ws",
		}
		//current request path
		requestPath := r.URL.Path

		//check if request does not need authentication, serve the request if it doesn't need it
		for _, value := range notAuth {
			if value == requestPath {
				next.ServeHTTP(w, r)
				return
			}
		}

		tokenHeader := r.Header.Get("Authorization") //Grab the token from the header

		if tokenHeader == "" { //Token is missing
			sendErrorResponse(w, "Missing auth token")
			return
		}

		splitted := strings.Split(tokenHeader, " ") //The token normally comes in format `Bearer {token-body}`
		if len(splitted) != 2 {
			sendErrorResponse(w, "Invalid/Malformed auth token")
			return
		}

		tokenPart := splitted[1] //Grab the token part, what we are truly interested in
		tk := &models.Token{}

		token, err := jwt.ParseWithClaims(tokenPart, tk, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("token_password")), nil
		})

		if err != nil { //Malformed token
			sendErrorResponse(w, "Malformed authentication token")
			return
		}

		if !token.Valid { //Token is invalid, maybe not signed on this server
			sendErrorResponse(w, "Token is not valid")
			return
		}

		//Everything went well, proceed with the request and set the caller to the user retrieved from the parsed token
		log.Printf("User %v", tk.UserID) //Useful for monitoring
		ctx := context.WithValue(r.Context(), "user", tk.UserID)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r) //proceed in the middleware chain!
	})
}

func sendErrorResponse(w http.ResponseWriter, message string) {
	response := u.Message(false, message)
	u.Respond(w, response, 401)
}
