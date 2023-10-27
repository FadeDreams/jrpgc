package apiserver

import (
	"encoding/json"
	//"fmt"
	"github.com/fadedreams/jrpgc/entity"
	"net/http"
	//"github.com/gorilla/mux"
	//"time"
	//"github.com/fadedreams/jrpgc/usecase"
	//"strconv"
)

func (s *APIService) handleGJWT(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(400)
		w.Write([]byte("Bad request"))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	post, err2 := s.authusecase.GenerateJWT("email", "role")
	if err2 != nil {
		w.WriteHeader(500)
		w.Write([]byte("Internal server error"))
		return
	}
	json.NewEncoder(w).Encode(post)
}

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
