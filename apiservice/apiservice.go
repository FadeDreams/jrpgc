package apiserver

import (
	"fmt"
	"github.com/gorilla/mux"
	//"github.com/fadedreams/jrpgc/entity"
	"github.com/fadedreams/jrpgc/usecase"

	"gorm.io/gorm"
	"net/http"
)

type Storage struct {
	db *gorm.DB
}

type APIService struct {
	address string
	//db      *gorm.DB
	storage     *Storage
	postusecase *usecase.PostUsecase
	authusecase *usecase.AuthUsecase
}

// func NewAPIServer(address string, db *gorm.DB) *APIService {
func NewAPIService(address string, db *gorm.DB, postUsecase *usecase.PostUsecase, authUsecase *usecase.AuthUsecase) *APIService {
	return &APIService{
		address: address,
		storage: &Storage{
			db: db,
		},
		postusecase: postUsecase,
		authusecase: authUsecase,
		//db:      db,
	}
}

//func (s *APIService) Run() {
//fmt.Println("tes")
//}

func UserIndex(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("Role") != "user" {
		w.Write([]byte("Not Authorized."))
		return
	}
	w.Write([]byte("Welcome, User."))
}

func AdminIndex(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("Role") != "admin" {
		w.Write([]byte("Not authorized."))
		return
	}
	w.Write([]byte("Welcome, Admin."))
}

func (s *APIService) Run() {
	//s.migrate()
	s.postusecase.Migrate()
	s.authusecase.Migrate()
	router := mux.NewRouter()
	// router.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
	// handleGreet(w, r)
	// })
	//router.HandleFunc("/hello", s.handleGreet)
	router.HandleFunc("/post/{id:[0-9]+}", s.handleGetPost).Methods("GET")
	router.HandleFunc("/post/{id:[0-9]+}", s.handleUpdatePost).Methods("PUT")
	router.HandleFunc("/post/{id:[0-9]+}", s.handleDeletePost).Methods("DELETE")
	router.HandleFunc("/posts", s.handleGetPosts)
	router.HandleFunc("/post/create", s.handleCreatePost).Methods("POST")
	//route for post
	router.HandleFunc("/auth/signup", s.handleSignup).Methods("POST")
	router.HandleFunc("/auth/signin", s.handleSignIn).Methods("POST")
	router.HandleFunc("/auth/signin", s.handleSignIn).Methods("POST")

	router.HandleFunc("/auth/user", s.handleIsAuthorized(UserIndex)).Methods("GET")
	router.HandleFunc("/auth/admin", s.handleIsAuthorized(AdminIndex)).Methods("GET")

	router.Methods("OPTIONS").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, Access-Control-Request-Headers, Access-Control-Request-Method, Connection, Host, Origin, User-Agent, Referer, Cache-Control, X-header")
	})

	fmt.Println("Starting server on port 8080")
	http.ListenAndServe(s.address, router)
}
