package handlers

import (
	"net/http"

	"github.com/Andressatass/trabalho-microinformatica-be/db/repository"
)

func RegisterRoutes(repo *repository.MockUserRepository) {
	http.HandleFunc("/GetUserInfo", func(w http.ResponseWriter, r *http.Request) {
		HandleGetUserInfo(w, r, *repo)
	})

	http.HandleFunc("/CreateUser", func(w http.ResponseWriter, r *http.Request) {
		HandleCreateUser(w, r, *repo)
	})

	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		Ping(w, r)
	})
}
