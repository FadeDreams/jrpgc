// main.go
package main

import (
	"github.com/fadedreams/jrpgc/apiservice"
	"github.com/fadedreams/jrpgc/db"
	"github.com/fadedreams/jrpgc/repository"
	"github.com/fadedreams/jrpgc/usecase"
)

func main() {
	dbURL := "postgresql://postgres:postgres@localhost:5432/postgres?sslmode=disable"
	db, err := db.GetDB(dbURL)
	if err != nil {
		panic(err)
	}

	postRepository := repository.NewConcretePostRepository(db)
	postUsecase := usecase.NewPostUsecase(postRepository)

	port := ":8080"
	s := apiserver.NewAPIService(port, db, postUsecase)
	s.Run()
}
