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
}

// func NewAPIServer(address string, db *gorm.DB) *APIService {
func NewAPIService(address string, db *gorm.DB, postUsecase *usecase.PostUsecase) *APIService {
	return &APIService{
		address: address,
		storage: &Storage{
			db: db,
		},
		postusecase: postUsecase,
		//db:      db,
	}
}

//func (s *APIService) Run() {
//fmt.Println("tes")
//}

func (s *APIService) Run() {
	//s.migrate()
	s.postusecase.Migrate()
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

	fmt.Println("Starting server on port 8080")
	http.ListenAndServe(s.address, router)
}
