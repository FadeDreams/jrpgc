package apiserver

import (
	"encoding/json"
	//"fmt"
	"net/http"

	"github.com/fadedreams/jrpgc/entity"
	"github.com/gorilla/mux"
	//"time"
	//"github.com/fadedreams/jrpgc/usecase"
	//"github.com/fadedreams/jrpgc/entity"
	"strconv"
)

type GreetRes struct {
	Hello string `json:"hello"`
}

func (s *APIService) handleGreet(w http.ResponseWriter, r *http.Request) {
	//s.postusecase.TestPost()
	//s.postusecase.Migrate()
	//p, _ := s.postusecase.CreatePost("hello", "world")
	//fmt.Println(p)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	res := GreetRes{
		Hello: "world",
	}
	json.NewEncoder(w).Encode(res)
	//w.Write([]byte("hello"))
}

func (s *APIService) Migrate(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(400)
		w.Write([]byte("Bad request"))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	status := s.postusecase.Migrate()
	if status != true {
		w.WriteHeader(500)
		w.Write([]byte("Internal server error"))
		return
	}
}

func (s *APIService) handleGetPost(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(400)
		w.Write([]byte("Bad request"))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	// res := &Post{
	// ID:        1,
	// Title:     "Hello",
	// Content:   "World",
	// CreatedAt: "2023-01-01",
	// }
	idStr := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte("Bad request"))
		return
	}
	post, err2 := s.postusecase.GetPost(id)
	if err2 != nil {
		w.WriteHeader(500)
		w.Write([]byte("Internal server error"))
		return
	}
	json.NewEncoder(w).Encode(post)

}

func (s *APIService) handleCreatePost(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Bad request: Method not allowed"))
		return
	}
	payload := new(entity.CreatePostPayload)
	err := json.NewDecoder(r.Body).Decode(payload)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Bad request: " + err.Error()))
		return
	}
	// fmt.Println(payload)
	_, err2 := s.postusecase.CreatePost(payload.Title, payload.Content)
	//res, err2 := s.storage.persistPost(&Post{
	//Title:   payload.Title,
	//Content: payload.Content,
	//})
	if err2 != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal server error"))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("created"))
}

func (s *APIService) handleGetPosts(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Bad request: Method not allowed"))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	posts, err2 := s.postusecase.GetPosts()
	if err2 != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal server error"))
		return
	}
	json.NewEncoder(w).Encode(posts)

}

func (s *APIService) handleUpdatePost(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Bad request: Method not allowed"))
		return
	}

	idStr := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Bad request: Invalid post ID"))
		return
	}

	payload := new(entity.UpdatePostPayload)
	err = json.NewDecoder(r.Body).Decode(payload)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Bad request: " + err.Error()))
		return
	}

	// Call the update method from the use case
	err = s.postusecase.UpdatePost(id, payload.Title, payload.Content)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal server error"))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Post updated"))
}

func (s *APIService) handleDeletePost(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Bad request: Method not allowed"))
		return
	}

	idStr := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Bad request: Invalid post ID"))
		return
	}

	// Call the delete method from the use case
	err = s.postusecase.DeletePost(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal server error"))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Post deleted"))
}
