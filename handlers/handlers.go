package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/Andressatass/trabalho-microinformatica-be/trabalho-microinformatica-be/db/entities"
	"github.com/Andressatass/trabalho-microinformatica-be/trabalho-microinformatica-be/db/repository"
)

type ErrorResponse struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

func HandleGetUserInfo(w http.ResponseWriter, r *http.Request, repo repository.MockUserRepository) {
	var userEntity entities.UserInfo

	bytes, err := io.ReadAll(r.Body)
	if err != nil {
		WriteErrorResponse(w, err, http.StatusInternalServerError)
		return
	}

	err = json.Unmarshal([]byte(bytes), &userEntity)
	if err != nil {
		WriteErrorResponse(w, err, http.StatusInternalServerError)

		return
	}

	userData, err := repo.FindById(userEntity.UUID)
	if err != nil {
		WriteErrorResponse(w, err, http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	byte, err := json.Marshal(userData)
	if err != nil {
		fmt.Fprintf(w, "Erro ao codificar o erro: %v", err)

		return
	}
	w.Write(byte)
}

func HandleCreateUser(w http.ResponseWriter, r *http.Request, repo repository.MockUserRepository) {
	var userEntity entities.UserInfo

	bytes, err := io.ReadAll(r.Body)
	if err != nil {
		WriteErrorResponse(w, err, http.StatusInternalServerError)
		return
	}

	err = json.Unmarshal([]byte(bytes), &userEntity)
	if err != nil {
		WriteErrorResponse(w, err, http.StatusInternalServerError)

		return
	}

	id, err := repo.Create(userEntity)
	if err != nil {
		WriteErrorResponse(w, err, http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	byte, err := json.Marshal(id)
	if err != nil {
		fmt.Fprintf(w, "Erro ao codificar o erro: %v", err)

		return
	}
	w.Write(byte)
}

func WriteErrorResponse(w http.ResponseWriter, err error, statusCode int) {
	errorResponse := ErrorResponse{
		Message: err.Error(),
		Code:    statusCode,
	}

	jsonData, err := json.Marshal(errorResponse)
	if err != nil {
		fmt.Fprintf(w, "Erro ao codificar o erro: %v", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}
