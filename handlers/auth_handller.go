package handlers

import (
	"AuthService/models"
	"AuthService/repositories"
	"AuthService/services"
	"encoding/json"
	"net/http"
)

type AuthHandler struct {
	repository *repositories.AuthRepository
}

func NewAuthRepository(repository *repositories.AuthRepository) *AuthHandler {
	return &AuthHandler{
		repository: repository,
	}
}

func (h *AuthHandler) LoginHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	var requests models.LoginRequest
	err := json.NewDecoder(r.Body).Decode(&requests)
	if err != nil {
		handleError(w, false, err)
	}

	response, err := services.Login(&requests, *h.repository)
	if err != nil {
		handleError(w, false, err)
		return
	}
	handleResult(w, true, response, "Logged In")
}

func (h *AuthHandler) VerifyTokenHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	var requests models.VerifyTokenRequest
	err := json.NewDecoder(r.Body).Decode(&requests)
	if err != nil {
		handleError(w, false, err)
	}

	response, err := services.VerifyToken(requests.Token)
	if err != nil {
		handleError(w, false, err)
		return
	}
	handleResult(w, true, response, "Token Match")
}