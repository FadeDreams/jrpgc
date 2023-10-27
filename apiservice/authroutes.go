package apiserver

import (
	"encoding/json"
	"fmt"
	//"log"
	"net/http"

	"github.com/fadedreams/jrpgc/entity"
	"github.com/golang-jwt/jwt"
	//"github.com/gorilla/mux"
	//"time"
	//"github.com/fadedreams/jrpgc/usecase"
	//"strconv"
)

func (s *APIService) handleSignup(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(400)
		w.Write([]byte("Bad request"))
		return
	}
	payload := new(entity.UserSignUp)
	err := json.NewDecoder(r.Body).Decode(payload)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Bad request: " + err.Error()))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	a, err2 := s.authusecase.SignUp(payload.Email, payload.Password)
	if err2 != nil {
		w.WriteHeader(500)
		w.Write([]byte("Internal server error"))
		return
	}
	json.NewEncoder(w).Encode(a)

}

func (s *APIService) handleSignIn(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Bad request"))
		return
	}

	// Decode the request payload containing user credentials
	var credentials entity.User
	err := json.NewDecoder(r.Body).Decode(&credentials)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Bad request: " + err.Error()))
		return
	}

	// Call the AuthUsecase to handle user sign-in
	user, err := s.authusecase.SignIn(credentials.Email, credentials.Password)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal server error"))
		return
	}

	// Respond with the user information or an authentication token if needed
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}

func (s *APIService) handleIsAuthorized(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Header["Token"] == nil {
			response := map[string]string{"message": "Token is null"}
			json.NewEncoder(w).Encode(response)
			return
		}

		var mySigningKey = []byte("secret")

		token, err := jwt.Parse(r.Header["Token"][0], func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("There was an error in parsing token.")
			}
			return mySigningKey, nil
		})

		if err != nil {
			response := map[string]string{"message": "Your Token has been expired."}
			json.NewEncoder(w).Encode(response)
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			if claims["role"] == "admin" {
				r.Header.Set("Role", "admin")
				handler.ServeHTTP(w, r)
				return
			} else if claims["role"] == "user" {
				r.Header.Set("Role", "user")
				handler.ServeHTTP(w, r)
				return
			} else {
				r.Header.Set("Role", "null")
				handler.ServeHTTP(w, r)
				return
			}
		}

		response := map[string]string{"message": "Not Authorized."}
		json.NewEncoder(w).Encode(response)
	}
}
